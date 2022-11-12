package utils

import "strings"

func JoinWithSpaces(inputs ...string) string {
	parts := []string{}
	parts = append(parts, inputs...)

	joined := strings.Join(parts, " ")
	return joined
}
