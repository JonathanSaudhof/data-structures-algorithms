package lists

type stackElement[T any] struct {
	value T
	prev  *stackElement[T]
}

type stack[T any] struct {
	size int
	head *stackElement[T]
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(value T) {

	newStackElement := stackElement[T]{
		value: value,
	}

	lastElement := s.head

	if lastElement == nil {
		s.head = &newStackElement
	} else {
		newStackElement.prev = lastElement
		s.head = &newStackElement
	}

	s.size++
}

func (s *stack[T]) Pop() T {

	var currElement stackElement[T] = *s.head
	prevElement := currElement.prev

	var isLastElement bool = prevElement == nil

	if isLastElement {
		s.head = nil
	} else {
		s.head = prevElement
	}

	return currElement.value

}

func (s *stack[T]) Top() T {
	return s.head.value
}
