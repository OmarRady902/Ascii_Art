package banner

import (
	"os"
	"testing"
)

func BenchmarkLoadBannerFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoadBannerFile("../banners/standard.txt")
	}
}

func BenchmarkBuildCharacterMap(b *testing.B) {
	content, _ := os.ReadFile("../banners/standard.txt")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuildCharacterMap(string(content))
	}
}
