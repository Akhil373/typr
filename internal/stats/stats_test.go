package stats

import (
	"testing"
	"time"
)

func TestCompute(t *testing.T) {
	target := "the cat sat"
	typed := "the cat sat"
	elapsed := 60 * time.Second

	got := Compute(target, typed, elapsed)

	// 11 runes / 5 / 1 minute = 2.2 WPM
	if got.WPM < 2.19 || got.WPM > 2.21 {
		t.Errorf("WPM = %v, want ~2.2", got.WPM)
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
	// 2 of 3 words correct
	if got.Accuracy < 66.6 || got.Accuracy > 66.7 {
		t.Errorf("Accuracy = %v, want ~66.67", got.Accuracy)
	}
}

func TestComputeNetWPM(t *testing.T) {
	// 13 runes → 13/5 = 2.6 gross; 1 wrong word → net 1.6 WPM in 1 minute
	got := Compute("one two three", "one too three", time.Minute)
	if got.WPM < 1.59 || got.WPM > 1.61 {
		t.Errorf("WPM = %v, want ~1.6 (gross 2.6 minus 1 error word)", got.WPM)
	}
}

func TestComputeZeroElapsed(t *testing.T) {
	got := Compute("hello", "hello", 0)
	if got.WPM != 0 {
		t.Errorf("WPM = %v, want 0 when elapsed is 0", got.WPM)
	}
}
