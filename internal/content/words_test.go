package content

import (
	"testing"
)

func TestLoadWords(t *testing.T) {
	words := LoadWords(30)
	if len(words) != 30 {
		t.Fatalf("got %d words, want 30", len(words))
	}
	for i, w := range words {
		if w == "" {
			t.Errorf("words[%d] is empty", i)
		}
	}
}

func TestLoadWordsZero(t *testing.T) {
	words := LoadWords(0)
	if len(words) != 0 {
		t.Fatalf("got %d words, want 0", len(words))
	}
}

func TestLoadWordsMoreThanAvailable(t *testing.T) {
	// With replacement, we can get more words than the source file has
	words := LoadWords(1000)
	if len(words) != 1000 {
		t.Fatalf("got %d words, want 1000", len(words))
	}
	for i, w := range words {
		if w == "" {
			t.Errorf("words[%d] is empty", i)
		}
	}
}
