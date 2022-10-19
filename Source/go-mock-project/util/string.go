package util

import "strings"

func SplitStringToArray(text string) []string {
	arr := make([]string, 0)
	if len(text) > 0 {
		arr = strings.Split(text, ",")
	}
	return arr
}
