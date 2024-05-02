package utils

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type EnvMap map[string]string

var Env EnvMap

func LoadEnv() {
	if Env != nil {
		return
	}

	godotenv.Load()
	items := EnvMap{}

	for _, item := range os.Environ() {
		key, val, _ := strings.Cut(item, "=")
		items[key] = val
	}

	Env = items
}

func (env EnvMap) Get(key, fallback string) string {
	if value, exists := env[key]; exists {
		return value
	}

	return fallback
}
