package design_circular_queue

import (
	"testing"
)

func TestMyCircularQueue(t *testing.T) {
	circularQueue := Constructor(3)     // set the size to be 3
	enqueue := circularQueue.EnQueue(1) // return true
	if !enqueue {
		t.Error("Enqueue Failed")
	}
	enqueue = circularQueue.EnQueue(2) // return true
	if !enqueue {
		t.Error("Enqueue Failed")
	}
	enqueue = circularQueue.EnQueue(3) // return true
	if !enqueue {
		t.Error("Enqueue Failed")
	}
	enqueue = circularQueue.EnQueue(4)
	if enqueue {
		t.Error("Enqueue Failed")
	}

	rear := circularQueue.Rear()
	if rear != 3 {
		t.Error("Rear Failed")
	}

	full := circularQueue.IsFull()
	if !full {
		t.Error("isFull Failed")
	}

	dequeue := circularQueue.DeQueue()
	if !dequeue {
		t.Error("DeQueue Failed")
	}

	dequeue = circularQueue.DeQueue()
	if !dequeue {
		t.Error("DeQueue Failed")
	}

	dequeue = circularQueue.DeQueue()
	if !dequeue {
		t.Error("DeQueue Failed")
	}

	dequeue = circularQueue.DeQueue()
	if dequeue {
		t.Error("DeQueue Failed")
	}

	dequeue = circularQueue.DeQueue()
	if dequeue {
		t.Error("DeQueue Failed")
	}

	enqueue = circularQueue.EnQueue(4)
	if !enqueue {
		t.Error("Enqueue Failed")
	}

	rear = circularQueue.Rear()
	if rear != 4 {
		t.Error("Rear Failed")
	}
}
