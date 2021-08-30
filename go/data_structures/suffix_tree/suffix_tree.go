package suffix_tree

type SuffixTree interface {
}

type suffixTree struct {
	inputString          string
	root                 *suffixTreeNode
	lastNewNode          *suffixTreeNode
	activeNode           *suffixTreeNode
	nodeCount            int
	activeEdge           int
	activeLength         int
	remainingSuffixCount int
	leafEnd              int
	rootEnd              *int
	splitEnd             *int
	inputStringSize      int
}

func NewSuffixTree(s string) SuffixTree {
	st := suffixTree{
		inputString:          s,
		root:                 nil,
		lastNewNode:          nil,
		activeNode:           nil,
		nodeCount:            0,
		activeEdge:           -1,
		activeLength:         0,
		remainingSuffixCount: 0,
		leafEnd:              -1,
		rootEnd:              nil,
		splitEnd:             nil,
		inputStringSize:      -1,
	}
	return &st
}

func (st *suffixTree) walkDown(currentNode *suffixTreeNode) bool {
	if st.activeLength >= currentNode.edgeLength() {
		st.activeEdge = st.activeEdge - currentNode.edgeLength()
		st.activeLength -= currentNode.edgeLength()
		st.activeNode = currentNode
		return true
	}

	return false
}

func (st *suffixTree) extend(position int) {
	leafEnd := position
	st.remainingSuffixCount++
	st.lastNewNode = nil

	for st.remainingSuffixCount > 0 {

	}

}
