package utils

import "strings"

type Input struct {
}

type InputInterface interface {
	Sanitize(input string) string
}

func NewInput() InputInterface {
	return &Input{}
}

func (u *Input) Sanitize(input string) string {
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\r", "")
	return strings.TrimSpace(input)
}
