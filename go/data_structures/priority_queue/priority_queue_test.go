package priority_queue

import "testing"

func Test_priorityQueue_Push(t *testing.T) {
	type fields struct {
		binaryHeap []int
	}
	type args struct {
		element int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &priorityQueue{
				binaryHeap: tt.fields.binaryHeap,
			}
			p.Push(tt.args.element)
		})
	}
}

func Test_priorityQueue_Pop(t *testing.T) {
	type fields struct {
		binaryHeap []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &priorityQueue{
				binaryHeap: tt.fields.binaryHeap,
			}
			if got := p.Pop(); got != tt.want {
				t.Errorf("priorityQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Peek(t *testing.T) {
	type fields struct {
		binaryHeap []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &priorityQueue{
				binaryHeap: tt.fields.binaryHeap,
			}
			if got := p.Peek(); got != tt.want {
				t.Errorf("priorityQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Remove(t *testing.T) {
	type fields struct {
		binaryHeap []int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &priorityQueue{
				binaryHeap: tt.fields.binaryHeap,
			}
			p.Remove(tt.args.i)
		})
	}
}
