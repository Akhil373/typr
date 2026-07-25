package stats

import (
	"strings"
	"time"
	"unicode/utf8"
)

type Result struct {
	Elapsed  time.Duration
	WPM      float64
	Accuracy float64
}

func Compute(target, typed string, elapsed time.Duration) Result {
	var errWords int
	charCount := utf8.RuneCountInString(target)

	targetWords := strings.Fields(target)
	typedWords := strings.Fields(typed)

	var accuracy float64
	if len(targetWords) > 0 {
		correctWords := 0
		n := min(len(targetWords), len(typedWords))
		for i := range n {
			if targetWords[i] == typedWords[i] {
				correctWords++
			}
		}
		accuracy = float64(correctWords) / float64(len(targetWords)) * 100
		errWords = (len(targetWords) - correctWords)
	}

	var wpm float64
	if secs := elapsed.Seconds(); secs > 0 {
		timeMinutes := secs / 60.0
		wpm = ((float64(charCount) / 5.0) - float64(errWords)) / timeMinutes
	}

	return Result{
		Elapsed:  elapsed,
		WPM:      max(0, wpm),
		Accuracy: accuracy,
	}
}
