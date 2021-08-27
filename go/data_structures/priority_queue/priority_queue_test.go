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
		binaryHeap []pq.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Normal conditions",
			fields: fields{
				binaryHeap: []pq.Item{
					pq.NewItem(nil, 1),
					pq.NewItem(nil, 4),
					pq.NewItem(nil, 7),
					pq.NewItem(nil, 9),
					pq.NewItem(nil, 2),
				},
			},
			want: 5,
		},
		{
			name: "Empty queue",
			fields: fields{
				binaryHeap: []pq.Item{
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pq.NewPriorityQueue()
			for _, item := range tt.fields.binaryHeap{
				p.Push(item)
			}
			if got := p.Length(); got != tt.want {
				t.Errorf("priorityQueue.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Push(t *testing.T) {
	type args struct {
		i pq.Item
	}
	tests := []struct {
		name   string
		args   args
	}{
		{
			name: "Normal conditions",
			args: args{
				pq.NewItem(nil, 1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pq.NewPriorityQueue()
			lengthBefore := p.Length()
			p.Push(tt.args.i)
			lengthAfter := p.Length()
			if lengthAfter != lengthBefore + 1{
				t.Errorf("priorityQueue.Push() did not insert %v", tt.args.i)
			}
		})
	}
}

func Test_priorityQueue_Pop(t *testing.T) {
	type fields struct {
		binaryHeap []pq.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   pq.Item
	}{
		{
			name: "Normal conditions",
			fields: fields{
				binaryHeap: []pq.Item{
					pq.NewItem(nil, 1),
					pq.NewItem(nil, 4),
					pq.NewItem(nil, 7),
					pq.NewItem(nil, 9),
					pq.NewItem(nil, 2),
				},
			},
			want: pq.NewItem(nil, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pq.NewPriorityQueue()
			for _, item := range tt.fields.binaryHeap{
				p.Push(item)
			}
			if got := p.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("priorityQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priorityQueue_Peek(t *testing.T) {
	type fields struct {
		binaryHeap []pq.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   pq.Item
	}{
		{
			name: "Normal conditions",
			fields: fields{
				binaryHeap: []pq.Item{
					pq.NewItem(nil, 1),
					pq.NewItem(nil, 4),
					pq.NewItem(nil, 7),
					pq.NewItem(nil, 9),
					pq.NewItem(nil, 2),
				},
			},
			want: pq.NewItem(nil, 9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pq.NewPriorityQueue()
			for _, item := range tt.fields.binaryHeap{
				p.Push(item)
			}
			originalLength := p.Length()
			if got := p.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("priorityQueue.Peek() = %v, want %v", got, tt.want)
			}
			currentLength := p.Length()
			if originalLength != currentLength{
				t.Errorf("priorityQueue.Peek() length was altered, was %v, is now %v", originalLength, currentLength)
			}
		})
	}
}

func Test_priorityQueue_Remove(t *testing.T) {
	type fields struct {
		binaryHeap []pq.Item
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Normal conditions",
			fields: fields{
				binaryHeap: []pq.Item{
					pq.NewItem(nil, 1),
					pq.NewItem(nil, 4),
					pq.NewItem(nil, 7),
					pq.NewItem(nil, 9),
					pq.NewItem(nil, 2),
				},
			},
			args: args{
				i: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pq.NewPriorityQueue()
			for _, item := range tt.fields.binaryHeap{
				p.Push(item)
			}
			originalLength := p.Length()
			p.Remove(tt.args.i)
			if originalLength == 0 {
				return
			}
			currentLength := p.Length()
			if originalLength - 1 != currentLength{
				t.Errorf("priorityQueue.Remove() length incorrect, was %v, is now %v", originalLength, currentLength)
			}
		})
	}
}
