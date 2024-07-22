package algorithms

import "fmt"

type MazePiece = string

const (
	WALL  MazePiece = "#"
	SPACE MazePiece = " "
	START MazePiece = "S"
	END   MazePiece = "E"
)

type Point struct {
	Column, Row int
}

func SolveMaze(maze [][]string) []Point {
	// Find start and end points
	start, end := findStartAndEnd(maze)
	// Create a visited matrix
	visited := make([][]bool, len(maze))

	path := make([]Point, 0)

	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	// Solve the maze
	walk(&maze, &visited, start, end)

	return path
}

func findStartAndEnd(maze [][]string) (start Point, end Point) {
	for i, row := range maze {
		for j, piece := range row {
			if piece == START {
				start = Point{i, j}
			}
			if piece == END {
				end = Point{i, j}
			}
		}
	}

	if start == (Point{}) || end == (Point{}) {
		panic("Maze must have a start and end point")
	}

	return
}

const dir = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func walk(maze *[][]string, visited *[][]bool, path *[]Point, current Point, end Point) bool {
	isEnd := current == end

	if isEnd {
		*path = append(*path, current)
		return true
	}

	isOutOfColumn := current.Column > len(*maze) || current.Column < 0
	isOutOfRow := current.Row > len((*maze)[current.Column]) || current.Row < 0
	isOutOfBoard := isOutOfColumn || isOutOfRow

	if isOutOfBoard {
		fmt.Println("Out of board")
		return false
	}

	isWall := (*maze)[current.Row][current.Column] == WALL

	if isWall {
		fmt.Println("Hit wall")
		return false
	}

	var hasAlreadyBeenVisited bool = (*visited)[current.Column][current.Row]

	if hasAlreadyBeenVisited {
		fmt.Println("Already visited")
		return false
	}

	for _, value := range dir {
		next := current.
			walk(maze, visited, path)
	}

	(*visited)[current.Column][current.Row] = true

	return false
}
