package helper

import (
	"os"
)

func GetEnvWithCallback(envVar, callback string) string {
	env := os.Getenv(envVar)
	if env != "" {
		return env
	} else {
		return callback
	}
}
