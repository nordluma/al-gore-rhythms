package main

import "golang.org/x/exp/constraints"

func bubbleSort[T constraints.Ordered](list *[]T) {
	for i := 0; i < len(*list); i++ {
		for j := 0; j < len(*list)-1-i; j++ {
			if (*list)[j] > (*list)[j+1] {
				tmp := (*list)[j]
				(*list)[j] = (*list)[j+1]
				(*list)[j+1] = tmp
			}
		}
	}
}
