package trees

type TrieNode struct {
	AtEnd    bool
	Children map[byte]*TrieNode
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[byte]*TrieNode),
		AtEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
	len  int
}

func NewTrie() *Trie {
	return &Trie{root: newTrieNode(), len: 0}
}

func (t *Trie) Insert(text string) {
	currentNode := t.root

	for i := 0; i < len(text); i++ {
		charIdx := text[i] - 'a'
		if currentNode.Children[charIdx] == nil {
			currentNode.Children[charIdx] = newTrieNode()
		}

		currentNode = currentNode.Children[charIdx]
	}

	currentNode.AtEnd = true
	t.len++
}

func (t *Trie) Contains(text string) bool {
	currentNode := t.root

	for i := 0; i < len(text); i++ {
		charIdx := text[i] - 'a'
		if currentNode.Children[charIdx] == nil {
			return false
		}
		currentNode = currentNode.Children[charIdx]
	}

	return currentNode.AtEnd
}

func (t *Trie) Length() int {
	return t.len
}
