package render

import (
	"ASCII_ART/banner"
	"os"
	"testing"
)

func BenchmarkRenderText(b *testing.B) {
	content, _ := os.ReadFile("../banners/standard.txt")
	charMap := banner.BuildCharacterMap(string(content))

	text := "Hello Omar!"

	for i := 0; i < b.N; i++ {
		RenderText(text, charMap)
	}
}
