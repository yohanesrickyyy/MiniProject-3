package config

import "os"

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetPort() string {
	return GetEnv("PORT")
}
