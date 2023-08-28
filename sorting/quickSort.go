package main

import "golang.org/x/exp/constraints"

func partition[T constraints.Ordered](list *[]T, low, high int) int {
	pivot := (*list)[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if (*list)[i] <= pivot {
			idx++
			tmp := (*list)[i]
			(*list)[i] = (*list)[idx]
			(*list)[idx] = tmp
		}
	}

	idx++
	(*list)[high] = (*list)[idx]
	(*list)[idx] = pivot

	return idx
}

func _qs[T constraints.Ordered](list *[]T, low, high int) {
	if low >= high {
		return
	}

	pivotIdx := partition(list, low, high)

	_qs(list, low, pivotIdx-1)
	_qs(list, pivotIdx+1, high)
}

func quickSort[T constraints.Ordered](list *[]T) {
	_qs(list, 0, len(*list)-1)
}
