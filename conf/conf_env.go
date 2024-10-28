package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port string
	Node string
}

func NewConf() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return &Config{
		Port: getEnv("PORT_LISTEN", ":8080"), // Default to "8080" if not set
		Node: getEnv("NODE", "defaultNode"),  // Default to "defaultNode" if not set
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
