package config

import (
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/joho/godotenv"
)

// Store app configuration
// Map to the environment variables sets by env flag
// To skip set the Parameter by environment variable you must sen env flag to "-"
type AppConfig struct {
	// App
	TCPPort  string `env:"PORT"`      // Port for rest API
	LogLevel string `env:"LOG_LEVEL"` // Logining DEBUG/ERROR/WARN

}

var conf AppConfig
var once sync.Once

// Get pointer to the app config
// The first run runs setUp function that get configuration values from environ variables
func GetConfig() *AppConfig {
	once.Do(setUp)
	return &conf
}

// Get configuration values from the environ variables
// If ENVIRONMENT variable is dev then trying to get environ variables from the .env file
func setUp() {
	if env := os.Getenv("ENVIRONMENT"); env == "dev" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("No .env files found. Using real environment")
		}

	}
	v := reflect.ValueOf(&conf).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		varName, _ := f.Tag.Lookup("env")
		if varName == "-" {
			continue
		}
		env, ok := os.LookupEnv(varName)
		if ok {
			v.Field(i).SetString(env)
		}

	}
}
