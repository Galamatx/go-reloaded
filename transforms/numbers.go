package transforms

import "strings"

func ApplyNumbers(text string) string{
	words := strings.Fields(text)

	for i, word := range words {
		switch word{
			case "(hex)" {
				if i > 0{
					words[i-1] = strconv.Itoa(hexTOdec(words[i-1]))
					words[i] = ""
				}
			}
			case "(bin)" {
				
			}
		}
	}
}

func hexTOdec(hex string) int {
	result := 0

	for i := 0; i < len(hex); i++ {
		char := hex[i]
		var val int

		if char >= '0' && char <= '9' {
			val = int(char - '0')
		} else if char >= 'a' && char <= 'z' {
			val = int(char - 'a' + 10)
		} else if char >= 'A' && char <= 'Z' {
			val = int(char - 'A' + 10)
		} else if i == 1 && (char == 'x' || char == 'X') {
			continue
		} else if i == 0 && char == '0' {
			continue
		} else {
			return 0
		}

		result = (result * 16) + val
	}

	return result
}