package lists

import (
	"testing"
)

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
			value, err := queue.Dequeue()
			dequeuedValues[i] = value

			if err != nil {
				t.Errorf("dequeue should not throw an error, Error: %v", err)
			}
		}

		if entries != dequeuedValues {
			t.Errorf("arrays are not equal, %v expected, %v", entries, dequeuedValues)
		}

	})

}

func TestEnqueueAndDequeue_DequeueEmptyQueue_ShouldReturnError(t *testing.T) {

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
			value, err := queue.Dequeue()
			if err != nil {
				t.Errorf("Should not return an error; Error %v", err)
			}
			dequeuedValues[i] = value
		}

		if entries != dequeuedValues {
			t.Errorf("arrays are not equal, %v expected, %v", entries, dequeuedValues)
		}

		if queue.size != 0 {
			t.Errorf("queue should be size of 0 but is %v instead", queue.size)
		}

		_, err := queue.Dequeue()
		if err == nil {
			t.Errorf("Dequeue should throw an error")
		}
	})

}

func TestPeek(t *testing.T) {

	t.Run("Queue Peek should return correct value", func(t *testing.T) {
		var queue = NewQueue[int]()

		value, err := queue.Peek()

		if err == nil {
			t.Errorf("Peek should return an error")
		}

		if value != 0 {
			t.Errorf("Peek should return a default value but returns %v", value)
		}

		queue.Enqueue(1)
		queue.Enqueue(2)

		value, err = queue.Peek()

		if err != nil {
			t.Errorf("Peek should not return an error; Error: %v", err)
		}

		if value != 1 {
			t.Errorf("should show 1 but shows %v instead", value)
		}

		if queue.size != 2 {
			t.Errorf("should be size of 2 but is %v instead", queue.size)
		}
	})

}
