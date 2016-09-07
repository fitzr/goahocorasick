package ahocorasick

type AhoCorasick struct {
	root *node
}

type node struct {
	children map[rune]*node
	depth    int
	next     *node
	hit      bool
}

func New(keywords []string) *AhoCorasick {
	a := AhoCorasick{root: newNode(0)}
	a.createTrie(keywords)
	a.createNext()
	return &a
}

func newNode(depth int) *node {
	return &node{
		children: map[rune]*node{},
		depth:    depth}
}

func (a *AhoCorasick) createTrie(keywords []string) {
	for _, keyword := range keywords {
		n := a.root
		for _, r := range keyword {
			v, ok := n.children[r]
			if !ok {
				v = newNode(n.depth + 1)
				n.children[r] = v
			}
			n = v
		}
		n.hit = true
	}
}

func (a *AhoCorasick) createNext() {
	for k, v := range a.root.children {
		a.walkCreateNext(v, []rune{k})
	}
}

func (a *AhoCorasick) walkCreateNext(n *node, text []rune) {
	n.next = a.backwardMatchNode(text)
	for k, v := range n.children {
		a.walkCreateNext(v, append(text, k))
	}
}

func (a *AhoCorasick) backwardMatchNode(text []rune) *node {
	for t := text[1:]; len(t) > 0; t = t[1:] {
		n, ok := a.matchNode(t)
		if ok {
			return n
		}
	}
	return a.root
}

func (a *AhoCorasick) matchNode(text []rune) (*node, bool) {
	n := a.root
	for _, r := range text {
		v, ok := n.children[r]
		if ok {
			n = v
		} else {
			return nil, false
		}
	}
	return n, true
}

func (a *AhoCorasick) Match(text string) [][]int {
	result := make([][]int, 0)
	runes := []rune(text)
	n := a.root

	for i := 0; i < len(runes); {
		if n.hit {
			result = append(result, []int{i-n.depth, n.depth})
		}
		v, ok := n.children[runes[i]]
		if ok {
			n = v
			i++
		} else if n.next != nil {
			n = n.next
		} else {
			i++
		}
	}

	return result
}
