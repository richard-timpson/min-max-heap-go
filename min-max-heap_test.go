package main

import (
	"fmt"
	"testing"
)

func TestBuildHeap(t *testing.T) {
	h := []int{41, 10, 71, 31, 8, 16, 46, 13, 21, 31, 51, 11}

	h = floydBuildHeap(h)

	expectedH := []int{8, 51, 71, 13, 10, 11, 46, 31, 21, 31, 41, 16}
	verifyArray(t, expectedH, h)

}

func TestInsert(t *testing.T) {
	testVals := []struct {
		in  int
		out []int
	}{
		{6, []int{6, 71, 41, 31, 10, 8, 16, 46, 51, 31, 21, 13, 11}},
		{81, []int{8, 71, 81, 31, 10, 11, 16, 46, 51, 31, 21, 13, 41}},
	}

	for i, vals := range testVals {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			h := []int{8, 71, 41, 31, 10, 11, 16, 46, 51, 31, 21, 13}
			actual := insert(h, vals.in)
			verifyArray(t, vals.out, actual)
		})
	}
}

func TestFindMin(t *testing.T) {
	h := []int{8, 71, 41, 31, 10, 11, 16, 46, 51, 31, 21, 13}

	min := findMin(h)
	if min != 8 {
		t.Errorf("Min %d is incorrect. Should be 8", min)
	}
}

func TestFindMax(t *testing.T) {
	h := []int{8, 71, 41, 31, 10, 11, 16, 46, 51, 31, 21, 13}

	max := findMax(h)
	if max != 71 {
		t.Errorf("Min %d is incorrect. Should be 71", max)
	}
}

func TestRemoveMin(t *testing.T) {
	h := []int{8, 71, 41, 31, 10, 11, 16, 46, 51, 31, 21, 13}

	h = removeMin(h)

	expectedH := []int{10, 71, 41, 31, 13, 11, 16, 46, 51, 31, 21}
	verifyArray(t, expectedH, h)
}

func TestRemoveMax(t *testing.T) {
	h := []int{8, 71, 41, 31, 10, 11, 16, 46, 51, 31, 21, 13}

	h = removeMax(h)

	expectedH := []int{8, 51, 41, 13, 10, 11, 16, 46, 31, 31, 21}
	verifyArray(t, expectedH, h)
}

func verifyArray(t *testing.T, expected []int, actual []int) {
	t.Logf("Slice values: %v", actual)
	for i, val := range actual {
		if expected[i] != val {
			t.Errorf("Value is incorrect. Got %d, want %d", val, expected[i])
		}
	}
}
