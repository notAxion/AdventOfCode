package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input")

	fmt.Println(part1(file))

	fmt.Println(part2(file))
}

func part1(file []byte) int {
	total, num := 0, 0

	for i := 0; i < len(file); {
		ldigit := 0

		for ; i < len(file); i++ {
			if file[i] > '0' && file[i] <= '9' {
				ldigit = int(file[i] - '0')
				break
			}
		}

		num = ldigit * 10

		for ; i < len(file); i++ {
			if file[i] > '0' && file[i] <= '9' {
				ldigit = int(file[i] - '0')
			}
			if file[i] == '\n' {
				break
			}
		}

		total += num + ldigit
	}

	return total
}

func part2(file []byte) int {
	total, num := 0, 0

	for i := 0; i < len(file); {
		ldigit := 0

		for ; i < len(file); i++ {
			if file[i] > '0' && file[i] <= '9' {
				ldigit = int(file[i] - '0')
				break
			}
			if strings.ContainsAny(string(file[i]), "otfsen") {
				ldigit = lettersToDigit(file, i)
				if ldigit > 0 {
					break
				}
			}
		}

		num = ldigit * 10

		for ; i < len(file); i++ {
			if file[i] > '0' && file[i] <= '9' {
				ldigit = int(file[i] - '0')
			}
			if strings.ContainsAny(string(file[i]), "otfsen") {
				tempD := lettersToDigit(file, i)
				if tempD > 0 {
					ldigit = tempD
				}
			}
			if file[i] == '\n' {
				break
			}
		}

		total += num + ldigit
	}

	return total
}

func lettersToDigit(file []byte, i int) int {
	letters := ""
	digits := []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	for j := i; j < i+6 && j < len(file); j++ {
		letters += string(file[j])
	}

	for i, digit := range digits {
		if strings.HasPrefix(letters, digit) {
			return i + 1
		}
	}

	return 0
}
