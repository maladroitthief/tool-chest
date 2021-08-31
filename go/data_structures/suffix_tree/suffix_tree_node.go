package suffix_tree

type suffixTreeNode struct {
	children    []*suffixTreeNode
	suffixLink  *suffixTreeNode
	start       int
	end         *int
	suffixIndex int
}

func (st *suffixTree) newNode(start int, end *int) *suffixTreeNode {
	st.nodeCount++
	return &suffixTreeNode{
		suffixLink:  st.root,
		start:       start,
		end:         end,
		suffixIndex: -1,
	}
}

func (stn *suffixTreeNode) edgeLength() int {
	return *stn.end - stn.start + 1
}

func (stn *suffixTreeNode) childAtIndex(i int) bool {
	if len(stn.children) <= i {
		return false
	}

	if stn.children[i] == nil {
		return false
	}

	return true
}

func (stn *suffixTreeNode) insertChildAtIndex(i int, s *suffixTreeNode) {
	for i >= len(stn.children) {
		stn.children = append(stn.children, &suffixTreeNode{})
	}

	stn.children[i] = s
}
