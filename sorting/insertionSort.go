package main

import "golang.org/x/exp/constraints"

func insertionSort[T constraints.Ordered](list *[]T) {
	for j := 1; j < len(*list); j++ {
		i := j - 1
		value := (*list)[j]
		for i >= 0 && (*list)[i] > value {
			(*list)[i+1] = (*list)[i]
			i--
		}
		(*list)[i+1] = value
	}
}
