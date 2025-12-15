package utils

import (
	"os"
	"strconv"
)

func GetEnv(key string, defaultVal string) string {
	val, found := os.LookupEnv(key)
	if found {
		return val
	}
	return defaultVal
}

func GetEnvInt(key string, defaultVal int) int {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}

	return res
}

func GetEnvBool(key string, defaultVal bool) bool {
	val, found := os.LookupEnv(key)
	if !found {
		return defaultVal
	}
	res, err := strconv.ParseBool(val)
	if err != nil {
		return defaultVal
	}

	return res
}
