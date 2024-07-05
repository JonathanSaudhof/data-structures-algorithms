package lists

import (
	"testing"
	"time"
)

type QueueTestType struct {
	name string
	age  time.Time
}

func TestEnqueueAndDequeue(t *testing.T) {

	var queue = NewQueue[int]()

	t.Run("Queue Should Enqueue and Dequeue correctly", func(t *testing.T) {

		entries := [5]int{1, 2, 3, 4, 5}
		for i, v := range entries {
			queue.Enqueue(v)
			expectedQueueSize := i + 1
			if queue.size != expectedQueueSize {
				t.Errorf("size of queue does not match: %v expected, %v actual", expectedQueueSize, queue.size)
			}
		}

		var dequeuedValues [5]int
		for i := range entries {
			dequeuedValues[i] = queue.Dequeue()
		}

		if entries != dequeuedValues {
			t.Errorf("arrays are not equal, %v expected, %v", entries, dequeuedValues)
		}

	})

}

func TestPeek(t *testing.T) {

	t.Run("Queue Peek should return correct value", func(t *testing.T) {
		var queue = NewQueue[int]()
		queue.Enqueue(1)
		queue.Enqueue(2)

		if queue.Peek() != 2 {
			t.Errorf("should show %v but shows %v instead", 2, queue.Peek())
		}

		if queue.size != 2 {
			t.Errorf("should be size of %v but is %v instead", 2, queue.size)
		}
	})

}
