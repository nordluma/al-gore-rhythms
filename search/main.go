package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

func createListOfStrings() []string {
	return []string{"hello", "mom", "im", "doing", "algorithms"}
}

func createListOfIntegers() []int {
	return []int{666, 420, 69, 13, 5}
}

func testSearch[T constraints.Ordered](
	target T,
	list []T,
	fn func(T, []T) bool,
) {
	fmt.Printf("Searching for: %v\n", target)
	if fn(target, list) {
		fmt.Println("Item Found")
	} else {
		fmt.Println("Item not Found")
	}
	fmt.Println("===================")
}

func main() {
	// Testing linearSearch
	fmt.Println("Testing linear search")
	testSearch("doing", createListOfStrings(), linearSearch)
	testSearch("world", createListOfStrings(), linearSearch)
	testSearch(69, createListOfIntegers(), linearSearch)
	testSearch(7, createListOfIntegers(), linearSearch)

	fmt.Println()

	// Testing binarySearch

	stringList := createListOfStrings()
	sort.Strings(stringList)

	intList := createListOfIntegers()
	sort.Ints(intList)

	fmt.Println("Testing iterating binary search")
	testSearch("doing", stringList, iterBinarySearch)
	testSearch("world", stringList, iterBinarySearch)
	testSearch(69, intList, iterBinarySearch)
	testSearch(7, intList, iterBinarySearch)
	testSearch(420, intList, iterBinarySearch)
	testSearch(666, intList, iterBinarySearch)
	testSearch(13, intList, iterBinarySearch)

	fmt.Println()

	fmt.Println("Testing recursive binary search")
	testSearch(69, intList, recBinarySearch)
	testSearch(7, intList, recBinarySearch)
	testSearch(420, intList, recBinarySearch)
	testSearch(666, intList, recBinarySearch)
	testSearch(13, intList, recBinarySearch)
}
