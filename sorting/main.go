package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func createListOfStrings() []string {
	return []string{"d", "k", "a", "j", "b"}
}

func createListOfIntegers() []int {
	return []int{5, 8, 4, 6, 1, 9, 2}
}

func testInsertionSort() {
	fmt.Println("Testing Insertion Sort")
	listOfStrings := createListOfStrings()
	listOfInts := createListOfIntegers()
	testSort(listOfStrings, insertionSort)
	testSort(listOfInts, insertionSort)
}

func testSelectionSort() {
	fmt.Println("Testing Selection Sort")
	listOfStrings := createListOfStrings()
	listOfInts := createListOfIntegers()
	testSort(listOfStrings, selectionSort)
	testSort(listOfInts, selectionSort)
}

func testQuickSort() {
	fmt.Println("Testing Quick Sort")
	listOfStrings := createListOfStrings()
	listOfInts := createListOfIntegers()
	testSort(listOfStrings, quickSort)
	testSort(listOfInts, quickSort)
}

func testSort[T constraints.Ordered](
	list []T,
	fn func(*[]T),
) {
	fn(&list)
	fmt.Println(list)
	fmt.Println("======================")
}

func main() {
	testInsertionSort()
	fmt.Println()
	testSelectionSort()
	fmt.Println()
	testQuickSort()
}
