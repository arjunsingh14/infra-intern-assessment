package main


// SolveSudoku is the main function that is called to solve the Sudoku puzzle.
// It calls the helper function 'solve' and returns the solved game.
func SolveSudoku(game [][]int) [][]int {
    helper(game)
    return game
}

// helper function that uses backtracking to solve the Sudoku puzzle.
// It iterates over each cell in the board. If it finds an empty cell (0), it tries to fill it with a number from 1 to 9.
// If the number is valid (doesn't violate Sudoku rules), it recursively calls solve to fill in the rest of the board.
// If it can't find a valid number, it backtracks by setting the cell back to 0 and returning false.
func helper(game [][]int) bool {
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if game[i][j] == 0 {
                for num := 1; num <= 9; num++ {
                    if check_valid(game, i, j, num) {
                        game[i][j] = num
                        if helper(game) {
                            return true
                        }
                        game[i][j] = 0
                    }
                }
                return false
            }
        }
    }
    return true
}


// The check_valid function checks if a number is valid in a given cell.
// It returns true if the number is not already used in the same row, column, and 3x3 box.
func check_valid(board [][]int, row, col, num int) bool {
    for i := 0; i < 9; i++ {
        if board[i][col] == num {
            return false
        }
        if board[row][i] == num {
            return false
        }
        if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
            return false
        }
    }
    return true
}
