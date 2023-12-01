package main

func main() {
	n := 4
	var res [][]string
	var board []string
	for i := 0; i < n; i++ {
		board[i] = "...."
	}
	helper(n, 0, &res, board)

}

// colwise
func helper(n, col int, res *[][]string, board []string) {

	if col == n {
		*res = append(*res, board)
		return
	}
	for row := 0; row < n; row++ {
		if isValid(row, col, n, board) {
			runeBoard := []rune(board[row])
			runeBoard[col] = 'Q'
			board[row] = string(runeBoard)
			helper(n, col+1, res, board)
			runeBoard = []rune(board[row])
			runeBoard[col] = '.'
			board[row] = string(runeBoard)
		}
	}
}

func isValid(row, col, n int, board []string) bool {

	// check horizontal left
	for j := col; j >= 0; j-- {
		r := []rune(board[row])
		if r[j] == 'Q' {
			return false
		}
	}
	// check diagonal top

	for i, j := row, col; i >= 0; i, j = i-1, j-1 {
		r := []rune(board[i])
		if r[j] == 'Q' {
			return false
		}
	}

	// check diagonal bottom
	for i := row; i < n; i++ {
		for j := col; j >= 0; j-- {
			r := []rune(board[i])
			if r[j] == 'Q' {
				return false
			}
		}
	}
	return true
}
