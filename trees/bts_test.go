package trees

import (
	"testing"
)

func TestCompare_IdenticalTrees_ShouldValidResult(t *testing.T) {

	t.Run("Should return true", func(t *testing.T) {

		test := BinaryNode[int]{
			Value: 4,
			left: &BinaryNode[int]{
				Value: 2,
				left: &BinaryNode[int]{
					Value: 1,
				},
				right: &BinaryNode[int]{
					Value: 3,
				},
			},
			right: &BinaryNode[int]{
				Value: 6,
				left: &BinaryNode[int]{
					Value: 5,
				},
				right: &BinaryNode[int]{
					Value: 7,
				},
			},
		}

		expected := BinaryNode[int]{
			Value: 4,
			left: &BinaryNode[int]{
				Value: 2,
				left: &BinaryNode[int]{
					Value: 1,
				},
				right: &BinaryNode[int]{
					Value: 3,
				},
			},
			right: &BinaryNode[int]{
				Value: 6,
				left: &BinaryNode[int]{
					Value: 5,
				},
				right: &BinaryNode[int]{
					Value: 7,
				},
			},
		}

		if !test.Compare(&expected) {
			t.Fatalf("should return  true")
		}
	})

	t.Run("Should return false", func(t *testing.T) {

		test := BinaryNode[int]{
			Value: 4,
			left: &BinaryNode[int]{
				Value: 2,
				left: &BinaryNode[int]{
					Value: 1,
				},
				right: &BinaryNode[int]{
					Value: 3,
				},
			},
			right: &BinaryNode[int]{
				Value: 6,
				left: &BinaryNode[int]{
					Value: 5,
				},
				right: &BinaryNode[int]{
					Value: 7,
				},
			},
		}

		expected := BinaryNode[int]{
			Value: 4,
			left: &BinaryNode[int]{
				Value: 2,
				left: &BinaryNode[int]{
					Value: 1,
				},
				right: &BinaryNode[int]{
					Value: 3,
				},
			},
			right: &BinaryNode[int]{
				Value: 6,
				left: &BinaryNode[int]{
					Value: 5,
				},
			},
		}

		if test.Compare(&expected) {
			t.Fatalf("should return false")
		}
	})
}

func TestBSTCreate_ShouldCreateTree(t *testing.T) {

	t.Run("Should create a balanced binary tree", func(t *testing.T) {
		head := CreateBST([]int{1, 2, 3, 4, 5, 6, 7})

		expected := &BinaryNode[int]{
			Value: 4,
			left: &BinaryNode[int]{
				Value: 2,
				left: &BinaryNode[int]{
					Value: 1,
				},
				right: &BinaryNode[int]{
					Value: 3,
				},
			},
			right: &BinaryNode[int]{
				Value: 6,
				left: &BinaryNode[int]{
					Value: 5,
				},
				right: &BinaryNode[int]{
					Value: 7,
				},
			},
		}

		if !head.compare(expected) {
			t.Fatalf("should be created like expected")
		}
	})
}

func TestBSTInsert_ValidValue_ShouldInsertCorrectly(t *testing.T) {

	t.Run("Should insert at the correct place", func(t *testing.T) {
		head := CreateBST([]int{1, 2, 3, 5, 6, 7})

		head.Insert(4)

		expected := &BinaryNode[int]{
			Value: 5,
			left: &BinaryNode[int]{
				Value: 2,
				left: &BinaryNode[int]{
					Value: 1,
				},
				right: &BinaryNode[int]{
					Value: 3,
					right: &BinaryNode[int]{
						Value: 4,
					},
				},
			},
			right: &BinaryNode[int]{
				Value: 7,
				left: &BinaryNode[int]{
					Value: 6,
				},
			},
		}

		if !head.compare(expected) {
			t.Fatalf("should be inserted at the correct place")
		}
	})
}

func TestBSTDelete_ValidValue_ShouldCreateTree(t *testing.T) {

	t.Run("Should delete expected value from root", func(t *testing.T) {
		head := CreateBST([]int{100})

		head.Insert(125)
		head.Insert(75)
		head.Insert(85)
		head.Insert(80)
		head.Insert(95)
		head.Insert(93)

		expected := &BinaryNode[int]{
			Value: 95,
			left: &BinaryNode[int]{
				Value: 75,
				right: &BinaryNode[int]{
					Value: 85,
					left: &BinaryNode[int]{
						Value: 80,
					},
					right: &BinaryNode[int]{
						Value: 95,
						left: &BinaryNode[int]{
							Value: 93,
						},
					},
				},
			},
			right: &BinaryNode[int]{
				Value: 125,
			},
		}

		head.Delete(100)

		if !head.compare(expected) {
			t.Fatalf("should be deleted")
		}
	})
}
