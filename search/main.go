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

func testLinearSearch() {
	fmt.Println("Testing linear search")
	testSearch("doing", createListOfStrings(), linearSearch)
	testSearch("world", createListOfStrings(), linearSearch)
	testSearch(69, createListOfIntegers(), linearSearch)
	testSearch(7, createListOfIntegers(), linearSearch)
}

func testBinarySearch() {
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

	fmt.Println()

	fmt.Println("Testing search or insert binary search")
	sortedIntList := createListOfIntegers()
	sort.Ints(sortedIntList)
	bsInsert(&sortedIntList, 33)
	fmt.Println(sortedIntList)
}


func testBSTSearch() {
	fmt.Println("Testing BST search")
	stringList := createListOfStrings()
	bst := NewBinaryTree[string]()
	for i := 0; i < len(stringList); i++ {
		bst.Insert(stringList[i])
	}

	target := "mom"
	res := bst.Contains(target)
	fmt.Printf("Tree contains: %s => %v\n", target, res)
	fmt.Printf(
		"Tree contains: %s => %v\n",
		"Skill issue",
		bst.Contains("Skill issue"),
	)

	intList := createListOfIntegers()
	bstNum := NewBinaryTree[int]()
	for i := 0; i < len(intList); i++ {
		bstNum.Insert(intList[i])
	}

	targetNum := 420
	fmt.Printf(
		"Tree contains: %d => %v\n",
		targetNum,
		bstNum.Contains(targetNum),
	)
	fmt.Printf(
		"Tree contains: %d => %v\n",
		54,
		bstNum.Contains(54),
	)
}

func testAVLSearch() {
	fmt.Println("Testing AVL search")
	stringList := createListOfStrings()
	avl := NewAVLTree[string]()
	for i := 0; i < len(stringList); i++ {
		avl.Insert(stringList[i])
	}

	fmt.Printf("AVL tree contains: %s => %v\n", "mom", avl.Contains("mom"))
	fmt.Printf(
		"AVL tree contains: %s => %v\n",
		"Skill issue",
		avl.Contains("Skill issue"),
	)

	numList := createListOfIntegers()
	avlNum := NewAVLTree[int]()
	for i := 0; i < len(numList); i++ {
		avlNum.Insert(numList[i])
	}

	fmt.Printf("AVL tree contains: %d => %v\n", 420, avlNum.Contains(420))
	fmt.Printf("AVL tree contains: %d => %v\n", 54, avlNum.Contains(54))
}

func main() {
	// Testing linearSearch
	testLinearSearch()

	fmt.Println()
	testBinarySearch()

	fmt.Println()
	testBSTSearch()

	fmt.Println()
	testAVLSearch()
}
