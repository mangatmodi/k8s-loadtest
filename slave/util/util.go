package util

import (
	"errors"
	"os"
)

func GetEnv(envVar string) (string, error) {
	var value = os.Getenv(envVar)
	if value == "" {
		return "", errors.New("Environment variable not set: " + envVar)
	}
	return value, nil
}
