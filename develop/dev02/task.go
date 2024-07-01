package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func unpackString(str string) (string, error) {
	result := ""
	n := 1
	var currSymbol int32

	for i, c := range str {
		if !unicode.IsNumber(c) {
			result += strings.Repeat(string(currSymbol), n)
			currSymbol = c
			n = 1
		} else {
			if _, err := fmt.Sscanf(str[i:], "%d", &n); err != nil {
				return "", err
			}
		}
	}

	result += strings.Repeat(string(currSymbol), n)
	result = strings.ReplaceAll(result, "\x00", "")

	if len(result) < len(str) {
		return "", errors.New("incorrect string")
	}

	return result, nil
}

func main() {}
