package transforms

import (
	"strconv"
	"strings"
)

func ApplyCasing(text string) string {
	text = strings.ReplaceAll(text, "(up)", " (up) ")
	text = strings.ReplaceAll(text, "(low)", " (low) ")
	text = strings.ReplaceAll(text, "(cap)", " (cap) ")
	text = strings.ReplaceAll(text, "(up,", " (up, ")
	text = strings.ReplaceAll(text, "(low,", " (low, ")
	text = strings.ReplaceAll(text, "(cap,", " (cap, ")

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == "" {
			continue
		}

		switch {
		case word == "(up)":
			prevIdx := findPrevWord(words, i)
			if prevIdx >= 0 {
				words[prevIdx] = strings.ToUpper(words[prevIdx])
			}
			words[i] = ""
		case word == "(low)":
			prevIdx := findPrevWord(words, i)
			if prevIdx >= 0 {
				words[prevIdx] = strings.ToLower(words[prevIdx])
			}
			words[i] = ""
		case word == "(cap)":
			prevIdx := findPrevWord(words, i)
			if prevIdx >= 0 {
				words[prevIdx] = safeCap(words[prevIdx])
			}
			words[i] = ""
		case word == "(up," || word == "(low," || word == "(cap,":
			if i+1 < len(words) {
				s := words[i+1]
				idx := 0
				isNegative := false

				// Check for an optional negative sign
				if idx < len(s) && s[idx] == '-' {
					isNegative = true
					idx++
				}

				countStr := ""
				for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
					countStr += string(s[idx])
					idx++
				}

				if countStr == "" {
					continue
				}

				count, _ := strconv.Atoi(countStr)
				if isNegative {
					count = -count
				}

				validTag := false
				if idx < len(s) && s[idx] == ')' {
					validTag = true
					words[i+1] = s[idx+1:]
				} else if idx == len(s) {
					nextPunctIdx := i + 2
					for nextPunctIdx < len(words) && words[nextPunctIdx] == "" {
						nextPunctIdx++
					}
					if nextPunctIdx < len(words) && strings.HasPrefix(words[nextPunctIdx], ")") {
						validTag = true
						words[nextPunctIdx] = words[nextPunctIdx][1:]
						words[i+1] = ""
					}
				}

				if !validTag {
					continue
				}

				if count > 0 {
					applied := 0
					for j := i - 1; j >= 0 && applied < count; j-- {
						if words[j] != "" {
							switch word {
							case "(up,":
								words[j] = strings.ToUpper(words[j])
							case "(low,":
								words[j] = strings.ToLower(words[j])
							case "(cap,":
								words[j] = safeCap(words[j])
							}
							applied++
						}
					}
				}
				words[i] = ""
			}
		}
	}

	var result []string
	for _, w := range words {
		if w != "" {
			result = append(result, w)
		}
	}
	return strings.Join(result, " ")
}

func findPrevWord(words []string, start int) int {
	p := start - 1
	for p >= 0 && words[p] == "" {
		p--
	}
	return p
}

func safeCap(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	for i, r := range runes {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return string(runes[:i]) + strings.ToUpper(string(runes[i:i+1])) + strings.ToLower(string(runes[i+1:]))
		}
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
