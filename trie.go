package structx

import "fmt"

type Trie[T any] struct {
	isEnd    bool
	path     rune   // current path(char)
	fullPath string // fullPath from root node
	data     T
	children []*Trie[T]
}

func NewTrie[T any]() *Trie[T] {
	return new(Trie[T])
}

func (t *Trie[T]) Insert(word string, data ...T) {
	var flag bool
	var node = t

	for _, ch := range word {
		flag = false
		// search children
		for _, child := range node.children {
			// match
			if child.path == ch {
				node = child
				flag = true
				break
			}
		}

		// create children
		if !flag {
			newNode := &Trie[T]{path: ch, fullPath: word}
			if len(data) > 0 {
				newNode.data = data[0]
			}
			node.children = append(node.children, newNode)
			node = newNode
		}
	}
	node.isEnd = true
}

func (t *Trie[T]) SearchPrefix(prefix string) *Trie[T] {
	node := t
	for _, ch := range prefix {
		for _, c := range node.children {
			if c.path == ch {
				node = c
				break
			}
		}
	}
	return node
}

func (t *Trie[T]) Search(word string) (v T, ok bool) {
	node := t.SearchPrefix(word)
	if node != nil && node.isEnd {
		v = node.data
		ok = true
	}
	return
}

func (t *Trie[T]) PrintChildren() {
	for _, c := range t.children {
		fmt.Printf("fullPath: %s, path: %c, isEnd: %v\n", c.fullPath, c.path, c.isEnd)
	}
}

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
