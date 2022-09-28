package main

import (
	"math"
)

func findMin(h []int) int {
	return h[0]
}

func findMax(h []int) int {
	return h[findMaxIdx(h)]
}

func findMaxIdx(h []int) int {
	if len(h) >= 3 {
		if h[1] > h[2] {
			return 1
		} else {
			return 2
		}
	} else if len(h) == 2 {
		return 1
	} else {
		return -1
	}
}

func removeMin(h []int) []int {
	if len(h) > 0 {
		lastVal := h[len(h)-1]
		h = h[:len(h)-1]
		h[0] = lastVal
		pushDown(h, 0)
	}
	return h
}

func removeMax(h []int) []int {
	if len(h) > 0 {
		lastVal := h[len(h)-1]
		h = h[:len(h)-1]
		maxIdx := findMaxIdx(h)
		if maxIdx != -1 {
			h[maxIdx] = lastVal
			pushDown(h, maxIdx)
		}
	}
	return h
}

func floydBuildHeap(h []int) []int {
	for i := len(h) / 2; i >= 0; i-- {
		pushDown(h, i)
	}
	return h
}

func pushDown(h []int, i int) {
	if isMinLevel(i) {
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
				if h[idx] < h[m] {
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
				if h[idx] > h[m] {
					m = idx
				}
			}
		}
		// if index is grandchild
		if isIdxGrandchild(i, m) {
			if h[m] > h[i] {
				h[m], h[i] = h[i], h[m]
				parentM := parent(m)
				if h[m] < h[parentM] {
					h[m], h[parentM] = h[parentM], h[m]
				}
				pushDownMax(h, m)
			}
		} else if h[m] > h[i] {
			h[m], h[i] = h[i], h[m]
		}
	}
}

func insert(h []int, value int) []int {
	h = append(h, value)
	return pushUp(h, len(h)-1)

}

func pushUp(h []int, i int) []int {
	if i != 0 {
		if isMinLevel(i) {
			parentIdx := parent(i)
			if h[i] > h[parentIdx] {
				h[i], h[parentIdx] = h[parentIdx], h[i]
				pushUpMax(h, parentIdx)
			} else {
				pushUpMin(h, i)
			}
		} else {
			parentIdx := parent(i)
			if h[i] < h[parentIdx] {
				h[i], h[parentIdx] = h[parentIdx], h[i]
				pushUpMin(h, parentIdx)
			} else {
				pushUpMax(h, i)
			}
		}
	}
	return h
}

func pushUpMin(h []int, i int) {
	if idxHasGrandparent(i) && h[i] < h[grandparent(i)] {
		h[i], h[grandparent(i)] = h[grandparent(i)], h[i]
		pushUpMin(h, grandparent(i))
	}
}

func pushUpMax(h []int, i int) {
	if idxHasGrandparent(i) && h[i] > h[grandparent(i)] {
		h[i], h[grandparent(i)] = h[grandparent(i)], h[i]
		pushUpMax(h, grandparent(i))
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

func isMinLevel(i int) bool {
	return getLevel(i)%2 == 0
}

func isMaxLevel(i int) bool {
	return getLevel(i)%2 != 0
}

func getLeftAndRightIdx(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

func isIdxGrandchild(i int, childIdx int) bool {
	_, rightIdx := getLeftAndRightIdx(i)
	return childIdx > rightIdx
}

func parent(i int) int {
	return int(math.Ceil(float64(i)/2)) - 1
}

func grandparent(i int) int {
	return parent(parent(i))
}

func idxHasGrandparent(i int) bool {
	return grandparent(i) >= 0
}
