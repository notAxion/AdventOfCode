package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input")

	fmt.Println(part1(file))
}

func part1(file []byte) int {
	width := strings.IndexByte(string(file), '\n') + 1
	height := (len(file) + 1) / width
	// creating a matrix with 1 byte padding all around
	// for wiggle room (so that i don't have to worry about
	// over stepping when traversing)
	matrix := make([][]byte, height+2)
	for i := range matrix {
		// (width - 1) is the no. of bytes without the \n
		matrix[i] = make([]byte, (width-1)+2)
		for j := range matrix[i] {
			matrix[i][j] = '.'
		}
	}

	total := 0

	for i, x := 0, 1; i < len(file); i, x = i+1, x+1 {
		for ; i < len(file); i++ {
			y := (i % width) + 1
			if file[i] == '\n' {
				break
			}
			matrix[x][y] = file[i]
		}
	}

	// lines := strings.FieldsFunc(string(file), func(r rune) bool { return r == '\n' })
	// for i := 1; i < len(matrix)-1; i++ {
	// 	if lines
	// }

	for i := 1; i < len(matrix)-1; i++ {
		num, numDigits := 0, 0
		for j := 1; j < len(matrix[i]); j++ {
			val := matrix[i][j]
			if val >= '0' && val <= '9' {
				num = num*10 + int(val-'0')
				numDigits++
				continue
			}

			if num > 0 {
				k := numDigits
				// [i][j]
				// [i+1][j]
				// for numDigits step minus in j
				// then another minus 1 in j; j = j-k-1
				// [i][current j]
				// [i-1][current j]
				// for numDigits step plus in j
				// [i-1][j]
				if matrix[i][j] != '.' ||
					matrix[i+1][j] != '.' ||
					matrix[i+1][j-k-1] != '.' ||
					matrix[i][j-k-1] != '.' ||
					matrix[i-1][j-k-1] != '.' ||
					matrix[i-1][j] != '.' {

					total += num
				} else {
					for x := 1; x <= k; x++ {
						if matrix[i+1][j-x] != '.' || matrix[i-1][j-x] != '.' {
							total += num
							break
						}
					}
				}
				num, numDigits = 0, 0
			}
		}
	}

	return total
}
