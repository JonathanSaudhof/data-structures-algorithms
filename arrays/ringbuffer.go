package arrays

type RingBuffer[T any] struct {
	size   int
	buf    []*T
	read   Index
	write  Index
	isFull bool
}

type Index struct {
	index int
	size  int
}

func (i *Index) next() {
	i.index = (i.index + 1) % (i.size)
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return ringBufferInit[T](size)
}

func ringBufferInit[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		size:   size,
		buf:    make([]*T, size),
		read:   Index{index: 0, size: size},
		write:  Index{index: -1, size: size},
		isFull: false,
	}
}

func (r *RingBuffer[T]) Put(value T) {
	if r.IsFull() {
		r.read.next()
	}

	if r.IsEmpty() {
		r.write.next()
	}

	r.buf[r.write.index] = &value
	r.write.next()

	if r.write.index == r.read.index {
		r.isFull = true
	}
}

func (r *RingBuffer[T]) Pop() (value T) {
	value = *r.buf[r.read.index]
	r.buf[r.read.index] = nil
	r.read.next()

	if r.isFull {
		r.isFull = false
	}

	if r.write.index == r.read.index {
		r.setEmpty()
	}

	return value
}

func (r *RingBuffer[T]) setEmpty() {
	r.write.index = -1
}

func (r *RingBuffer[T]) Length() int {
	if r.IsEmpty() {
		return 0
	}

	if r.IsFull() {
		return r.size
	}

	if r.write.index > r.read.index {
		return r.write.index - r.read.index
	}

	return r.size - r.read.index + r.write.index
}

func (r *RingBuffer[T]) Flush() []T {
	values := make([]T, r.Length())
	for idx := range values {
		values[idx] = r.Pop()
	}

	return values
}

func (r *RingBuffer[T]) Capacity() int {
	return r.size - r.Length()
}

func (r *RingBuffer[T]) IsEmpty() bool {
	return r.write.index == -1
}

func (r *RingBuffer[T]) IsFull() bool {
	return r.isFull
}

func (r *RingBuffer[T]) Peek() T {
	return *r.buf[r.read.index]
}
