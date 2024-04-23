package env

import (
	"fmt"
	"os"
)

func Get(key string) (string, error) {
	res := os.Getenv(key)
	if res == "" {
		return "", fmt.Errorf("%s is not set", key)
	}
	return res, nil
}
