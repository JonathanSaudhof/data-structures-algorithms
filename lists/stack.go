package lists

import "errors"

type stackElement[T any] struct {
	value T
	prev  *stackElement[T]
}

type stack[T any] struct {
	size int
	head *stackElement[T]
}

const noStackElementError = "stack has no elements"

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

func (s *stack[T]) Pop() (value T, err error) {

	if s.head == nil {
		err = errors.New(noStackElementError)
		return
	}

	var currElement stackElement[T] = *s.head
	prevElement := currElement.prev

	var isLastElement bool = prevElement == nil

	if isLastElement {
		s.head = nil
	} else {
		s.head = prevElement
	}

	s.size--

	value = currElement.value
	return

}

func (s *stack[T]) Top() T {
	return s.head.value
}
