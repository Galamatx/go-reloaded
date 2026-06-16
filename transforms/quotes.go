package transforms

import "strings"

func ApplyQuotes(text string) string {
	text = cleanQuoteMargins(text, "'")
	text = cleanQuoteMargins(text, "\"")
	return text
}

func cleanQuoteMargins(text string, target string) string {
	var result strings.Builder
	open := false
	runes := []rune(text)
	i := 0

	for i < len(runes) {
		if string(runes[i]) == target {
			if target == "'" && i > 0 && i+1 < len(runes) &&
				((runes[i-1] >= 'a' && runes[i-1] <= 'z') || (runes[i-1] >= 'A' && runes[i-1] <= 'Z')) &&
				((runes[i+1] >= 'a' && runes[i+1] <= 'z') || (runes[i+1] >= 'A' && runes[i+1] <= 'Z')) {
				result.WriteRune(runes[i])
				i++
				continue
			}

			if !open {
				hasClosing := false
				for nextIdx := i + 1; nextIdx < len(runes); nextIdx++ {
					if string(runes[nextIdx]) == target {
						hasClosing = true
						break
					}
				}

				if hasClosing {
					if result.Len() > 0 && !strings.HasSuffix(result.String(), " ") {
						result.WriteRune(' ')
					}
					result.WriteString(target)
					i++
					for i < len(runes) && runes[i] == ' ' {
						i++
					}
					open = true
				} else {
					result.WriteString(target)
					i++
				}
			} else {
				s := result.String()
				for len(s) > 0 && s[len(s)-1] == ' ' {
					s = s[:len(s)-1]
				}
				result.Reset()
				result.WriteString(s)
				result.WriteString(target)
				result.WriteRune(' ')
				i++

				for i < len(runes) && runes[i] == ' ' {
					i++
				}

				for i < len(runes) && (runes[i] == '.' || runes[i] == ',' || runes[i] == '!' || runes[i] == '?' || runes[i] == ':' || runes[i] == ';') {
					s2 := result.String()
					for len(s2) > 0 && s2[len(s2)-1] == ' ' {
						s2 = s2[:len(s2)-1]
					}
					result.Reset()
					result.WriteString(s2)
					result.WriteRune(runes[i])
					result.WriteRune(' ')
					i++
				}
				open = false
			}
		} else {
			result.WriteRune(runes[i])
			i++
		}
	}
	return strings.Join(strings.Fields(result.String()), " ")
}
