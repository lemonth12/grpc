package service

import (
	"encoding/json"
	"fmt"
)

func GetPrime(val string) (string, error) {
	stringss := make([]string, 0)
	err := json.Unmarshal([]byte(val), stringss)
	if err != nil {
		return "", err
	}
	fmt.Println(stringss)

	return "hhhhhhhh", nil
}
