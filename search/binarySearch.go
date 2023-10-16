package main

import (
	"math"

	"golang.org/x/exp/constraints"
)

func iterBinarySearch[T constraints.Ordered](target T, list []T) bool {
	l := float64(0)
	h := float64(len(list) - 1)

	for l <= h {
		mid := math.Floor(l + (h-l)/2)
		curItem := list[int(mid)]

		if target < curItem {
			h = mid - 1
		} else if target > curItem {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}

func recBinarySearch[T constraints.Ordered](target T, list []T) bool {
	return _recBinarySearch(target, list, 0, len(list)-1)
}

func _recBinarySearch[T constraints.Ordered](
	target T,
	list []T,
	low, hi int,
) bool {
	if low <= hi {
		mid := int(math.Floor(float64(low) + (float64(hi)-float64(low))/2))

		if target == list[mid] {
			return true
		} else if target < list[mid] {
			return _recBinarySearch(target, list, low, mid-1)
		}
		return _recBinarySearch(target, list, mid+1, hi)
	}
	return false
}

func bsInsert[T constraints.Ordered](list *[]T, target T) {
	idx := bsContains(list, target)

	if idx < 0 {
		list = insert(list, -(idx + 1), target)
	}
}

func bsContains[T constraints.Ordered](list *[]T, target T) int {
	low := 0
	high := len(*list) - 1

	for low <= high {
		mid := (low + high) / 2

		if target < (*list)[mid] {
			high = mid - 1
		} else if target > (*list)[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -(low + 1)
}

func insert[T constraints.Ordered](a *[]T, index int, value T) *[]T {
	if len(*a) == index {
		*a = append(*a, value)
		return a
	}

	*a = append((*a)[:index+1], (*a)[index:]...)
	(*a)[index] = value

	return a
}
