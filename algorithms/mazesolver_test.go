package algorithms_test

import (
	"testing"

	"portfolio/das/algorithms"
)

func TestMazeSolver_ReachAbleEnd_ShouldReturnPath(t *testing.T) {
	maze := [][]string{
		{"#", "#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", "S", " ", " ", " ", " ", " ", " ", "#"},
		{"#", "#", "#", "#", "#", "#", " ", "#", "#"},
		{"#", " ", " ", " ", " ", " ", " ", " ", "#"},
		{"#", " ", "#", "#", "#", "#", "#", "#", "#"},
		{"#", " ", " ", " ", " ", " ", " ", " ", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "E", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#", "#"},
	}

	expectedPath := []algorithms.Point{
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 1},
		{6, 2},
		{6, 3},
		{5, 3},
		{4, 3},
		{3, 3},
		{2, 3},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 5},
		{5, 5},
		{6, 5},
		{7, 5},
		{7, 6},
	}

	t.Run("Maze solver should return the path to the end", func(t *testing.T) {
		path := algorithms.SolveMaze(maze)

		algorithms.PrintMaze(&maze, &path)

		if len(path) != len(expectedPath) {
			t.Fatalf("Expected path to be %v, but got %v", expectedPath, path)
		}

		for i := range path {
			if path[i] != expectedPath[i] {
				t.Fatalf("Expected path to be %v, but got %v", expectedPath, path)
			}
		}
	})
}

func TestMazeSolver_NoReachAbleEnd_ShouldReturnEmptyPath(t *testing.T) {
	maze := [][]string{
		{"#", "S", "#", "#", "#", "#", "#", "#", "#"},
		{"#", " ", " ", " ", " ", "#", " ", " ", "#"},
		{"#", "#", "#", "#", "#", "#", "E", "#", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#", "#"},
	}

	t.Run("Maze solver should return an empty path", func(t *testing.T) {
		path := algorithms.SolveMaze(maze)

		algorithms.PrintMaze(&maze, &path)

		if len(path) != 0 {
			t.Fatalf("Expected path to be empty, but got %v", path)
		}
	})
}
