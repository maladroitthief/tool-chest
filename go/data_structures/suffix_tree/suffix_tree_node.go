package suffix_tree

type suffixTreeNode struct {
	children map[rune]*suffixTreeNode
	// pointer to other node via suffix link
	suffixLink *suffixTreeNode
	// (start, end) interval specifies the edge, by which the node is connected to its parent node. Each edge will connect two nodes, one parent and one child, and (start, end) interval of a given edge will be stored in the child node.
	start int
	end   *int
	// for leaf nodes, it stores the index of suffix for the path from root to leaf
	suffixIndex int
	isRoot      bool
}

func (st *suffixTree) newNode(start int, end *int) *suffixTreeNode {
	st.nodeCount++
	stn := suffixTreeNode{
		// For root node, suffixLink will be set to nil.
		// For internal nodes, suffixLink will be set to root by default in current extension and may change in next extension
		suffixLink: st.root,
		start:      start,
		end:        end,
		// suffixIndex will be set to -1 by default and actual suffix index will be set later for leaves at the end of all phases
		suffixIndex: -1,
		isRoot:      false,
	}
	stn.children = make(map[rune]*suffixTreeNode)
	return &stn
}

func (stn *suffixTreeNode) setRoot() {
	stn.isRoot = true
}

func (stn *suffixTreeNode) edgeLength() int {
	if stn.isRoot {
		return 0
	}

	return *stn.end - stn.start + 1
}
