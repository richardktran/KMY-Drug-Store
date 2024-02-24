package app

import (
	"os"
	"strconv"
	"strings"
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

func RootPath() string {
	dir, err := os.Getwd()

	if err != nil {
		return ""
	}

	dir = strings.Replace(dir, "/pkg/app", "", 1)
	dir = strings.Replace(dir, "\\pkg\\app", "", 1)

	return dir
}
