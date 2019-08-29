package helpers

import "testing"

func BarTest(t *testing.T) {
	want := 60
	if got := Bar(); got != want {
		t.Errorf("Bar() = %v, but got %v", want, got)
	}
} 