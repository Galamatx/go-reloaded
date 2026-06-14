package transforms

import "strings"

func ApplyPunctuation(text string) string {
	var prep strings.Builder
	for _, r := range text {
		if r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';' || r == ')' {
			prep.WriteRune(' ')
			prep.WriteRune(r)
			prep.WriteRune(' ')
		} else {
			prep.WriteRune(r)
		}
	}

	words := strings.Fields(prep.String())

	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}

		if isPunct(words[i]) {
			prevIdx := i - 1
			for prevIdx >= 0 && words[prevIdx] == "" {
				prevIdx--
			}

			if prevIdx >= 0 {
				if words[i] == "." && strings.HasSuffix(words[prevIdx], ".") {
					words[prevIdx] = words[prevIdx] + "."
				} else {
					words[prevIdx] = words[prevIdx] + words[i]
				}
				words[i] = ""
			}
		}
	}

	result := []string{}
	for _, w := range words {
		if w != "" {
			result = append(result, w)
		}
	}

	joined := strings.Join(result, " ")
	var final strings.Builder
	runes := []rune(joined)
	for k := 0; k < len(runes); k++ {
		final.WriteRune(runes[k])
		if isPunctCh(byte(runes[k])) && k+1 < len(runes) {
			next := runes[k+1]
			if (next >= 'a' && next <= 'z') || (next >= 'A' && next <= 'Z') || (next >= '0' && next <= '9') {
				final.WriteRune(' ')
			}
		}
	}
	return final.String()
}

func isPunct(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		if c == ',' || c == '.' || c == '!' || c == '?' || c == ';' || c == ':' || c == ')' {
			return true
		}
	}

	return false
}

func isPunctCh(c byte) bool {
	if c == ',' || c == '.' || c == '!' || c == '?' || c == ';' || c == ':' || c == ')' {
		return true
	}
	return false
}
