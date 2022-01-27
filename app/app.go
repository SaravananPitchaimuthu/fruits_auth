package app

import (
	"fmt"
	"os"

	"github.com/SaravananPitchaimuthu/Bookstore_users-api/logger"
)

func Start() {
	sanityCheck()
	router := mux.NewRouter()

}

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Error(fmt.Sprintf("Environment variable %s is not defined", k))
		}
	}
}
