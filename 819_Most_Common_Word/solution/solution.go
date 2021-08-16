package solution

/*
Runtime: 4 ms, faster than 45.68% of Go online submissions for Most Common Word.
Memory Usage: 3.8 MB, less than 22.22% of Go online submissions for Most Common Word.
*/

import (
	"regexp"
	"strings"
)

var (
	res    string
	maxcnt int
)

func MostCommonWord(paragraph string, banned []string) string {
	res = ""
	maxcnt = 0
	root := &TrieNode{}
	re := regexp.MustCompile(`\w+`)
	words := re.FindAll([]byte(paragraph), -1)
	for _, w := range words {
		root.insert(string(w))
	}
	for _, b := range banned {
		root.ban(b)
	}
	findMax(root)
	return res
}

type TrieNode struct {
	word  string
	cnt   int
	links [26]*TrieNode
}

func (t *TrieNode) insert(s string) {
	var curr *TrieNode
	curr = t
	ss := strings.ToLower(s)
	for _, r := range ss {
		i := int(r) - int('a')
		if curr.links[i] == nil {
			curr.links[i] = &TrieNode{}
			curr.links[i].word = curr.word + string(r)
		}
		curr = curr.links[i]
	}
	curr.cnt += 1
}
func (t *TrieNode) ban(s string) {
	for _, r := range s {
		i := int(r) - 'a'
		if t.links[i] == nil {
			return
		}
		t = t.links[i]
	}
	t.cnt = 0
}

func findMax(node *TrieNode) string {
	t := node
	if t == nil {
		return ""
	}

	if t.cnt > maxcnt {
		res = t.word
		maxcnt = t.cnt
	}
	for _, n := range t.links {
		if n != nil {
			findMax(n)
		}
	}
	return res
}
