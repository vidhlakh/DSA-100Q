package main

func main() {
	n := 4
	var res [][]string
	var board []string
	helper(n, 0, 0, &res, board)

}

// colwise
func helper(n, row, col int, res *[][]string, board []string) {

	if col == n {
		return
	}
	for i := col; i < n; i++ {
		if isValid(row, col, res) {
			board[row][col] = 'Q'
			helper(row+1, col, res, board)
			board[row][col] = '.'

		}
	}
}

func isValid(row, col int, res *[][]string) bool {

	// check vertically

	// check diagonal top

	// check diagonal bottom

	return true
}
