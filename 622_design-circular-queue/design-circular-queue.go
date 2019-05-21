package design_circular_queue

// MyCircularQueue ...
type MyCircularQueue struct {
	data []int
	cap  int
	f    int
	r    int
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.r] = value
	this.r = (this.r + 1)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.f = this.f + 1
	return true
}

func (this *MyCircularQueue) Front() int {
	return this.data[this.f+1]
}

func (this *MyCircularQueue) Rear() int {
	return this.data[this.r-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.r-this.f == 1
}

func (this *MyCircularQueue) IsFull() bool {
	return this.r == this.cap
}

// Constructor ...
func Constructor(cap int) *MyCircularQueue {
	return &MyCircularQueue{data: make([]int, cap), cap: cap, f: -1, r: 0}
}
