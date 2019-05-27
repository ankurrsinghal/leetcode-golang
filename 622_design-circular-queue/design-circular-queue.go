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
	this.r = (this.r + 1) % this.cap
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.f = (this.f + 1) % this.cap
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.f]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.cap+this.r-1)%this.cap]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.r == this.f
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size() == this.cap-1
}

func (this *MyCircularQueue) size() int {
	return (this.cap - this.f + this.r) % this.cap
}

// Constructor ...
func Constructor(cap int) MyCircularQueue {
	return MyCircularQueue{data: make([]int, cap+1), cap: cap + 1, f: 0, r: 0}
}
