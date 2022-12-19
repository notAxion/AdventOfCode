package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input")
	// file, err := os.Open("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rawCrates := make([]string, 0, 9)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rawCrates = append(rawCrates, line)
	}
	crateNums := len(strings.Fields(rawCrates[len(rawCrates)-1]))
	crates := make([][]byte, crateNums)
	rawCrates = rawCrates[:len(rawCrates)-1]
	// fmt.Println(rawCrates)
	for i := len(rawCrates) - 1; i >= 0; i-- {
		line := rawCrates[i]
		// line = strings.ReplaceAll(line, "[", "")
		// line = strings.ReplaceAll(line, "]", "")
		// crateNames := strings.Split(line, " ")
		// fmt.Println(crateNames, len(crates))
		// for i, name := range crateNames {
		for i, j := 1, 0; i < len(line); i += 4 {
			name := line[i]
			if name != ' ' {
				crates[j] = append(crates[j], name)
			}
			j++
		}
	}
	fmt.Println(crates)

	for scanner.Scan() {
		procedure := strings.Fields(scanner.Text())
		move, _ := strconv.Atoi(procedure[1])
		from, _ := strconv.Atoi(procedure[3])
		from--
		to, _ := strconv.Atoi(procedure[5])
		to--
		for i := 0; i < move; i++ {
			lastFrom := len(crates[from]) - 1
			// lastTo := len(crates[to]) - 1
			// fmt.Printf("%v\n", crates)
			if lastFrom < 0 {
				fmt.Println(procedure, "\n", string(crates[from]))
				fmt.Printf("%v\n", crates)
				return
			}
			crate := crates[from][lastFrom]
			crates[from] = crates[from][:lastFrom]
			crates[to] = append(crates[to], crate)
		}
	}

	for _, arr := range crates {
		fmt.Print(string(arr[len(arr)-1]))
	}
	fmt.Println()
}

func part2() {
	file, err := os.Open("input")
	// file, err := os.Open("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rawCrates := make([]string, 0, 9)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rawCrates = append(rawCrates, line)
	}
	crateNums := len(strings.Fields(rawCrates[len(rawCrates)-1]))
	crates := make([][]byte, crateNums)
	rawCrates = rawCrates[:len(rawCrates)-1]
	// fmt.Println(rawCrates)
	for i := len(rawCrates) - 1; i >= 0; i-- {
		line := rawCrates[i]
		// line = strings.ReplaceAll(line, "[", "")
		// line = strings.ReplaceAll(line, "]", "")
		// crateNames := strings.Split(line, " ")
		// fmt.Println(crateNames, len(crates))
		// for i, name := range crateNames {
		for i, j := 1, 0; i < len(line); i += 4 {
			name := line[i]
			if name != ' ' {
				crates[j] = append(crates[j], name)
			}
			j++
		}
	}
	fmt.Println(crates)

	for scanner.Scan() {
		procedure := strings.Fields(scanner.Text())
		move, _ := strconv.Atoi(procedure[1])
		from, _ := strconv.Atoi(procedure[3])
		from--
		to, _ := strconv.Atoi(procedure[5])
		to--

		lenFrom := len(crates[from])
		crateChunk := crates[from][lenFrom-move:]
		crates[from] = crates[from][:lenFrom-move]
		crates[to] = append(crates[to], crateChunk...)

	}

	for _, arr := range crates {
		fmt.Print(string(arr[len(arr)-1]))
	}
	fmt.Println()
}
