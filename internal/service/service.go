package service

import (
	"strings"

	"MorzeText/pkg/morse"
)

func Service(input string) (string, error) {
	if input == "" {
		return "", morse.ErrNoEncoding{Text: "Can't translate empty string"}
	}

	isMorse := strings.ContainsFunc(input, func(r rune) bool {
		return r == '.' || r == '-'
	})

	if isMorse {
		return morse.ToText(input), nil
	} else {
		return morse.ToMorse(input), nil
	}
}
