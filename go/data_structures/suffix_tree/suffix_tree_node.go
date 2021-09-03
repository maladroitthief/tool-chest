package suffix_tree

type suffixTreeNode struct {
	children    map[int]*suffixTreeNode
	suffixLink  *suffixTreeNode
	start       int
	end         *int
	suffixIndex int
}

func (st *suffixTree) newNode(start int, end *int) *suffixTreeNode {
	st.nodeCount++
	stn := suffixTreeNode{
		suffixLink:  st.root,
		start:       start,
		end:         end,
		suffixIndex: -1,
	}
	stn.children = make(map[int]*suffixTreeNode)
	return &stn
}

func (stn *suffixTreeNode) edgeLength() int {
	return *stn.end - stn.start + 1
}
