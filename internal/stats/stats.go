package stats

import (
	"time"
)

type Result struct {
	Elapsed  time.Duration
	WPM      float64
	Accuracy float64
}

func Compute(target, typed string, elapsed time.Duration) Result {
	typedRunes := []rune(typed)
	targetRunes := []rune(target)

	targetCount := len(targetRunes)
	if targetCount == 0 {
		return Result{Elapsed: elapsed}
	}

	correctRunes := 0
	typedCount := len(typedRunes)
	n := min(targetCount, typedCount)

	for i := range n {
		if typedRunes[i] == targetRunes[i] {
			correctRunes++
		}
	}

	var accuracy float64
	if typedCount > 0 {
		accuracy = (float64(correctRunes) / float64(typedCount)) * 100.0
	}

	var wpm float64
	if secs := elapsed.Seconds(); secs > 0 {
		timeMinutes := secs / 60.0
		wpm = (float64(correctRunes) / 5.0) / timeMinutes
	}

	return Result{
		Elapsed:  elapsed,
		WPM:      max(0, wpm),
		Accuracy: accuracy,
	}
}
