package render

import "strings"

func RenderText(text string, charMap map[rune][]string) string {
	var result strings.Builder
	result.Grow(len(text) * 64)
	for i := 0; i < 8; i++ {
		for _, char := range text {
			result.WriteString(charMap[char][i])
		}
		result.WriteString("\n")
	}

	return result.String()
}
