/*
假定一种编码的编码范围是a ~ y的25个字母，从1位到4位的编码，如果我们把该编码按字典序排序，形成一个数组如下： a, aa, aaa, aaaa, aaab, aaac, … …, b, ba, baa, baaa, baab, baac … …, yyyw, yyyx, yyyy 其中a的Index为0，aa的Index为1，aaa的Index为2，以此类推。 编写一个函数，输入是任意一个编码，输出这个编码对应的Index.
*/
package main

import "fmt"

type Trie struct {
	chirdren [25]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.chirdren[ch] == nil {
			node.chirdren[ch] = &Trie{}
		}
		node = node.chirdren[ch]
	}
	node.isEnd = true
}
func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.chirdren[ch] == nil {
			return nil
		}
		node = node.chirdren[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func main() {
	tree := &Trie{}
	for i := 'a'; i < 'y'; i++ {
		tree.Insert(string(i))
		for j := 'a'; j < 'y'; j++ {
			tree.Insert(string(i + j))
			for k := 'a'; k < 'y'; k++ {
				tree.Insert(string(i + j + k))
				for n := 'a'; n < 'y'; n++ {
					tree.Insert(string(i + j + k + n))
				}
			}
		}
	}

	fmt.Print(tree.Search("abbu"))
}
