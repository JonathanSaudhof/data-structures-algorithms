package trees

import (
	"cmp"
	"errors"
)

type BinaryNode[T cmp.Ordered] struct {
	Value  T
	parent *BinaryNode[T]
	left   *BinaryNode[T]
	right  *BinaryNode[T]
}

func insert[T cmp.Ordered](parent, node *BinaryNode[T], value T) {

	if node == nil {

		if value <= parent.Value {
			parent.left = &BinaryNode[T]{
				Value:  value,
				parent: parent,
			}
			return
		}

		parent.right = &BinaryNode[T]{
			Value:  value,
			parent: parent,
		}
		return
	}

	if value <= (*node).Value {
		insert(node, node.left, value)
	}

	if value > (*node).Value {
		insert(node, node.right, value)
	}

}

func (t *BinaryNode[T]) Insert(value T) {
	insert(t, t, value)
}

func createBalanced[T cmp.Ordered](nodeValues []T, parent *BinaryNode[T]) (node *BinaryNode[T], err error) {

	if len(nodeValues) == 0 {
		err = errors.New("no elements left")
		return
	}

	midIdx := len(nodeValues) / 2
	value := nodeValues[midIdx]

	node = &BinaryNode[T]{
		Value:  value,
		parent: parent,
	}

	leftRest := nodeValues[0:midIdx]
	rightRest := nodeValues[midIdx+1:]

	leftElement, leftErr := createBalanced(leftRest, node)
	rightElement, rightErr := createBalanced(rightRest, node)

	if leftErr == nil {
		node.left = leftElement
	}

	if rightErr == nil {
		node.right = rightElement
	}

	return
}

func (a *BinaryNode[T]) compare(b *BinaryNode[T]) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return a.left.compare(b.left) && a.right.compare(b.right)

}

func (a *BinaryNode[T]) Compare(b *BinaryNode[T]) bool {
	return a.compare(b)
}

func (t *BinaryNode[T]) findLargestValue() *BinaryNode[T] {
	if t.right == nil {
		return t
	}

	return t.right.findLargestValue()
}

func (t *BinaryNode[T]) Delete(value T) (err error) {
	node, err := t.findNode(value)

	if err != nil {
		return
	}

	isLeftItem := t.parent != nil && value <= t.parent.Value

	if node.left == nil && node.right == nil {
		if isLeftItem {
			t.parent.left = nil
			node.parent = nil
			return
		}

		t.parent.right = nil
		return
	}

	if node.left == nil || node.right == nil {

		if isLeftItem {
			t.parent.left = node.left
			node.parent = nil
			return
		}

		t.parent.right = node.right
		node.parent = nil
		return
	}

	replacement := t.left.findLargestValue()

	hasSmallerItem := replacement.left != nil

	if hasSmallerItem {
		replacement.parent.right = node.left
		replacement.parent.right.parent = replacement.parent

	} else {
		replacement.parent.right = nil
	}

	replacement.right = node.right
	replacement.left = node.left
	*node = *replacement
	node.parent = nil

	return
}

func (t *BinaryNode[T]) findNode(value T) (node *BinaryNode[T], err error) {
	if value == t.Value {
		node = t
		return
	}

	if value <= t.Value {
		return t.left.findNode(value)
	}

	if value > t.Value {
		return t.left.findNode(value)
	}

	err = errors.New("not found")

	return

}

// Creates Binary Search Tree
func CreateBST[T cmp.Ordered](nodeValues []T) (head *BinaryNode[T]) {
	head, _ = createBalanced(nodeValues, nil)
	return
}
