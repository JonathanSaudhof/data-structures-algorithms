package arrays

import (
	"testing"
)

// fill a ring buffer with 5 elements
func TestRingBuffer(t *testing.T) {
	ringBuffer := NewRingBuffer[int](5)

	t.Run("Capacity", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}

		for idx, value := range list {
			ringBuffer.Put(value)
			if *ringBuffer.buf[idx] != value {
				t.Fatalf("Expected %v, got %v", &value, ringBuffer.buf[idx])
			}

			capacity := ringBuffer.Capacity()
			expectedCapacity := 5 - (idx + 1)

			if capacity != expectedCapacity {
				t.Fatalf("Expected %v, got %v", expectedCapacity, capacity)
			}
		}
	})
}
