package main

import (
	"golang.org/x/exp/constraints"
)

func merge[T constraints.Ordered](left, right []T) []T {
	final := []T{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}

	for i < len(left) {
		final = append(final, left[i])
		i++
	}

	for j < len(right) {
		final = append(final, right[j])
		j++
	}

	return final
}

func mergeSort[T constraints.Ordered](list []T) []T {
	if len(list) < 2 {
		return list
	}

	left := mergeSort(list[:len(list)/2])
	right := mergeSort(list[len(list)/2:])
	return merge(left, right)
}
