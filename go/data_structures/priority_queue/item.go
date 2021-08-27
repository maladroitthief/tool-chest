package priority_queue

type Item interface {
	GetPriority() int
	GetValue() interface{}
	SetPriority(int)
	SetValue(interface{})
}

type item struct {
	value interface{}
	priority int
}

func NewItem(value interface{}, priority int) Item{
	return &item{
		value: value,
		priority: priority,
	}
}

func (i *item) GetPriority() int {
	return i.priority
}

func (i *item) GetValue() interface{} {
	return i.value
}

func (i *item) SetPriority(p int){
	i.priority = p
}

func (i *item) SetValue(v interface{}){
	i.value = v
}
