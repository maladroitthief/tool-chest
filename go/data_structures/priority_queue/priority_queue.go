package priority_queue

type PriorityQueue interface{
	Push(int)
	Pop() int
	Peek() int
	Remove(int)
}

type priorityQueue struct {
	binaryHeap []int
}

func (p *priorityQueue) Push(element int){
	p.binaryHeap = append(p.binaryHeap, element)
	p.shiftUpNode(len(p.binaryHeap)-1)
}

func (p *priorityQueue) Pop() int{
	result := p.binaryHeap[0]
	// replace root node with last node
	p.binaryHeap[0] = p.binaryHeap[len(p.binaryHeap)-1]
	// trim off the last leaf
	p.binaryHeap = p.binaryHeap[:len(p.binaryHeap)-1]
	// rebalance
	p.shiftDownNode(0)

	return result
}

func (p *priorityQueue) Peek() int {
	return p.binaryHeap[0]
}

func (p *priorityQueue) Remove(i int){
	p.binaryHeap[i] = p.Peek() + 1
	p.shiftUpNode(i)
	p.Pop()
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}

func getLeftChildIndex(i int) int {
	return (2 * i) + 1
}

func getRightChildIndex(i int) int {
	return (2 * i) + 2
}

func (p *priorityQueue) swapNodes(i, j int) {
	p.binaryHeap[i], p.binaryHeap[j] = p.binaryHeap[j], p.binaryHeap[i]
}

func (p *priorityQueue) shiftUpNode(i int){
	for i > 0 && p.binaryHeap[getParentIndex(i)] < p.binaryHeap[i]{
		// swap parent and node
		p.swapNodes(getParentIndex(i), i)
		i = getParentIndex(i)
	}
}

func (p *priorityQueue) shiftDownNode(i int){
	maxIndex := i
	leftChildIndex := getLeftChildIndex(i)

	if leftChildIndex <= len(p.binaryHeap) && p.binaryHeap[leftChildIndex] > p.binaryHeap[maxIndex]{
		maxIndex = leftChildIndex
	}

	rightChildIndex := getRightChildIndex(i)
	if rightChildIndex <= len(p.binaryHeap) && p.binaryHeap[rightChildIndex] > p.binaryHeap[maxIndex]{
		maxIndex = rightChildIndex
	}

	if i != maxIndex{
		p.swapNodes(i, maxIndex)
		p.shiftDownNode(maxIndex)
	}
}
