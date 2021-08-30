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
		st.activeEdge += currentNode.edgeLength()
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
		if st.activeLength == 0 {
			st.activeEdge = position
		}

		if st.activeNode.childAtIndex(st.activeEdge) {
			next := st.activeNode.children[st.activeEdge]
			if st.walkDown(next) {
				continue
			}

			if st.inputString[next.start+st.activeLength] == st.inputString[position] {
				if st.lastNewNode != nil && st.activeNode != st.root {
					st.lastNewNode.suffixLink = st.activeNode
					st.lastNewNode = nil
				}

				st.activeLength++

				break
			}

			splitEnd := next.start + st.activeLength - 1
			split := st.newNode(next.start, &splitEnd)
			st.activeNode.insertChildAtIndex(st.activeEdge, split)

			if st.lastNewNode != nil {
				st.lastNewNode.suffixLink = split
			}

			st.lastNewNode = split
		} else {
			stn := st.newNode(position, &leafEnd)
			st.activeNode.insertChildAtIndex(st.activeEdge, stn)

			if st.lastNewNode != nil {
				st.lastNewNode.suffixLink = st.activeNode
				st.lastNewNode = nil
			}
		}

		st.remainingSuffixCount--
		if st.activeNode == st.root && st.activeLength > 0 {
			st.activeLength--
			st.activeEdge = position - st.remainingSuffixCount + 1
		} else if st.activeNode != st.root {
			st.activeNode = st.activeNode.suffixLink
		}
	}
}
