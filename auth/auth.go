package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"os/exec"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	appconfig "TODOCLI/config"
)

func authenticate() (*oauth2.Token, error) {
    config := &oauth2.Config{
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        Endpoint:     google.Endpoint,
        Scopes: []string{
            "https://www.googleapis.com/auth/calendar.readonly",
        },
        RedirectURL: "http://localhost:8080/callback",
    }

    server := &http.Server{
        Addr: "localhost:8080",
    }

    codeChan := make(chan string)

    http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")

        if code == "" {
            http.Error(w, "Authorization failed", http.StatusBadRequest)
            return
        }

        fmt.Fprintln(w, "Authorization successful! You can close this window.")

        codeChan <- code
    })

    go func() {
        if err := server.ListenAndServe(); err != nil &&
            err != http.ErrServerClosed {
            log.Println(err)
        }
    }()

    authURL := config.AuthCodeURL(
        "random-state-value",
        oauth2.AccessTypeOffline,
    )

    fmt.Println("Opening browser...")
    exec.Command("open", authURL).Start()

    code := <-codeChan

    token, err := config.Exchange(context.Background(), code)
    if err != nil {
        return nil, err
    }

    server.Shutdown(context.Background())

    return token, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := filepath.Join(appconfig.ReadPath("configpath"), "token.json")
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	codeChan := make(chan string)

	server := &http.Server{
		Addr: "localhost:8080",
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		if code == "" {
			http.Error(w, "Authorization failed", http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, "Authorization successful! You can close this window.")

		codeChan <- code
	})

	server.Handler = mux

	go func() {
		if err := server.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	// IMPORTANT:
	// This must match the redirect URI registered in Google Cloud.
	config.RedirectURL = "http://localhost:8080/callback"

	authURL := config.AuthCodeURL(
		"state-token",
		oauth2.AccessTypeOffline,
	)

	fmt.Println("Opening browser...")

	if err := exec.Command("open", authURL).Start(); err != nil {
		log.Printf("Unable to open browser: %v", err)
		fmt.Println("Open this URL manually:")
		fmt.Println(authURL)
	}

	// Wait for Google to redirect back to localhost.
	authCode := <-codeChan

	// Exchange authorization code for tokens.
	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	// Stop the temporary callback server.
	server.Shutdown(context.Background())

	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		log.Fatalf("Unable to create config directory: %v", err)
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Calender() {
	ctx := context.Background()
	b, err := os.ReadFile(filepath.Join(appconfig.ReadPath("configpath"), "credentials.json"))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	config.RedirectURL = "http://localhost:8080/callback"

	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}