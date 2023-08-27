package main

import "fmt"

func createListOfStrings() []string {
	return []string{"d", "k", "a", "j", "b"}
}

func createListOfIntegers() []int {
	return []int{5, 8, 4, 6, 1, 9, 2}
}

func testInsertionSort() {
	fmt.Println("Testing Insertion Sort")
	listOfStrings := createListOfStrings()
	insertionSort(&listOfStrings)
	fmt.Println(listOfStrings)

	fmt.Println("======================")

	listOfInts := createListOfIntegers()
	insertionSort(&listOfInts)
	fmt.Println(listOfInts)
}

func testSelectionSort() {
	fmt.Println("Testing Selection Sort")
	listOfStrings := createListOfStrings()
	selectionSort(&listOfStrings)
	fmt.Println(listOfStrings)

	fmt.Println("======================")

	listOfInts := createListOfIntegers()
	selectionSort(&listOfInts)
	fmt.Println(listOfInts)
}

func main() {
	testInsertionSort()
	fmt.Println()
	testSelectionSort()
}
