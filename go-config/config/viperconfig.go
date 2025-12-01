package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperConfig struct{}

func (v ViperConfig) build() {
	initConfigs()
}

func initConfigs() {
	var configFilePath string
	viper.SetConfigName("config")
	if configFilePath != "" {
		stat, err := os.Stat(configFilePath)
		if err != nil {
			fmt.Println("Error reading config file:", err)
		}
		if stat.IsDir() {
			viper.AddConfigPath(configFilePath)
		} else {
			viper.AddConfigPath(configFilePath)
		}
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath(os.Getenv("/etc/appname/"))
	viper.AddConfigPath(os.Getenv("$HOME/.appname"))
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file : %s", err)
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("Config file changed: %s", e.Name)
		})
	}
	viper.WatchConfig()
}

func (v ViperConfig) isSet(key string) bool {
	return viper.IsSet(key)
}

func (v ViperConfig) GetString(key string) string {
	// Implementation for getting a string value
	return viper.GetString(key)
}

func (v ViperConfig) GetInt(key string) int {
	// Implementation for getting an int value
	return viper.GetInt(key)
}

func (v ViperConfig) GetBool(key string) bool {
	// Implementation for getting a bool value
	return viper.GetBool(key)
}

func (v ViperConfig) GetFloat(key string) float64 {
	// Implementation for getting a float64 value
	return viper.GetFloat64(key)
}
func (v ViperConfig) GetIntSlice(key string) []int {
	// Implementation for getting a slice of int values
	return viper.GetIntSlice(key)
}
func (v ViperConfig) GetStringSlice(key string) []string {
	// Implementation for getting a slice of string values
	return viper.GetStringSlice(key)
}

func (v ViperConfig) GetStringMap(key string) map[string]interface{} {
	// Implementation for getting a map of string to interface{}
	return viper.GetStringMap(key)
}

func (v ViperConfig) GetStringMapString(key string) map[string]string {
	// Implementation for getting a map of string to string
	return viper.GetStringMapString(key)
}
