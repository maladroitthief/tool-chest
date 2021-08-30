package priority_queue

import "errors"

// PriorityQueue describes all the the functions that a priorityQueue object must implement
type PriorityQueue interface {
	Push(Item)
	Pop() (Item, error)
	Peek() (Item, error)
	Remove(int) error
	Length() int
}

type priorityQueue struct {
	binaryHeap []*item
}

var (
	ErrPriorityQueueEmpty = errors.New("priority queue is empty")
	ErrIndexOutOfBounds   = errors.New("index is out of bounds of the queue")
)

// NewPriorityQueue returns an instance of a priorityQueue struct
func NewPriorityQueue() PriorityQueue {
	return &priorityQueue{}
}

// Length returns the current length of the priorityQueue
func (p *priorityQueue) Length() int {
	return len(p.binaryHeap)
}

// Push inserts an Item into the priorityQueue and shifts it into position based
// on the Item's priority
func (p *priorityQueue) Push(i Item) {
	item := item{
		value:    i.GetValue(),
		priority: i.GetPriority(),
	}
	p.binaryHeap = append(p.binaryHeap, &item)
	p.shiftUpNode(p.Length() - 1)
}

// Pop returns the Item with the highest priority and removes it from the
// priorityQueue. An error is returned for any illegal operations

func (p *priorityQueue) Pop() (Item, error) {
	if p.Length() == 0 {
		return nil, ErrPriorityQueueEmpty
	}

	result := p.binaryHeap[0]
	// replace root node with last node
	p.binaryHeap[0] = p.binaryHeap[p.Length()-1]
	// Avoid memory leaks
	p.binaryHeap[p.Length()-1] = nil
	// trim off the last leaf
	p.binaryHeap = p.binaryHeap[:p.Length()-1]
	// rebalance
	p.shiftDownNode(0)

	return result, nil
}

// Peek returns the highest priority Item without removing it from the queue
func (p *priorityQueue) Peek() (Item, error) {
	if p.Length() == 0 {
		return nil, ErrPriorityQueueEmpty
	}

	return p.binaryHeap[0], nil
}

// Remove removes the Item located at index i from the priorityQueue
func (p *priorityQueue) Remove(i int) error {
	if i < 0 || i >= p.Length() {
		return ErrIndexOutOfBounds
	}

	currentRoot, err := p.Peek()
	if err != nil {
		return err
	}

	p.binaryHeap[i].SetPriority(currentRoot.GetPriority() + 1)
	p.shiftUpNode(i)
	_, err = p.Pop()

	return err
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return (2 * index) + 1
}

func getRightChildIndex(index int) int {
	return (2 * index) + 2
}

func (p *priorityQueue) swapNodes(i, j int) {
	p.binaryHeap[i], p.binaryHeap[j] = p.binaryHeap[j], p.binaryHeap[i]
}

func (p *priorityQueue) shiftUpNode(i int) {
	// Check that index is not 0 and the that parent is lower priority
	for i > 0 && p.binaryHeap[getParentIndex(i)].GetPriority() < p.binaryHeap[i].GetPriority() {
		// swap parent and node
		p.swapNodes(getParentIndex(i), i)
		i = getParentIndex(i)
	}
}

func (p *priorityQueue) shiftDownNode(i int) {
	maxIndex := i
	leftChildIndex := getLeftChildIndex(i)

	if leftChildIndex < p.Length() && p.binaryHeap[leftChildIndex].GetPriority() > p.binaryHeap[maxIndex].GetPriority() {
		maxIndex = leftChildIndex
	}

	rightChildIndex := getRightChildIndex(i)
	if rightChildIndex < p.Length() && p.binaryHeap[rightChildIndex].GetPriority() > p.binaryHeap[maxIndex].GetPriority() {
		maxIndex = rightChildIndex
	}

	if i != maxIndex {
		p.swapNodes(i, maxIndex)
		p.shiftDownNode(maxIndex)
	}
}
