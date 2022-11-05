package main

import (
	"fmt"
	"os"
)

//############ Test le board sudoku ############

// fonction
func SudokuErreur(arr []string) bool {
	for i, str := range arr {
		if len(str) != 9 {
			return true
		}
		for _, ch := range arr[i] {
			if !(ch >= '1' && ch <= '9' || ch == '.') {
				return true
			}
		}
	}
	return false
}

// fonction
func Test(arr [][]int, row int, column int, num int) bool {
	return !VerifieLigne(arr, row, num) && !VerifieColonne(arr, column, num) && !VerifieSubSudoku(arr, row-(row%3), column-(column%3), num)
}

// fonction
func VerifieLigne(arr [][]int, row int, num int) bool {
	for column := 0; column < len(arr); column++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

// fonction
func VerifieColonne(arr [][]int, column int, num int) bool {
	for row := 0; row < len(arr); row++ {
		if arr[row][column] == num {
			return true
		}
	}
	return false
}

// fonction
func VerifieSubSudoku(arr [][]int, rowIndex int, columnIndex int, num int) bool {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if arr[rowIndex+row][columnIndex+column] == num {
				return true
			}
		}
	}
	return false
}

//##############################################
//################# Solveur ####################

// fonction
func Solver(arr *[][]int, length int) bool {
	isEmpty := true
	row := -1
	column := -1
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if (*arr)[i][j] == 0 {
				row = i
				column = j
				isEmpty = false
				break
			}
		}
	}
	if isEmpty {
		return true
	}
	for number := 1; number <= 9; number++ {
		if Test(*arr, row, column, number) {
			(*arr)[row][column] = number

			if Solver(arr, length) {
				return true
			} else {
				(*arr)[row][column] = 0
			}
		}
	}
	return false
}

//##############################################
//############### Converstion ##################

// fonction
func RuneToInt(number rune) int {
	count := 0
	for i := '1'; i <= number; i++ {
		count++
	}
	return count
}

// fonction
func ToInt(arr []string) [][]int {
	sudoku := make([][]int, 9)
	for i := range sudoku {
		sudoku[i] = make([]int, 9)
	}
	for i, str := range arr {
		for j, ch := range str {
			sudoku[i][j] = RuneToInt(ch)
		}
	}
	return sudoku
}

//###############################################
//############# Affichage du sudoku #############

// fonction
func AfficherSudoku(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j])
			if j < len(arr)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//################################################
//############# Fonction main ####################

func main() {
	arr := os.Args[1:]

	if SudokuErreur(arr) {
		fmt.Println("Error")
	} else {
		sudoku := ToInt(arr)
		if Solver(&sudoku, len(sudoku)) {
			AfficherSudoku(sudoku)
		} else {
			fmt.Println("Error")
		}
	}
}

//################################################
