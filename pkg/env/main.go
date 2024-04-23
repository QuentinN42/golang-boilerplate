package env

import (
	"fmt"
	"os"
	"strconv"
)

func Get(key string) (string, error) {
	res := os.Getenv(key)
	if res == "" {
		return "", fmt.Errorf("%s is not set", key)
	}
	return res, nil
}

func GetBool(key string) bool {
	res, err := Get(key)
	if err != nil {
		return false
	}
	ok := true
	for _, v := range []string{"false", "0", "", "no"} {
		if res == v {
			ok = false
			break
		}
	}
	return ok
}

func GetInt(key string) (int, error) {
	res, err := Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}
