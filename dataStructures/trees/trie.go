package trees

const u8 = 255

type TrieNode struct {
	AtEnd    bool
	Children [u8]*TrieNode
}

type Trie struct {
	root *TrieNode
	len  int
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}, len: 0}
}

func (t *Trie) Insert(text string) {
	currentNode := t.root

	for i := 0; i < len(text); i++ {
		charIdx := text[i] - 'a'
		if currentNode.Children[charIdx] == nil {
			currentNode.Children[charIdx] = &TrieNode{}
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
