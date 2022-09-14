package main

import "testing"

func TestBuildHeap(t *testing.T) {
	h := []int{41, 10, 71, 13, 8, 16, 46, 13, 21, 31, 51, 11}

	h = floydBuildHeap(h)

	t.Logf("Slice values: %v", h)
	expectedH := []int{8, 71, 41, 31, 10, 11, 16, 46, 51, 31, 21, 13}
	for i, val := range h {
		t.Logf("%d,", val)
		if expectedH[i] != val {
			t.Errorf("Value is incorrect. Got %d, want %d", val, expectedH[i])
		}
	}
}
