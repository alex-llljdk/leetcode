package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	children [26]*Trie
	v        string
}

func newTrie() *Trie {
	return &Trie{}
}
func (this *Trie) insert(word string) {
	node := this
	for _, c := range word {
		c -= 'a'
		if node.children[c] == nil {
			node.children[c] = newTrie()
		}
		node = node.children[c]
	}
	node.v = word
}

func (this *Trie) search(word string) string {
	node := this
	for _, c := range word {
		c -= 'a'
		if node.children[c] == nil {
			break
		}
		node = node.children[c]
		if node.v != "" {
			return node.v
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	trie := newTrie()
	for _, v := range dictionary {
		trie.insert(v)
	}
	var ans []string
	for _, v := range strings.Split(sentence, " ") {
		ans = append(ans, trie.search(v))
	}
	return strings.Join(ans, " ")
}

// 下面是非结构化
type prefixTree struct {
	children [26]*prefixTree
	isEnd    bool
	word     string
}

func replaceWords(dictionary []string, sentence string) string {
	pt := prefixTree{}
	for _, word := range dictionary {
		var node *prefixTree = &pt
		for _, c := range word {
			idx := c - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &prefixTree{}
			}
			node = node.children[idx]
		}
		node.isEnd = true
		node.word = word
	}
	sp := strings.Split(sentence, " ")
	for i, word := range sp {
		var node *prefixTree = &pt
		for _, c := range word {
			idx := c - 'a'
			if node.children[idx] == nil {
				break
			}
			node = node.children[idx]
			if node.isEnd {
				sp[i] = node.word
				break
			}
		}
	}
	fmt.Println(strings.Join(sp, " "))
	return strings.Join(sp, " ")
}

func main() {

}
