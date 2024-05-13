package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetFileExtension(filename string) string {
	ext := filepath.Ext(filename)
	return ext
}

func GetEnv(prefix, key string) string {
	if prefix != "" {
		value, ok := os.LookupEnv(fmt.Sprintf("%s_%s", prefix, key))
		if ok {
			return value
		}
	} else {
		value, ok := os.LookupEnv(key)
		if ok {
			return value
		}
	}
	return ""
}

func GetEnvKey(prefix string, key string) string {
	if prefix != "" {
		return fmt.Sprintf("%s_%s", prefix, key)
	}
	return key
}
