package environment

import "os"

func SetEnv(env string, value string) {
	os.Setenv(env, GetEnv(env, value))
}

func Current() string {
	return GetEnv("GO_ENV", "development")
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
