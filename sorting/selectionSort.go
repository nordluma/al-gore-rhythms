package main

import "golang.org/x/exp/constraints"

func selectMax[T constraints.Ordered](list *[]T, left, right int) int {
	maxPos := left
	i := left
	for i <= right {
		if (*list)[i] >= (*list)[maxPos] {
			maxPos = i
		}
		i++
	}

	return maxPos
}

func selectionSort[T constraints.Ordered](list *[]T) {
	for i := len((*list)) - 1; i >= 1; i-- {
		maxPos := selectMax(list, 0, i)
		if maxPos != i {
			tmp := (*list)[i]
			(*list)[i] = (*list)[maxPos]
			(*list)[maxPos] = tmp
		}
	}
}
