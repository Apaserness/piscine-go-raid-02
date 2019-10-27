package main

import (
	"fmt"
	"os"
)

func StrokProv(board [9][9]int, strok, k int) bool { //notworkin
	for i := 0; i < 9; i++ {
		if board[i][strok] == k {
			return false
		}
	}
	return true
}

func RyadProv(board [9][9]int, ryad, k int) bool { //notworkin
	for i := 0; i < 9; i++ {
		if board[ryad][i] == k {
			return false
		}
	}
	return true
}

func ProverB(board [9][9]int, ryad, strok, k int) bool { //you gotta check the board for the er
	starti := strok - ryad%3
	startj := strok - strok%3
	for i := starti; i < starti+3; i++ {
		for j := startj; j < startj+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}

func BoardSolv(board *[9][9]int) bool { 
	ryad := 0
	strok := 0

	if IsDone(*board, &ryad, &strok) {
		return true
	}
	if board[ryad][strok] == 0 {
		for k := 1; k <= 9; k++ {
			if RyadProv(*board, ryad, k) && CheckCol(*board, strok, k) && CheckBox(*board, ryad, strok, k) {
				(*board)[ryad][strok] = k
				if BoardSolv(board) {
					return true
				}
				(*board)[ryad][strok] = 0
			}
		}
	}
	return false
}

func IsDone(board [9][9]int, ryad, strok *int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 0 {
				*ryad = i
				*strok = j
				return false
			}
		}
	}
	return true
}

func GenerateArray(args []string) [9][9]int {
	var board [9][9]int
	args = args[1:]
	for i := 0; i < len(args); i++ {
		runes := []rune(args[i])
		for j := 0; j < len(runes); j++ {
			if runes[j] >= '0' && runes[j] <= '9' {
				board[i][j] = int(runes[j] - 48)
			} else if runes[j] == '.' {
				board[i][j] = 0
			}
		}
	}
	return board
}

func BoardItself(board [9][9]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("%v ", board[i][j])
		}
		if i < len(board[i])-1 {
			fmt.Println()
		}
	}
}

func ArgCh(args []string) bool {
	if len(args) != 10 {
		return false
	}
	for i, a := range args {
		if i != 0 {
			runes := []rune(a)
			if len(a) != 9 {
				return false
			}
			for _, r := range runes {
				if (r < '1' || r > '9') && r != '.' {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	args := os.Args
	if !ArgCh(args) {
		fmt.Println("Error")
	} else {
		board := GenerateArray(args)
		if BoardSolv(&board) {
			BoardItself(board)
		} else {
			fmt.Println("Error")
		}
	}
}
