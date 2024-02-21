package main

import (
	"os"
	"strconv"
)

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

func IsDebug() bool {
	isDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))

	if err != nil {
		return false
	}

	return isDebug
}
