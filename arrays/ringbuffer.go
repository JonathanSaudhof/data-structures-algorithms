package arrays

type RingBuffer[T any] struct {
	size int
	buf  []*T
	head Index
	tail Index
}

type Index struct {
	index int
	size  int
}

func (i *Index) next() {
	i.index = (i.index + 1) % (i.size - 1)
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		size: size,
		buf:  make([]*T, size),
		head: Index{index: 0, size: size},
		tail: Index{index: -1, size: size},
	}
}

func (r *RingBuffer[T]) Put(value T) {
	if r.tail.index == -1 {
		r.tail.next()
	}

	r.buf[r.tail.index] = &value
	r.tail.next()

	if r.tail.index == r.head.index {
		r.head.next()
	}
	// if the head is at the same index as the tail, move the head
	// now the last element will be overwritten
}

func (r *RingBuffer[T]) Pop() (value T) {
	value = *r.buf[r.head.index]
	if r.buf[r.head.index] == nil {
		return value
	}
	r.buf[r.head.index] = nil
	r.head.next()
	return value
}

// 1 2 3 4 5
// - - h t -
// 0 1 2 3 4
// t - h - -
// 0 1 2 3 4
func (r *RingBuffer[T]) Capacity() int {
	if r.head.index == -1 {
		return r.size
	}

	if r.head.index < r.tail.index {
		return r.size - r.tail.index - r.head.index
	}

	return 0
}

func (r *RingBuffer[T]) IsEmpty() bool {
	return r.tail.index == -1
}

func (r *RingBuffer[T]) IsFull() bool {
	return r.tail.index == r.head.index
}
