package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	EnvPrefix  string
	HTTPServer *HTTPServerConfig `JSON:"httpServer,omitempty" yaml:"httpServer,omitempty"`
}

type HTTPServerConfig struct {
	Enabled bool   `JSON:"enabled,omitempty" yaml:"enabled,omitempty"`
	Host    string `JSON:"host,omitempty" yaml:"host,omitempty"`
	Port    string `JSON:"port,omitempty" yaml:"port,omitempty"`
}

func NewConfig() (*Config, error) {
	var err error
	config := &Config{}

	// Default Configuration
	config.HTTPServer = &HTTPServerConfig{
		Enabled: true,
		Host:    "0.0.0.0",
		Port:    "8080",
	}

	// Set Default Configration File
	var configFilePath string = "./cmd/ config.json"
	value := GetEnv(config.EnvPrefix, "CONFIG_FILE")
	if value != "" {
		configFilePath = value
		log.Println(fmt.Sprintf("loading config file '%s' specified via environment variable %s", configFilePath, GetEnvKey(config.EnvPrefix, "CONFIG_FILE")))
	} else {
		log.Println(fmt.Sprintf("no config file specified via environment variable %s, using default filepath '%s'", GetEnvKey(config.EnvPrefix, "CONFIG_FILE"), configFilePath))
	}

	// Validate that Configuration file that matches stack exists.
	if _, err := os.Stat(fmt.Sprintf(configFilePath)); errors.Is(err, os.ErrNotExist) {
		// Config File Doesn't Exist
		return config, fmt.Errorf("file: '%s' does not exist", configFilePath)
	}
	// Validate Configuration file exists and Read to Byte Stream
	file, err := os.ReadFile(fmt.Sprintf("%s", configFilePath))
	if err != nil {
		return config, err
	}
	switch ext := GetFileExtension(configFilePath); strings.ToLower(ext) {
	case ".json":
		// Unmarshal the Configuration Byte Stream to Struct
		err = json.Unmarshal([]byte(file), &config)
		if err != nil {
			return config, err
		}

	default:
		// File Extension not supported
		return config, fmt.Errorf("file extension '%s' not supported", ext)

	}

	// Environment Variable Overrides

	// httpServer:host
	value = GetEnv(config.EnvPrefix, "HTTP_HOST")
	if value != "" {
		log.Println(fmt.Sprintf("config override: using environment variable %s=%s", GetEnvKey(config.EnvPrefix, "HTTP_HOST"), value))
		config.HTTPServer.Host = value
	}

	// httpServer:port
	value = GetEnv(config.EnvPrefix, "HTTP_PORT")
	if value != "" {
		log.Println(fmt.Sprintf("config override: using environment variable %s=%s", GetEnvKey(config.EnvPrefix, "HTTP_PORT"), value))
		config.HTTPServer.Port = value
	}

	// Success, Configuration Loaded.
	return config, err
}

func (c *Config) Validate() error {

	if c.HTTPServer.Enabled {
		// Validate the Configuration of the HTTP Server Adapter if Enabled
		if c.HTTPServer.Host == "" {
			return fmt.Errorf("config error: httpServer:host is required when running http server")
		}
		if c.HTTPServer.Port == "" {
			return fmt.Errorf("config error: httpServer:port is required when running http server")
		}
	}
	// Success Configuration is Valid
	return nil
}
