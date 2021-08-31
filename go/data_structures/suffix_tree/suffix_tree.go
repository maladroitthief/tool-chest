package suffix_tree

import "log"

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
	}

	st.rootEnd = new(int)
	*st.rootEnd = -1
	st.root = st.newNode(-1, st.rootEnd)
	st.activeNode = st.root
	for i := range s {
		st.extend(i)
	}

	labelHeight := 0
	st.setSuffixIndexByDFS(st.root, labelHeight)
	st.freeByPostOrder(st.root)

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

func (st *suffixTree) setSuffixIndexByDFS(stn *suffixTreeNode, labelHeight int) {
	if stn == nil {
		return
	}

	if stn.start != -1 {
		for i := stn.start; i <= *stn.end; i++ {
			log.Printf("%s", string(st.inputString[i]))
		}
	}

	isLeaf := true
	for _, child := range stn.children {
		if child != nil {
			if isLeaf && stn.start != -1 {
				log.Printf(" [%v]\n", stn.suffixIndex)
			}
			isLeaf = false
			st.setSuffixIndexByDFS(child, labelHeight+child.edgeLength())
		}
	}

	if isLeaf {
		stn.suffixIndex = st.size() - labelHeight
		log.Printf(" [%v]\n", stn.suffixIndex)
	}
}

func (st *suffixTree) freeByPostOrder(stn *suffixTreeNode) {
	if stn == nil {
		return
	}

	for _, child := range stn.children {
		if child != nil {
			st.freeByPostOrder(child)
		}
	}

	if stn.suffixIndex == -1 {
		stn.end = nil
	}

	stn = nil
}

func (st *suffixTree) size() int {
	return len(st.inputString)
}
