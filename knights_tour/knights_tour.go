package knights_tour

import "fmt"

type cell struct {
	row,
	col int
}

// board
var board [][]int

// board size
var rowCount int
var colCount int

// available movement for knight
var dirs = []cell{
	{-1, -2}, // left top
	{-1, 2},  // right top
	{-2, -1}, // top left
	{-2, 1},  // top right
	{1, -2},  // left bottom
	{1, 2},   // right bottom
	{2, -1},  // bottom left
	{2, 1},   // bottom right
}

func Run(startRow, startCol, rowCountParam, colCountParam int) bool {
	// init board
	rowCount, colCount = rowCountParam, colCountParam
	board = make([][]int, rowCount)
	for i := 0; i < rowCount; i++ {
		board[i] = make([]int, colCount)
		for j := 0; j < colCount; j++ {
			board[i][j] = -1
		}
	}

	// start from row 0, col 0, step 0
	r, c, step := startRow, startCol, 0
	found := dfs(r, c, step)

	// print result
	fmt.Printf("Knights tour ")
	if found {
		fmt.Printf("CAN ")
	} else {
		fmt.Printf("CAN'T ")
	}
	fmt.Printf("be done with %d x %d board starting from row: %d, col: %d\n", rowCount, colCount, startRow, startCol)

	// print board
	if found {
		fmt.Println("Board:")
		for i := 0; i < rowCount; i++ {
			for j := 0; j < colCount; j++ {
				fmt.Printf("%3d", board[i][j])
			}
			fmt.Println()
		}
	}

	return found
}

func dfs(row, col, step int) bool {
	// All cell visited
	if step == rowCount*colCount {
		return true
	}

	// Check border
	if row < 0 || row >= rowCount || col < 0 || col >= colCount {
		return false
	}

	// Check visited
	if board[row][col] != -1 {
		return false
	}

	// if all checks above valid, store the step in the board
	board[row][col] = step

	// loop for next move in each direction
	for _, dir := range dirs {
		// recurse to next move,
		// if the solution LEADS to true (bubbling up from base cases)
		solutionFound := dfs(row+dir.row, col+dir.col, step+1)

		if solutionFound {
			return true
		}
	}

	// if no solution found, reset
	board[row][col] = -1

	return false
}
