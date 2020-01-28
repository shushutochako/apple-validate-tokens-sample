package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TeamID   string
	ClientID string
}

func NewConfig() *Config {
	// err := godotenv.Load(fmt.Sprintf("./.env", os.Getenv("GO_ENV")))
	err := godotenv.Load("./pkg/config/.env")
	if err != nil {
		fmt.Println(err)
	}
	return &Config{
		TeamID:   os.Getenv("TEAM_ID"),
		ClientID: os.Getenv("CLIENT_ID"),
	}
}
