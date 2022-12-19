package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	score := make(map[string]int)
	opp, me := "ABC", "XYZ"
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			str := fmt.Sprintf("%c %c", opp[i], me[j])
			sc := 0
			if i == j {
				sc = 3 + j + 1
			} else if i < j {
				if i == 0 && j == 2 {
					// loss
					sc = 0 + j + 1
				} else {
					// win
					sc = 6 + j + 1
				}
			} else {
				if i == 2 && j == 0 {
					// win
					sc = 6 + j + 1
				} else {
					// loss
					sc = 0 + j + 1
				}
			}
			score[str] = sc
		}
	}
	fmt.Println(score)
	// for str, val := range score {
	// 	fmt.Printf("%#v %v\n", str, val)
	// }

	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line, score[line])
		total += score[line]
		if _, ok := score[line]; !ok {
			fmt.Println(line)
		}
	}
	fmt.Println(total)
}

func part2() {
	score := make(map[string]int)
	opp, me := "ABC", "XYZ"
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			str := fmt.Sprintf("%c %c", opp[i], me[j])
			sc := 0
			if j == 1 {
				sc = 3 + i + 1
			} else if j == 2 {
				// win
				if i == 2 {
					sc = 6 + 1
				} else {
					sc = 6 + i + 2
				}
			} else {
				// loss
				if i == 0 {
					sc = 0 + 3
				} else {
					sc = 0 + i
				}
			}
			score[str] = sc
		}
	}
	fmt.Println(score)
	// for str, val := range score {
	// 	fmt.Printf("%#v %v\n", str, val)
	// }

	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line, score[line])
		total += score[line]
		if _, ok := score[line]; !ok {
			fmt.Println(line)
		}
	}
	fmt.Println(total)
}
