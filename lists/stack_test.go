package lists

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	stack := NewStack[int]()

	t.Run("Queue Should Enqueue and Dequeue correctly", func(t *testing.T) {
		entries := [5]int{1, 2, 3, 4, 5}
		for i, v := range entries {
			stack.Push(v)
			expectedQueueSize := i + 1
			if stack.size != expectedQueueSize {
				t.Errorf("size of queue does not match: %v expected, %v actual", expectedQueueSize, stack.size)
			}
		}

		var dequeuedValues [5]int
		for i := range entries {
			value, _ := stack.Pop()
			dequeuedValues[i] = value
		}

		expectedEntries := [5]int{5, 4, 3, 2, 1}
		if expectedEntries != dequeuedValues {
			t.Errorf("arrays are not equal, %v expected, %v", entries, dequeuedValues)
		}

		if stack.size != 0 {
			t.Errorf("stack should be size of 0 but is %v", stack.size)
		}
	})
}

func TestTop(t *testing.T) {
	t.Run("Queue Peek should return correct value", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)

		topElement := stack.Top()
		stackSize := stack.size

		if topElement != 2 {
			t.Errorf("should show %v but shows %v instead", 2, topElement)
		}

		if stackSize != 2 {
			t.Errorf("should be size of %v but is %v instead", 2, stackSize)
		}
	})
}
