package config

import (
	"os"
	"strings"
)

var CurrentVersion = readVersion()

func readVersion() string {
	contents, err := os.ReadFile("VERSION")
	if err != nil {
		return "HEAD"
	}

	version := strings.TrimSpace(string(contents))
	if version == "" {
		return "HEAD"
	}

	return version
}
