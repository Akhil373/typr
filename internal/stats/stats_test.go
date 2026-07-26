package stats

import (
	"math"
	"testing"
	"time"
)

func TestCompute(t *testing.T) {
	target := "the cat sat"
	typed := "the cat sat"
	elapsed := 60 * time.Second

	got := Compute(target, typed, elapsed)

	// 11 runes / 5 / 1 minute = 2.2 WPM
	if math.Abs(got.WPM-2.2) > 0.01 {
		t.Errorf("WPM = %v, want 2.2", got.WPM)
	}
	if got.Accuracy != 100 {
		t.Errorf("Accuracy = %v, want 100", got.Accuracy)
	}
	if got.Elapsed != elapsed {
		t.Errorf("Elapsed = %v, want %v", got.Elapsed, elapsed)
	}
}

func TestComputePartialAccuracy(t *testing.T) {
	got := Compute("one two three", "one too three", time.Minute)

	// Accuracy: 12 correct out of 13 typed = ~92.31%
	wantAccuracy := (12.0 / 13.0) * 100.0
	if math.Abs(got.Accuracy-wantAccuracy) > 0.01 {
		t.Errorf("Accuracy = %v, want ~%v", got.Accuracy, wantAccuracy)
	}

	// Net WPM: (12 correct chars / 5) / 1 min = 2.4 WPM
	if math.Abs(got.WPM-2.4) > 0.01 {
		t.Errorf("WPM = %v, want 2.4", got.WPM)
	}
}

func TestComputeZeroElapsed(t *testing.T) {
	got := Compute("hello", "hello", 0)
	if got.WPM != 0 {
		t.Errorf("WPM = %v, want 0 when elapsed is 0", got.WPM)
	}
}

func TestComputeEmptyTyped(t *testing.T) {
	got := Compute("hello", "", time.Minute)
	if got.WPM != 0 || got.Accuracy != 0 {
		t.Errorf("WPM = %v, Accuracy = %v, want 0 for both when typed is empty", got.WPM, got.Accuracy)
	}
}
