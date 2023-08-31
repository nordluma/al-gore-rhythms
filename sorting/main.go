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
	fmt.Println()
}

func testSelectionSort() {
	fmt.Println("Testing Selection Sort")
	listOfStrings := createListOfStrings()
	listOfInts := createListOfIntegers()
	testSort(listOfStrings, selectionSort)
	testSort(listOfInts, selectionSort)
	fmt.Println()
}

func testQuickSort() {
	fmt.Println("Testing Quick Sort")
	listOfStrings := createListOfStrings()
	listOfInts := createListOfIntegers()
	testSort(listOfStrings, quickSort)
	testSort(listOfInts, quickSort)
	fmt.Println()
}

func testBubbleSort() {
	fmt.Println("Testing Bubble Sort")
	listOfStrings := createListOfStrings()
	listOfInts := createListOfIntegers()
	testSort(listOfStrings, bubbleSort)
	testSort(listOfInts, bubbleSort)
	fmt.Println()
}

func testHeapSort() {
	fmt.Println("Testing Heap Sort")
	listOfStrings := createListOfStrings()
	listOfInts := createListOfIntegers()
	testSort(listOfStrings, heapSort)
	testSort(listOfInts, heapSort)
	fmt.Println()
}

func testSort[T constraints.Ordered](
	list []T,
	fn func(*[]T),
) {
	fn(&list)
	fmt.Println(list)
}

func main() {
	testInsertionSort()
	testSelectionSort()
	testQuickSort()
	testBubbleSort()
	testHeapSort()
}
