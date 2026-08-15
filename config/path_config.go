package config

import (
	"os"
	"github.com/spf13/viper"
	"path/filepath" 
)

func Init() {
	viper.SetEnvPrefix("todocli")
	// automatically gets all env variables with the todocli prefixed variables 
	// just gets them for us otherwise we have to bind each one of them seperately using viper.Bindenv
	viper.AutomaticEnv()
	// use default
	home, _ := os.UserHomeDir()
	viper.SetDefault("filepath", filepath.Join(home, ".todocli", "Book1.xlsx"))
	viper.SetDefault("backupfilepath", filepath.Join(home, ".todocli", "Backup.xlsx"))
	viper.SetDefault("configpath", filepath.Join(home, ".todocli"))

}

func ReadPath(file string) string {
	return viper.GetString(file)
}