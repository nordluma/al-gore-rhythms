package main

import "golang.org/x/exp/constraints"

func linearSearch[T constraints.Ordered](target T, list []T) bool {
	for _, elem := range list {
		if elem == target {
			return true
		}
	}

	return false
}
