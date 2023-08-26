package main

import "fmt"

func createListOfStrings() []string {
	return []string{"d", "k", "a", "j", "b"}
}

func createListOfIntegers() []int {
	return []int{5, 8, 4, 6, 1, 9, 2}
}

func main() {
	listOfStrings := createListOfStrings()
	insertionSort(&listOfStrings)
	fmt.Println(listOfStrings)

	fmt.Println("======================")

	listOfInts := createListOfIntegers()
	insertionSort(&listOfInts)
	fmt.Println(listOfInts)
}
