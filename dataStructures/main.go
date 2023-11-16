package main

import (
	"ds/trees"
	"fmt"
)

func main() {
	trie := trees.NewTrie()
	trie.Insert("https://www.reddit.com/r/rust")
	trie.Insert("https://www.manning.com")

	if trie.Contains("https://www.reddit.com/r/rust") {
		fmt.Println("Trie contains reddit")
	} else {
		fmt.Println("Something is wrong")
	}

	fmt.Println(trie)
}
