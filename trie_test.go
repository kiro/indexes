package indexes 

import(
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	for i := 500; i <= 15000; i++ {
		trie.Add(fmt.Sprintf("%d", i))
	}

	for i := 500; i <= 15000; i++ {
		assert.True(t, trie.Contains(fmt.Sprintf("%d", i)), fmt.Sprintf("contains(%d)", i))
	}

	for i := -500; i < 500; i++ {
		assert.False(t, trie.Contains(fmt.Sprintf("%d", i)), fmt.Sprintf("contains(%d)", i))
	}

	for i := 15001; i < 20000; i++ {
		assert.False(t, trie.Contains(fmt.Sprintf("%d", i)), fmt.Sprintf("contains(%d)", i))
	}

	assert.True(t, trie.Contains("800"), "contains(800)")
	trie.Remove("800")
	assert.False(t, trie.Contains("800"), "contains(800)")

	assert.Equal(t, trie.Next("10", 6), []string{"1000", "10000", "10001", "10002", "10003", "10004"})	
	trie.Remove("10001")
	trie.Remove("10002")
	trie.Remove("10003")
	trie.Remove("10004")
	assert.Equal(t, trie.Next("10", 6), []string{"1000", "10000", "10005", "10006", "10007", "10008"})	
}