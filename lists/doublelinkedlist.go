package lists

import "errors"

type Node[T comparable] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoubleLinkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewDoubleLinkedList[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (l *DoubleLinkedList[T]) Length() int {
	return l.length
}

const (
	EMPTY_LIST_ERROR      = "list is empty"
	VALUE_NOT_FOUND_ERROR = "could not find value in list"
)

func (l *DoubleLinkedList[T]) InsertAt(idx int, value T) (err error) {

	if idx > l.length || idx < 0 {
		err = errors.New("index out of bound")
		return
	}

	if idx == 0 {
		l.Prepend(value)
		return
	}

	if idx == l.length {
		l.Append(value)
		return
	}

	newElement := Node[T]{
		value: value,
	}

	element, err := l.getByIdx(idx)

	newElement.prev = element.prev
	newElement.next = element

	element.prev.next = &newElement
	element.prev = &newElement

	l.length++

	return
}

func (l *DoubleLinkedList[T]) getByIdx(idx int) (node *Node[T], err error) {

	if l.head == nil {
		err = errors.New(EMPTY_LIST_ERROR)
		return
	}

	curr := l.head

	for i := 0; i < idx; i++ {

		if curr.next == nil {
			err = errors.New(VALUE_NOT_FOUND_ERROR)
			return
		}
		curr = curr.next
	}

	node = curr

	return
}

func (l *DoubleLinkedList[T]) addAsFirstElement(value T) bool {
	if l.tail == nil {

		newElement := Node[T]{
			value: value,
		}

		l.tail = &newElement
		l.head = &newElement

		return true
	}

	return false
}

func (l *DoubleLinkedList[T]) Append(value T) {
	l.length++

	if l.addAsFirstElement(value) {
		return
	}

	newElement := Node[T]{
		value: value,
	}

	lastElement := l.tail
	lastElement.next = &newElement
	newElement.prev = lastElement
	l.tail = &newElement
}

func (l *DoubleLinkedList[T]) Prepend(value T) {
	l.length++

	if l.addAsFirstElement(value) {
		return
	}

	newElement := Node[T]{
		value: value,
	}

	firstElement := *l.head
	firstElement.prev = &newElement
	newElement.next = &firstElement

}

func (l *DoubleLinkedList[T]) findValueInList(value T) (idx int, node *Node[T], err error) {

	idx = -1

	if l.head == nil {
		err = errors.New(EMPTY_LIST_ERROR)
		return
	}

	curr := l.head

	for curr != nil {
		idx++

		if curr.value == value {
			node = curr
			return
		}

		curr = curr.next
	}

	return -1, nil, errors.New(VALUE_NOT_FOUND_ERROR)

}

func (l *DoubleLinkedList[T]) Remove(value T) (idx int, err error) {
	var element *Node[T]

	idx, element, err = l.findValueInList(value)

	if err != nil {
		return
	}

	element.detachNode()

	l.length--

	return
}

func (n *Node[T]) detachNode() {
	prevElement := n.prev
	nextElement := n.next

	prevElement.next = nextElement
	nextElement.prev = prevElement

	n.next = nil
	n.prev = nil
}

func (l *DoubleLinkedList[T]) RemoveAt(index int) (err error) {
	element, err := l.getByIdx(index)

	if err != nil {
		return
	}

	element.detachNode()

	return
}

func (l *DoubleLinkedList[T]) Get(idx int) (value T, err error) {
	node, err := l.getByIdx(idx)
	if err != nil {
		return
	}
	value = node.value
	return
}
