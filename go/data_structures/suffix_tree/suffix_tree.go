package suffix_tree

import (
	"errors"
	"log"
)

var (
	ErrNoRepeatedSubstrings = errors.New("no repeated substrings")
)

type SuffixTree interface {
	LongestRepeatedSubstring() (string, error)
}

type suffixTree struct {
	inputString         string
	terminatorCharacter string
	root                *suffixTreeNode
	// lastNewNode will point to newly created internal node, waiting for it's suffix link to be set, which might get a new suffix link (other than root) in next extension of same phase. lastNewNode will be set to nil when last newly created internal node (if there is any) got it's suffix link reset to new internal node created in next extension of same phase.
	lastNewNode *suffixTreeNode
	activeNode  *suffixTreeNode
	nodeCount   int
	// activeEdge is represented as input string character index (not the character itself)
	activeEdge   int
	activeLength int
	// remainingSuffixCount tells how many suffixes yet to be added in tree
	remainingSuffixCount int
	leafEnd              int
	rootEnd              *int
	splitEnd             *int
}

func NewSuffixTree(s string) SuffixTree {
	st := suffixTree{
		inputString:          s,
		terminatorCharacter:  "\u0000",
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

	st.inputString += st.terminatorCharacter
	st.rootEnd = new(int)
	*st.rootEnd = -1
	st.root = st.newNode(-1, st.rootEnd)
	st.root.setRoot()
	st.activeNode = st.root
	for i := range st.inputString {
		st.extend(i)
	}

	labelHeight := 0
	st.setSuffixIndexByDFS(st.root, labelHeight)
	// st.freeByPostOrder(st.root)

	return &st
}

func (st *suffixTree) RootTraversal(labelHeight int, maxHeight, substringStartIndex *int) {
	st.root.traversal(labelHeight, maxHeight, substringStartIndex)
}

func (st *suffixTree) LongestRepeatedSubstring() (string, error) {
	maxHeight := 0
	substringStartIndex := 0

	st.RootTraversal(0, &maxHeight, &substringStartIndex)

	// if maxHeight > 0 {
	// 	return st.inputString[substringStartIndex:maxHeight], nil
	// }

	results := ""
	for k := 0; k < maxHeight; k++ {
		results += string(st.inputString[k+substringStartIndex])
	}

	if results == "" {
		return "", ErrNoRepeatedSubstrings
	}

	return results, nil
}

func (st *suffixTree) walkDown(currentNode *suffixTreeNode) bool {
	// activePoint change for walk down
	if st.activeLength >= currentNode.edgeLength() {
		st.activeEdge += currentNode.edgeLength()
		st.activeLength -= currentNode.edgeLength()
		st.activeNode = currentNode
		return true
	}

	return false
}

func (st *suffixTree) extend(position int) {
	// Extension Rule 1
	// If the path from the root labelled S[j..i] ends at leaf edge (i.e. S[i] is last character on leaf edge) then character S[i+1] is just added to the end of the label on that leaf edge.
	st.leafEnd = position
	// Increment remainingSuffixCount
	st.remainingSuffixCount++
	// set lastNewNode to nil when starting a new phase
	st.lastNewNode = nil

	// add all remaining suffixes to the tree, one by one
	for st.remainingSuffixCount > 0 {
		// activePoint change for activeLength zero
		if st.activeLength == 0 {
			st.activeEdge = position
		}

		// There is no outgoing edge starting with activeEdge from activeNode
		if st.activeNode.children[st.activeRune()] == nil {
			// Extension Rule 2:
			// If the path from the root labelled S[j..i] ends at non-leaf edge (i.e. there are more characters after S[i] on path) and next character is not s[i+1], then a new leaf edge with label s{i+1] and number j is created starting from character S[i+1].  A new internal node will also be created if s[1..i] ends inside (in-between) a non-leaf edge.
			st.activeNode.children[st.activeRune()] = st.newNode(position, &st.leafEnd)

			// if there is any internal node waiting for it's suffix link get reset, point the suffix link from that last internal node to current activeNode. Then set lastNewNode to NULL indicating no more node waiting for suffix link reset.
			if st.lastNewNode != nil {
				st.lastNewNode.suffixLink = st.activeNode
				st.lastNewNode = nil
			}
		} else {
			// There is an outgoing edge starting with activeEdge from activeNode
			// Get the next node at the end of edge starting with activeEdge
			next := st.activeNode.children[st.activeRune()]
			if st.walkDown(next) {
				continue
			}

			// Extension rule 3
			// If the path from the root labelled S[j..i] ends at non-leaf edge (i.e. there are more characters after S[i] on path) and next character is s[i+1] (already in tree), do nothing.
			if st.inputString[next.start+st.activeLength] == st.inputString[position] {
				// If a newly created node waiting for it's suffix link to be set, then set suffix link of that waiting node to current active node
				if st.lastNewNode != nil && !st.activeNode.isRoot {
					st.lastNewNode.suffixLink = st.activeNode
					st.lastNewNode = nil
				}
				// activePoint change for extension rule 3
				st.activeLength++
				// stop all further processing in this phase and move on to next phase
				break
			}

			// We will be here when activePoint is in middle of the edge being traversed and current character being processed is not on the edge (we fall off the tree). In this case, we add a new internal node and a new leaf edge going out of that new node. This is Extension Rule 2, where a new leaf edge and a new internal node get created
			st.splitEnd = new(int)
			*st.splitEnd = next.start + st.activeLength - 1

			// new internal node
			split := st.newNode(next.start, st.splitEnd)
			st.activeNode.children[st.activeRune()] = split

			// new leaf coming off of internal node
			split.children[rune(st.inputString[position])] = st.newNode(position, &st.leafEnd)
			next.start += st.activeLength
			split.children[rune(st.inputString[next.start])] = next

			// We got a new internal node here. If there is any internal node created in last extensions of same phase which is still waiting for it's suffix link reset, do it now.
			if st.lastNewNode != nil {
				// suffixLink of lastNewNode points to current newly created internal node
				st.lastNewNode.suffixLink = split
			}

			// Make the current newly created internal node waiting for it's suffix link reset (which is pointing to root at present). If we come across any other internal node (existing or newly created) in next extension of same phase, when a new leaf edge gets added (i.e. when Extension Rule 2 applies is any of the next extension of same phase) at that point, suffixLink of this node will point to that internal node.
			st.lastNewNode = split
		}

		// One suffix got added in tree, decrement the count of suffixes yet to be added.
		st.remainingSuffixCount--
		// activeNode is root and activeLength is greater than zero
		if st.activeNode.isRoot && st.activeLength > 0 {
			st.activeLength--
			st.activeEdge = position - st.remainingSuffixCount + 1
		} else if !st.activeNode.isRoot {
			// If activeNode is not root, then follow the suffix link from current activeNode
			st.activeNode = st.activeNode.suffixLink
		}
	}
}

func (st *suffixTree) setSuffixIndexByDFS(stn *suffixTreeNode, labelHeight int) {
	if stn == nil {
		return
	}

	if stn.start != -1 {
		log.Print(st.inputString[stn.start : *(stn.end)+1])
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

func (st *suffixTree) activeRune() rune {
	return rune(st.inputString[st.activeEdge])
}
