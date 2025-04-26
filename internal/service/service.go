package service

import (
	"fmt"
	"morse"
	"strings"
)

func Service(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("Can't translate empty string")
	}

	if strings.HasPrefix(input, ".") || strings.HasPrefix(input, "-") {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}
}
