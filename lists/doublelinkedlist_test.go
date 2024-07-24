package lists_test

import (
	"portfolio/das/lists"
	"testing"
)

func TestAppendToList_EmptyList_ShouldAddFirstElement(t *testing.T) {
	t.Run("Append Should add first element to empty list", func(t *testing.T) {

		list := lists.NewDoubleLinkedList[int]()
		list.Append(10)

		if list.Length() != 1 {
			t.Fatalf("List should have length of 1")
		}

		element, _ := list.Get(0)

		if element != 10 {
			t.Fatalf("Item should be as expected")
		}
	})
}

func TestAppendToList_ExistingList_ShouldAddToTheEnd(t *testing.T) {
	t.Run("Append Should add first element to empty list", func(t *testing.T) {

		list := lists.NewDoubleLinkedList[int]()
		list.Append(10)

		if list.Length() != 1 {
			t.Fatalf("List should have length of 1")
		}

		element, _ := list.Get(0)

		if element != 10 {
			t.Fatalf("Item should be as expected")
		}
	})
}

func TestPrependToList_EmptyList_ShouldAddFirstElement(t *testing.T) {
	t.Run("Prepend should add first element to empty list", func(t *testing.T) {
		list := lists.NewDoubleLinkedList[int]()
		list.Prepend(10)
		if list.Length() != 1 {
			t.Fatalf("List should have length of 1")
		}

		element, _ := list.Get(0)

		if element != 10 {
			t.Fatalf("Item should be as expected")
		}
	})
}

func TestGet_WrongIndex_ShouldReturnError(t *testing.T) {

	list := lists.NewDoubleLinkedList[int]()

	t.Run("Get should return an error on empty list", func(t *testing.T) {

		_, err := list.Get(0)

		if err.Error() != lists.EMPTY_LIST_ERROR {
			t.Fatalf("Get should return an errror as list is empty")
		}
	})

	t.Run("Get should return an error on wrong index", func(t *testing.T) {

		list.Append(10)
		_, err := list.Get(0)

		if err != nil {
			t.Fatalf("Get should not return an errror for given index")
		}
		_, err = list.Get(1)

		if err.Error() == lists.VALUE_NOT_FOUND_ERROR {
			t.Fatalf("Get should return error for value not found")
		}

	})
}

func TestGet_FilledList_ShouldReturnInTheRightOrder(t *testing.T) {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := lists.NewDoubleLinkedList[int]()

	fillList(items, list, t)

	t.Run("Get Should return elements in correct orrder", func(t *testing.T) {

		for i := 0; i < len(items); i++ {
			element, err := list.Get(i)
			if element != items[i] {
				t.Fatalf("Expected element to be %v, but got %v", items[i], element)
			}

			if err != nil {
				t.Fatalf("Error should be nil")
			}
		}
	})

}

func TestRemove_ValidValue_ShouldRemoveElement(t *testing.T) {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}

	list := lists.NewDoubleLinkedList[int]()

	fillList(items, list, t)

	t.Run("Remove should remove element from list", func(t *testing.T) {

		list.Remove(5)

		for i := 0; i < len(expected); i++ {
			element, err := list.Get(i)
			if element != expected[i] {
				t.Fatalf("Expected element to be %v, but got %v", items[i], element)
			}

			if err != nil {
				t.Fatalf("Error should be nil")
			}
		}
	})

}

func TestRemove_InvalidValue_ShouldRemoveElement(t *testing.T) {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	list := lists.NewDoubleLinkedList[int]()

	t.Run("Remove should return error on empty list", func(t *testing.T) {
		_, err := list.Remove(5)
		if err.Error() != lists.EMPTY_LIST_ERROR {
			t.Fatalf("Should return error for empty list")
		}
	})

	fillList(items, list, t)

	t.Run("Remove should return error on not found element", func(t *testing.T) {
		_, err := list.Remove(11)
		if err.Error() != lists.VALUE_NOT_FOUND_ERROR {
			t.Fatalf("Should return error on for not found value")
		}
	})

}

func TestRemoveAt_ValidValue_ShouldRemoveElement(t *testing.T) {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}

	list := lists.NewDoubleLinkedList[int]()

	fillList(items, list, t)

	t.Run("Remove should remove element from list", func(t *testing.T) {

		list.RemoveAt(4)

		for i := 0; i < len(expected); i++ {
			element, err := list.Get(i)
			if element != expected[i] {
				t.Fatalf("Expected element to be %v, but got %v", items[i], element)
			}

			if err != nil {
				t.Fatalf("Error should be nil")
			}
		}
	})

}

func TestRemoveAt_InvalidValue_ShouldRemoveElement(t *testing.T) {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	list := lists.NewDoubleLinkedList[int]()

	t.Run("Remove should return error on empty list", func(t *testing.T) {
		err := list.RemoveAt(5)

		if err.Error() != lists.EMPTY_LIST_ERROR {
			t.Fatalf("Should return error for empty list")
		}
	})

	fillList(items, list, t)

	t.Run("Remove should return error on not found element", func(t *testing.T) {
		err := list.RemoveAt(10)
		if err.Error() != lists.VALUE_NOT_FOUND_ERROR {
			t.Fatalf("Should return error on for not found value")
		}
	})

}

func fillList(items []int, list *lists.DoubleLinkedList[int], t *testing.T) {
	for _, value := range items {
		list.Append(value)
	}

	if list.Length() != len(items) {
		t.Fatalf("List should have length of %v", len(items))
	}
}
