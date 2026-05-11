package bootstrap

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var Env map[string]string

func SetupConfig(path string) error {
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = godotenv.Load(filepath.Join(workDir, path))
	if err != nil {
		return err
	}
	return nil
}

func GetEnv(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}
	return result
}
