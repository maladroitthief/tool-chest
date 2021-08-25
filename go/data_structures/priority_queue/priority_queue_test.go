package priority_queue_test

import (
	"reflect"
	"testing"

	pq "github.com/maladroitthief/tool-chest/v2/go/data_structures/priority_queue"
)

func TestNewPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		want pq.PriorityQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pq.NewPriorityQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPriorityQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Length(t *testing.T) {
	type fields struct {
		binaryHeap []*item
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
			if got := p.Length(); got != tt.want {
				t.Errorf("priorityQueue.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Push(t *testing.T) {
	type fields struct {
		binaryHeap []*item
	}
	type args struct {
		i Item
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
			p.Push(tt.args.i)
		})
	}
}

func Test_priorityQueue_Pop(t *testing.T) {
	type fields struct {
		binaryHeap []*item
	}
	tests := []struct {
		name   string
		fields fields
		want   Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &priorityQueue{
				binaryHeap: tt.fields.binaryHeap,
			}
			if got := p.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("priorityQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Peek(t *testing.T) {
	type fields struct {
		binaryHeap []*item
	}
	tests := []struct {
		name   string
		fields fields
		want   Item
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &priorityQueue{
				binaryHeap: tt.fields.binaryHeap,
			}
			if got := p.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("priorityQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Remove(t *testing.T) {
	type fields struct {
		binaryHeap []*item
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
