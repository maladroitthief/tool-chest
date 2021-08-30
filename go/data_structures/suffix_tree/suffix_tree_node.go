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
