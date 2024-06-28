package config

import (
	"os"
)

func Load() {
	{ /*err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}*/
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
