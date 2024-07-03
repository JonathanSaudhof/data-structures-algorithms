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

//	tail  head
//  []  =>  []

func Enqueue[T any](q *queue[T], element T) {

	newQueueElement := queueElement[T]{}
	newQueueElement.value = element

	if q.head == nil {
		q.head = &newQueueElement
	}

	if q.tail == nil {
		q.tail = &newQueueElement
	}

}
