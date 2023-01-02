package structx

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Trie[T any] struct {
	isEnd    bool
	path     rune   // current path
	fullPath string // fullPath from root node
	data     T

	parent   *Trie[T]
	children []*Trie[T]
}

// NewTrie
func NewTrie[T any]() *Trie[T] {
	return new(Trie[T])
}

// Insert:
func (t *Trie[T]) Insert(word string, data ...T) error {
	cur := t
	for index, ch := range word {
		// match
		i := slices.IndexFunc(cur.children, func(t *Trie[T]) bool {
			return t.path == ch
		})

		if i >= 0 {
			cur = cur.children[i]

		} else {
			// add children
			node := &Trie[T]{path: ch, fullPath: word[:index] + string(ch), parent: cur}
			cur.children = append(cur.children, node)
			cur = node
		}
	}
	cur.isEnd = true
	if len(data) > 0 {
		cur.data = data[0]
	}
	return nil
}

// Delete
func (t *Trie[T]) Delete(word string) error {
	// search
	node := t.Search(word)
	if node == nil {
		return fmt.Errorf("word[%s] not exist", word)
	}

	p := node.parent
	i := slices.IndexFunc(p.children, func(t *Trie[T]) bool {
		return t == node
	})

	// delete
	p.children = slices.Delete(p.children, i, i+1)
	return nil
}

// Search
func (t *Trie[T]) Search(prefix string) *Trie[T] {
	node := t
	for _, ch := range prefix {
		// match
		i := slices.IndexFunc(node.children, func(t *Trie[T]) bool {
			return t.path == ch
		})
		if i < 0 {
			return nil
		}
		node = node.children[i]
	}
	return node
}

// GetData
func (t *Trie[T]) GetData() T {
	return t.data
}

// PrintChildren
func (t *Trie[T]) PrintChildren() {
	fmt.Printf("[parent] fullPath: %s, path: %c, isEnd: %v, data: %v\n",
		t.fullPath, t.path, t.isEnd, t.data)
	for _, c := range t.children {
		fmt.Printf("-[child] fullPath: %s, path: %c, isEnd: %v, data: %v\n",
			c.fullPath, c.path, c.isEnd, c.data)
	}
}

// Print
func (t *Trie[T]) Print() {
	var printChildrens func(n *Trie[T])

	printChildrens = func(n *Trie[T]) {
		for _, c := range n.children {
			fmt.Printf("%c", c.path)
			if c.isEnd {
				fmt.Println()
			}
			printChildrens(c)
		}
	}
	printChildrens(t)
}
