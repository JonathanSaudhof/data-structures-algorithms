package arrays_test

import (
	"testing"

	"portfolio/das/arrays"
)

const BUFFER_SIZE = 5

// fill a ring buffer with 5 elements
func TestRingBuffer_FillUp_ShouldBeFilled(t *testing.T) {
	ringBuffer := createTestBuffer()

	t.Run("Should fill buffer", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		if ringBuffer.IsEmpty() == false {
			t.Fatalf("Expected ring buffer to be empty")
		}

		fillBuffer(ringBuffer, list, t)

		if ringBuffer.IsFull() == false {
			t.Fatalf("Expected ring buffer to be full")
		}

		if ringBuffer.IsEmpty() == true {
			t.Fatalf("Expected ring buffer to NOT be empty")
		}
	})
}

func TestRingBuffer_Put_ShouldOverwrite(t *testing.T) {
	ringBuffer := createTestBuffer()

	t.Run("Should Overwrite", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		fillBuffer(ringBuffer, list, t)

		// overwrite the first element
		ringBuffer.Put(6)

		// check if the first element was overwritten
		if ringBuffer.Peek() != 2 {
			t.Fatalf("Expected 2, got %v", ringBuffer.Peek())
		}

		if ringBuffer.IsFull() == false {
			t.Fatalf("Expected ring buffer to be full")
		}
	})
}

func TestRingBuffer_Pop_ShouldRemoveElement(t *testing.T) {
	ringBuffer := createTestBuffer()

	t.Run("Should pop", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		fillBuffer(ringBuffer, list, t)

		// remove the first element
		element := ringBuffer.Pop()

		if element != 1 {
			t.Fatalf("Expected 1, got %v", element)
		}

		if ringBuffer.Length() != 4 {
			t.Fatalf("Expected 4, got %v", ringBuffer.Length())
		}

		if ringBuffer.IsFull() == true {
			t.Fatalf("Expected ring buffer not to be full")
		}
	})
}

func TestRingBuffer_Flush_ShouldReturnAllValuesAndReset(t *testing.T) {
	ringBuffer := createTestBuffer()

	t.Run("Should pop", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		fillBuffer(ringBuffer, list, t)

		// remove the first element
		elements := ringBuffer.Flush()

		for idx, element := range elements {

			expectedValue := list[idx]

			if element != expectedValue {
				t.Fatalf("Expected value %v, got %v", element, expectedValue)
			}

			if len(elements) != 5 {
				t.Fatalf("Expected 5, got %v", len(elements))
			}

			if ringBuffer.Length() != 0 {
				t.Fatalf("Expected Length to be 0, got %v", ringBuffer.Length())
			}

			if ringBuffer.IsEmpty() == false {
				t.Fatalf("Expected ring buffer to be empty")
			}

			if ringBuffer.IsFull() == true {
				t.Fatalf("Expected ring buffer not to be full")
			}

		}
	})
}

func TestRingBuffer_FlushNoFullBuffer_ShouldReturnAllValuesAndReset(t *testing.T) {
	ringBuffer := createTestBuffer()

	t.Run("Should pop", func(t *testing.T) {
		list := []int{1, 2, 3}

		for _, value := range list {
			ringBuffer.Put(value)
		}

		if ringBuffer.IsFull() == true {
			t.Fatalf("Expected ring buffer not to be full")
		}

		if ringBuffer.Capacity() != 2 {
			t.Fatalf("Expected 2, got %v", ringBuffer.Capacity())
		}

		if ringBuffer.Length() != 3 {
			t.Fatalf("Expected 3, got %v", ringBuffer.Length())
		}

		// remove the first element
		elements := ringBuffer.Flush()

		for idx, element := range elements {

			expectedValue := list[idx]

			if element != expectedValue {
				t.Fatalf("Expected value %v, got %v", element, expectedValue)
			}
		}

		if ringBuffer.Length() != 0 {
			t.Fatalf("Expected 0, got %v", ringBuffer.Length())
		}

		if ringBuffer.IsEmpty() == false {
			t.Fatalf("Expected ring buffer to be empty")
		}
	})
}

///////////////////////////////////////////////////////////////////////

func fillBuffer(r *arrays.RingBuffer[int], list []int, t *testing.T) {
	t.Run("FillBuffer", func(t *testing.T) {
		for idx, value := range list {
			r.Put(value)

			capacity := r.Capacity()
			expectedCapacity := 5 - (idx + 1)

			if capacity != expectedCapacity {
				t.Fatalf("Expected %v, got %v", expectedCapacity, capacity)
			}
		}

		if r.Length() != 5 {
			t.Fatalf("Expected 5, got %v", r.Length())
		}

		if r.IsEmpty() == true {
			t.Fatalf("Expected ring buffer to NOT be empty")
		}

		if r.IsFull() == false {
			t.Fatalf("Expected ring buffer to be full")
		}
	})
}

func createTestBuffer() *arrays.RingBuffer[int] {
	return arrays.NewRingBuffer[int](BUFFER_SIZE)
}
