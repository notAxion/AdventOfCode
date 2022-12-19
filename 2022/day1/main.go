package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	localMax, max := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			if localMax > max {
				max = localMax
			}
			localMax = 0
			continue
		}

		food, _ := strconv.Atoi(line)
		localMax += food
	}
	if localMax > max {
		max = localMax
	}
	fmt.Println(max)
}
