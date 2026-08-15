package google

import (
	"TODOCLI/utility"
	"context"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	appconfig "TODOCLI/config"
	"path/filepath"
	"os"
	"log"
	"fmt"
	"golang.org/x/oauth2/google"
	"TODOCLI/auth"
)


func Calender(taskType string, startTime string, endTime string, title string, location string, description string, timeZone string) {
	ctx := context.Background()
	b, err := os.ReadFile(filepath.Join(appconfig.ReadPath("configpath"), "credentials.json"))
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	config.RedirectURL = "http://localhost:8080/callback"

	client := auth.GetClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	switch taskType {
	case "event":
		CreateEvent(srv, startTime, endTime, title, location, description, timeZone)
	default:
		utility.PrintCMD("No task type defined", "Red")
	}

}

func CreateEvent(srv *calendar.Service, startTime string, endTime string, title string, location string, description string, timeZone string){
	event := &calendar.Event{
    Summary:     title,
    Location:    location,
    Description: description,
    Start: &calendar.EventDateTime{
        DateTime: startTime,
        TimeZone: timeZone,
    },
    End: &calendar.EventDateTime{
        DateTime: endTime,
        TimeZone: timeZone,
    },
	}

	createdEvent, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		log.Fatalf("Unable to create event: %v", err)
	}
	fmt.Printf("Event created: %s\n", createdEvent.HtmlLink)

}