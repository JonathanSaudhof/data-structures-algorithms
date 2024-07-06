package lists

import "errors"

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

const noQueueElementError = "queue has no elements"

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

func (q *queue[T]) Dequeue() (value T, err error) {

	if q.head == nil {
		err = errors.New(noQueueElementError)
		return
	}

	var currElement queueElement[T] = *q.head

	prevElement := currElement.prev
	var isLastElement bool = prevElement == nil

	if isLastElement {
		q.tail = nil
		q.head = nil
	} else {
		prevElement.next = nil
		q.head = prevElement
	}

	value = currElement.value
	q.size--

	return

}

func (q *queue[T]) Peek() (value T, err error) {

	if q.head == nil {
		err = errors.New(noQueueElementError)
		return
	}

	value = q.head.value
	return
}
