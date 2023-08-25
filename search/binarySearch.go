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
