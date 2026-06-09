package transforms

import "strings"

func ApplyCasing(text string) string {
	words := strings.Fields(text)

	for i, word := range words {
		switch word {
		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words[i] = ""
			}
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words[i] = ""
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
				words[i] = ""
			}
		case strings.Contains(word, "up,"):
		case strings.Contains(word, "low,"):
		case strings.Contains(word, "cap,"):
		}
	}
}
