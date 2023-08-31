package main

import "golang.org/x/exp/constraints"

func heapify[T constraints.Ordered](list *[]T, idx, max int) {
	var largest int
	left := 2*idx + 1
	right := 2*idx + 2

	if left < max && (*list)[left] > (*list)[idx] {
		largest = left
	} else {
		largest = idx
	}

	if right < max && (*list)[right] > (*list)[largest] {
		largest = right
	}

	// If largest is not already the parent the swap and propagate
	if largest != idx {
		tmp := (*list)[idx]
		(*list)[idx] = (*list)[largest]
		(*list)[largest] = tmp

		heapify(list, largest, max)
	}
}

func buildHeap[T constraints.Ordered](list *[]T, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		heapify(list, i, n)
	}
}

func heapSort[T constraints.Ordered](list *[]T) {
	n := len(*list)
	buildHeap(list, n)

	for i := n - 1; i >= 0; i-- {
		tmp := (*list)[0]
		(*list)[0] = (*list)[i]
		(*list)[i] = tmp

		heapify(list, 0, i)
	}
}
