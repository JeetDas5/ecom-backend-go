package env

import (
	"fmt"
	"os"
)

func GetString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		fmt.Println("env val:", val)
		return val
	}

	return fallback
}
