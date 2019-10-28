package main

import (
	"fmt"
	"os"
)

func NormArgs(str string) bool {
	for i := range str {
		if !(str[i] == '.' ||(str[i] >= '1' && str[i] <= '9')) {
			return false
		}
	}
	return true
}

func Parsing(args []string, result *[9][9]int) bool {
	for i := 1; i < len(args); i++ {
		if len(args[i]) != 9 {
			return false
		} else {
			if NormArgs(args[i]) {
				for j := 0; j < 9; j++ {
					if args[i][j] != '.' {
						result[i-1][j] = int(args[i][j] - 48)
					}
				}
			} else {
				return false
			}
		}
	}
	return true
}

func Backtrack(result *[9][9]int) bool {
	if !EmptyCell(*result) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if result[i][j] == 0 {
				for k := 1; k <= 9; k++ {
					result[i][j] = k
					if isTableValid(*result) {
						if Backtrack(result) {
							return true
						}
						result[i][j] = 0
					} else {
						result[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func revBacktrack(result *[9][9]int) bool {
	if !EmptyCell(*result) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if result[i][j] == 0 {
				for k := 9; k >= 1; k-- {
					result[i][j] = k
					if isTableValid(*result) {
						if Backtrack(result) {
							return true
						}
						result[i][j] = 0
					} else {
						result[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func EmptyCell(result [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if result[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func HasDuplicate(counter [10]int) bool {
	for i := 1; i < 10; i++ {
		if counter[i] > 1 {
			return true
		}
	}
	return false
}

func isTableValid(result [9][9]int) bool {
	for i := 0; i < 9; i++ {
		counter := [10]int{}
		for j := 0; j < 9; j++ {
			counter[result[i][j]]++
		}
		if HasDuplicate(counter) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		counter := [10]int{}
		for j := 0; j < 9; j++ {
			counter[result[j][i]]++
		}
		if HasDuplicate(counter) {
			return false
		}
	}
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[result[row][col]]++
				}
				if HasDuplicate(counter) {
					return false
				}
			}
		}
	}
	return true

}

func main() {

	arguments := os.Args
	arglen := len(arguments)
	if arglen == 10 {
		var result [9][9]int
		if Parsing(arguments, &result) {
			revResult := result
			if Backtrack(&result) {
				if revBacktrack(&revResult) {
					if result == revResult {
						for i := 0; i < 9; i++ {
							for j := 0; j < 9; j++ {
								fmt.Print(result[i][j])
								if j != 8 {
									fmt.Print(" ")

								}
							}
							fmt.Println()
						}
						return
					}
				}

			}

		}
	}
	fmt.Println("Error")
}
