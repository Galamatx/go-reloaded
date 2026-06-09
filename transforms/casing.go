package transforms

import (
	"strconv"
	"strings"
)

func ApplyCasing(text string) string {
	words := strings.Fields(text)

	for i, word := range words {
		switch {
		case word == "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words[i] = ""
			}
		case word == "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words[i] = ""
			}
		case word == "(cap)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
				words[i] = ""
			}
		case strings.Contains(word, "up,"):
			inner := strings.Trim(word, "()")
			parts := strings.Split(inner, ",")
			count, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

			for j := 1; j <= count && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words[i] = ""
		case strings.Contains(word, "low,"):
			inner := strings.Trim(word, "()")
			parts := strings.Split(inner, ",")
			count, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

			for j := 1; j <= count && i-j >= 0; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words[i] = ""
		case strings.Contains(word, "cap,"):
			inner := strings.Trim(word, "()")
			parts := strings.Split(inner, ",")
			count, _ := strconv.Atoi(strings.TrimSpace(parts[1]))

			for j := 1; j <= count && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j][:1]) + strings.ToLower(words[i-j][1:])
			}
			words[i] = ""
		}
	}
	result := []string{}
	for _, w := range words {
		if w != "" {
			result = append(result, w)
		}
	}
	return strings.Join(result, " ")
}
