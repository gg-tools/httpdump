package utils

import (
	"os"
	"strconv"
)

func Env(name string, defaultValue string) string {
	val, ok := os.LookupEnv(name)
	if !ok || val == "" {
		return defaultValue
	}

	return val
}

func EnvInt(name string, defaultValue int) int {
	val, ok := os.LookupEnv(name)
	if !ok || val == "" {
		return defaultValue
	}

	if intVal, err := strconv.Atoi(val); err == nil {
		return intVal
	} else {
		return defaultValue
	}
}

func Int(val string, defaultValue int) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}

	return num
}
