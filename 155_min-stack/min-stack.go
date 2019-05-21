package min_stack

type Data struct {
	val int
	min *Data
}

// MinStack ...
type MinStack struct {
	data    []*Data
	cap     int
	pointer int
}

// Constructor ...
func Constructor() MinStack {
	return MinStack{cap: 0, pointer: -1}
}

// Pop ...
func (this *MinStack) Pop() {
	this.data[this.pointer] = nil
	this.pointer--
}

// Top ...
func (this *MinStack) top() *Data {
	return this.data[this.pointer]
}

// Top ...
func (this *MinStack) Top() int {
	return this.top().val
}

// Push ...
func (this *MinStack) Push(x int) {
	if this.size() == this.cap {
		if this.size() == 0 {
			this.data = make([]*Data, 1)
			this.cap = 1
		} else {
			this.data = append(this.data, make([]*Data, this.cap*2)...)
			this.cap *= 2
		}
	}

	node := &Data{val: x}

	if this.pointer == -1 || x < this.GetMin() {
		node.min = node
	} else {
		node.min = this.min()
	}

	this.pointer++
	this.data[this.pointer] = node
}

func (this *MinStack) min() *Data {
	return this.top().min
}

// GetMin ...
func (this *MinStack) GetMin() int {
	return this.min().val
}

func (this *MinStack) size() int {
	return this.pointer + 1
}
