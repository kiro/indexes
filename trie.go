package indexes

type Trie struct {
	root *Node
}

type Node struct {
	isEnd bool
	next []*Node
}

func NewNode() *Node {
	return &Node{
		isEnd: false,
		next: make([]*Node, 256),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (trie *Trie) traverse(value string, addIfMissing bool) *Node {
	current := trie.root
	for _, c := range value {
		if current.next[c] == nil {
			if addIfMissing {
				current.next[c] = NewNode()
			} else {
				return nil
			}
		}

		current = current.next[c]
 	}
 	return current
}

func (trie *Trie) Add(value string) {
	node := trie.traverse(value, true)
	node.isEnd = true
}

func (trie *Trie) Remove(value string) {
	node := trie.traverse(value, false)
	if node != nil {
		node.isEnd = false
	}
}

func (trie *Trie) Contains(value string) bool {
	node := trie.traverse(value, false)
	return node != nil && node.isEnd
}

func dfs(node *Node, limit int, value string) []string {
	result := make([]string, 0)

	if node == nil || limit == 0 {
		return result
	}
	
	if node.isEnd {
		result = append(result, value)
	}
	
	for c, next := range node.next {
		if next != nil {
			values := dfs(next, limit - len(result), value + string([]byte{byte(c)}))
			result = append(result, values...)
		}
	}

	return result
}

func (trie *Trie) Next(prefix string, limit int) []string {
	node := trie.traverse(prefix, false)
	return dfs(node, limit, prefix)
}