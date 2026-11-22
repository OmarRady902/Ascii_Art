package banner

import "testing"

func TestBuildCharacterMap(t *testing.T) {
	content := "ignore\n" +
		"A1\n" +
		"A2\n" +
		"A3\n" +
		"A4\n" +
		"A5\n" +
		"A6\n" +
		"A7\n" +
		"A8\n" +
		"ignore\n" +
		"B1\n" +
		"B2\n" +
		"B3\n" +
		"B4\n" +
		"B5\n" +
		"B6\n" +
		"B7\n" +
		"B8\n"

	got := BuildCharacterMap(content)

	tests := []struct {
		name      string
		char      rune
		lineIndex int
		expected  string
	}{
		{
			name:      "Space first line",
			char:      ' ',
			lineIndex: 0,
			expected:  "A1",
		},
		{
			name:      "Space last line",
			char:      ' ',
			lineIndex: 7,
			expected:  "A8",
		},
		{
			name:      "Exclamation first line",
			char:      '!',
			lineIndex: 0,
			expected:  "B1",
		},
		{
			name:      "Exclamation last line",
			char:      '!',
			lineIndex: 7,
			expected:  "B8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue := got[tt.char][tt.lineIndex]

			if gotValue != tt.expected {
				t.Errorf("got %q, want %q", gotValue, tt.expected)
			}
		})
	}
}
