package lists

type queueElement[T any] struct {
	value T
	next  *queueElement[T]
	prev  *queueElement[T]
}

type queue[T any] struct {
	size int
	tail *queueElement[T]
	head *queueElement[T]
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Enqueue(value T) {

	newQueueElement := queueElement[T]{
		value: value,
	}

	var isFirstElement bool = q.head == nil && q.tail == nil

	if isFirstElement {
		q.head = &newQueueElement
		q.tail = &newQueueElement
	} else {
		lastElement := q.tail
		lastElement.prev = &newQueueElement
		newQueueElement.next = lastElement
		q.tail = &newQueueElement
	}

	q.size++
}

func (q *queue[T]) Dequeue() T {

	var currElement queueElement[T] = *q.head

	prevElement := currElement.prev
	var isLastElement bool = prevElement == nil

	if isLastElement {
		q.tail = nil
	} else {
		prevElement.next = nil
		q.head = prevElement
	}

	return currElement.value

}

func (q *queue[T]) Peek() T {
	return q.head.value
}
