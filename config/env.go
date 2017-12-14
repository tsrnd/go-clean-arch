package config

import (
	"fmt"
	"os"
)

// Port func
func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}
