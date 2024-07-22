package arrays

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

	walk(&maze, &visited, &path, start, end)

	return path
}

func findStartAndEnd(maze [][]string) (start Point, end Point) {
	for x, row := range maze {
		for y, piece := range row {
			if piece == START {
				start = Point{
					y, x,
				}
			}
			if piece == END {
				end = Point{y, x}
			}
		}
	}

	if start == (Point{}) || end == (Point{}) {
		panic("Maze must have a start and end point")
	}

	return
}

var dir = [][]int{
	{-1, 0}, // TOP
	{0, 1},  // RIGHT
	{1, 0},  // BOTTOM
	{0, -1}, // LEFT
}

func walk(maze *[][]string, visited *[][]bool, path *[]Point, current Point, end Point) bool {
	isEnd := current.Column == end.Column && current.Row == end.Row

	if isEnd {
		*path = append(*path, current)
		fmt.Println("Found end", current)
		return true
	}

	isOutOfBoard := current.Column > len(*maze) ||
		current.Column < 0 ||
		current.Row > len((*maze)[0]) ||
		current.Row < 0

	if isOutOfBoard {
		fmt.Println("Out of board", current)
		return false
	}

	isWall := (*maze)[current.Row][current.Column] == WALL

	if isWall {
		fmt.Println("Hit wall", current)
		return false
	}

	var hasAlreadyBeenVisited bool = (*visited)[current.Row][current.Column]

	if hasAlreadyBeenVisited {
		fmt.Println("Already visited", current)
		return false
	}

	(*visited)[current.Row][current.Column] = true
	*path = append(*path, current)
	fmt.Println("Remember path", current)

	for _, dir := range dir {
		next := Point{
			Column: current.Column + dir[1],
			Row:    current.Row + dir[0],
		}

		if walk(maze, visited, path, next, end) {
			return true
		}
	}
	*path = (*path)[:len(*path)-1]
	return false
}

func PrintMaze(maze *[][]string, path *[]Point) {
	for y, row := range *maze {
		for x, piece := range row {
			isPath := false
			for _, path := range *path {
				if path.Column == x && path.Row == y {
					isPath = true
				}
			}
			if isPath {
				fmt.Print("X")
			} else {
				fmt.Print(piece)
			}
		}
		fmt.Println()
	}
}
