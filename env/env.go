package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Test() {
	os.Setenv("SERVER_ENV", "test")
	Load()
}

func Load(forced ...string) {
	env := os.Getenv("SERVER_ENV")
	if env == "" {
		env = "prod"
	}

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(dirname + "/Work/Projects/server-go/.env." + env)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
