package main

import (
	"math"
)

func floydBuildHeap(h []int) []int {
	for i := len(h) / 2; i > 0; i-- {
		pushDown(h, i)
	}
	return h
}

func pushDown(h []int, i int) {
	// if i is even, push down min
	level := getLevel(i)
	if level%2 == 0 {
		pushDownMin(h, i)
	} else {
		pushDownMax(h, i)
	}
}

func pushDownMin(h []int, i int) {
	descendants := getChildrenAndGrandchildren(h, i, math.MaxInt)
	// if i has children
	if descendants[0] != -1 || descendants[1] != -1 {
		// find index of smallest value
		m := descendants[0]
		for _, idx := range descendants {
			if idx > -1 {
				if h[m] < h[idx] {
					m = idx
				}
			}
		}
		// if index is grandchild
		if isIdxGrandchild(i, m) {
			if h[m] < h[i] {
				h[m], h[i] = h[i], h[m]
				if h[m] > h[parent(m)] {
					h[m], h[parent(m)] = h[parent(m)], h[m]
				}
				pushDownMin(h, m)
			}
		} else if h[m] < h[i] {
			h[m], h[i] = h[i], h[m]
		}
	}
}

func pushDownMax(h []int, i int) {
	descendants := getChildrenAndGrandchildren(h, i, math.MaxInt)
	// if i has children
	if descendants[0] != -1 || descendants[1] != -1 {
		// find index of greatest value
		m := descendants[0]
		for _, idx := range descendants {
			if idx > -1 {
				if h[m] > h[idx] {
					m = idx
				}
			}
		}
		// if index is grandchild
		if isIdxGrandchild(i, m) {
			if h[m] > h[i] {
				h[m], h[i] = h[i], h[m]
				if h[m] < h[parent(m)] {
					h[m], h[parent(m)] = h[parent(m)], h[m]
				}
				pushDownMax(h, m)
			}
		} else if h[m] > h[i] {
			h[m], h[i] = h[i], h[m]
		}
	}
}

func getChildrenAndGrandchildren(h []int, i int, absentValue int) [6]int {
	// ret will be a fix sized array.
	// there are 2 children, and 4 grandchildren
	children := getChildren(h, i, absentValue)
	leftIdx, rightIdx := getLeftAndRightIdx(i)
	leftChildren := getChildren(h, leftIdx, absentValue)
	rightChildren := getChildren(h, rightIdx, absentValue)
	return [6]int{
		children[0],
		children[1],
		leftChildren[0],
		leftChildren[1],
		rightChildren[0],
		rightChildren[1],
	}
}

func getChildren(h []int, i int, absentValue int) [2]int {
	var ret [2]int
	leftIdx, rightIdx := getLeftAndRightIdx(i)
	if len(h) > leftIdx {
		ret[0] = leftIdx
	} else {
		ret[0] = -1
	}
	if len(h) > rightIdx {
		ret[1] = rightIdx
	} else {
		ret[1] = -1
	}
	return ret
}

func getLevel(i int) int {
	return int(math.Floor(math.Log2(float64(i + 1))))
}

func getLeftAndRightIdx(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

func isIdxGrandchild(i int, childIdx int) bool {
	leftIdx, rightIdx := getLeftAndRightIdx(i)
	return childIdx > leftIdx || childIdx > rightIdx
}

func parent(i int) int {
	return int(math.Ceil(float64(i)/2)) - 1
}

/*


level:  0  1   1   2   2   2   2   3   3   3   3   3
index:  0  1   2   3   4   5   6   7   8   9   10  11
array: [8, 71, 41, 31, 10, 11, 16, 46, 51, 31, 21, 13]

how to determine if i has children?

children of 71 (1,1) -> 31(3), 10(4) : i + 2, i +3
log 1
children of 41 (2,2) -> 11(5), 16(6) : i + 3, i + 4
children of 31 (3,2) -> 46(7), 51(8) : i + 4, i + 5
children of 10 (4,2) -> 31(9), 21(10): i + 5, i + 6
children of 11 (5,2) -> 13(11)       : i + 6

children of i = i * 2 + 1, i * 2 + 2

parent of i = i / 2 - 1
*/
