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
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	c := 0
	for scanner.Scan() {
		line := scanner.Text()
		idsStr := strings.Split(strings.ReplaceAll(line, ",", "-"), "-")
		ids := make([]int, len(idsStr))
		// all this to have the 4 ids
		// in 1 array of int
		for i := range idsStr {
			ids[i], err = strconv.Atoi(idsStr[i])
			if err != nil {
				fmt.Println(err, line, i)
				return
			}
		}
		if ids[0] <= ids[2] {
			if ids[1] >= ids[3] {
				c++
				// fmt.Print("> ")
			}
		} else if ids[1] <= ids[3] {
			c++
			// fmt.Print("> ")
		}
		if ids[0] == ids[2] && ids[1] < ids[3] {
			c++
			// fmt.Print("> ")
		}
		// fmt.Println(line)
		// if c > 20 {
		// 	fmt.Println(c)
		// 	break
		// }
	}
	fmt.Println(c)
}

func part2() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	c := 0
	for scanner.Scan() {
		line := scanner.Text()
		idsStr := strings.Split(strings.ReplaceAll(line, ",", "-"), "-")
		ids := make([]int, len(idsStr))
		// all this to have the 4 ids
		// in 1 array of int
		for i := range idsStr {
			ids[i], err = strconv.Atoi(idsStr[i])
			if err != nil {
				fmt.Println(err, line, i)
				return
			}
		}

		if ids[1] < ids[2] || ids[3] < ids[0] {
		} else {
			c++
		}
		// fmt.Println(line)
		// if c > 20 {
		// 	fmt.Println(c)
		// 	break
		// }
	}
	fmt.Println(c)
}
