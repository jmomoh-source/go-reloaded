package textprocessor

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// === Hex conversion ===
		{
			name:     "hex conversion",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "hex in sentence",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},

		// === Bin conversion ===
		{
			name:     "bin conversion",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},

		// === Uppercase ===
		{
			name:     "up single word",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},

		// === Lowercase ===
		{
			name:     "low single word",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},

		// === Capitalize ===
		{
			name:     "cap single word",
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},

		// === Numbered flags ===
		{
			name:     "up with number",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},
		{
			name:     "cap with number",
			input:    "it was the age of foolishness (cap, 6)",
			expected: "It Was The Age Of Foolishness",
		},
		{
			name:     "low with number",
			input:    "IT WAS THE (low, 3) winter of despair.",
			expected: "it was the winter of despair.",
		},

		// === Punctuation ===
		{
			name:     "punctuation comma attached",
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			name:     "punctuation ellipsis",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "punctuation question mark",
			input:    "Punctuation tests are ... kinda boring ,what do you think ?",
			expected: "Punctuation tests are... kinda boring, what do you think?",
		},

		// === Quotes ===
		{
			name:     "single word quotes",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "multi word quotes",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},

		// === a/an correction ===
		{
			name:     "a to an before vowel",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
		{
			name:     "a to an before h",
			input:    "This is a huge deal",
			expected: "This is an huge deal",
		},
		{
			name:     "a before consonant stays",
			input:    "This is a big deal",
			expected: "This is a big deal",
		},
		{
			name:     "a/an full sentence",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},

		// === Full integration test from spec ===
		{
			name:     "full spec example 1",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Process(tt.input)
			if got != tt.expected {
				t.Errorf("\nInput:    %q\nExpected: %q\nGot:      %q", tt.input, tt.expected, got)
			}
		})
	}
}
