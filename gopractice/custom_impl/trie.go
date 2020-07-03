package custom

const (
	//AlphabetSize total characters in english alphabet
	AlphabetSize = 26
)

// TrieNode is a node of trie
type TrieNode struct {
	Childrens [AlphabetSize]*TrieNode
	IsWordEnd bool
}

// Trie is impl of trie
type Trie struct {
	Root *TrieNode
}

// InitTrie initializes a empty trie
func InitTrie() *Trie {
	return &Trie{
		Root: &TrieNode{},
	}
}

// Insert inserts a new word
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	current := t.Root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.Childrens[index] == nil {
			current.Childrens[index] = &TrieNode{}
		}
		current = current.Childrens[index]
	}
	current.IsWordEnd = true
}

// Find seraches for a valid word in trie
func (t *Trie) Find(word string) bool {
	wordLength := len(word)
	current := t.Root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.Childrens[index] == nil {
			return false
		}
		current = current.Childrens[index]
	}
	if current.IsWordEnd {
		return true
	}
	return false
}
