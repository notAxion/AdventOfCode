package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// for i := 0; i < 3; i++ {
		// 	scanner.Scan()
		cache := make(map[byte]int)
		line := scanner.Text()
		llen := len(line)
		for i := range line[:llen/2] {
			ch := line[i]
			// fmt.Println(string(ch))
			cache[ch]++
		}
		for i := range line[llen/2:] {
			i += llen / 2
			ch := line[i]
			// fmt.Println(string(ch))
			if _, ok := cache[ch]; !ok {
				continue
			}
			if ch >= 97 && ch <= 122 {
				sum += int(ch - 'a' + 1)
				// fmt.Println(sum, ch, string(ch), int(ch-'a'+1))
			} else if ch >= 65 && ch <= 90 {
				sum += int(ch - 'A' + 27)
				// fmt.Println(sum, ch, string(ch), int(ch-'A'+27))
			}
			delete(cache, ch)
		}
		// fmt.Println(line[:llen/2], line[llen/2:], sum)
	}
	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("input")
	// file, err := os.Open("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// for i := 0; i < 3; i++ {
		// 	scanner.Scan()
		line := scanner.Text()
		cache := make(map[byte]int)
		cache2 := make(map[byte]int)
		// llen := len(line)
		for i := range line {
			ch := line[i]
			// fmt.Println(string(ch))
			cache[ch]++
		}

		scanner.Scan()
		line = scanner.Text()
		for i := range line {
			ch := line[i]
			if _, ok := cache[ch]; ok {
				cache2[ch]++
			}
		}

		scanner.Scan()
		line = scanner.Text()
		c := 0
		for i := range line {
			// i += llen / 2
			ch := line[i]
			// fmt.Println(string(ch))
			if _, ok := cache2[ch]; !ok {
				continue
			}
			if ch >= 97 && ch <= 122 {
				sum += int(ch - 'a' + 1)
				// fmt.Println(sum, ch, string(ch), int(ch-'a'+1))
			} else if ch >= 65 && ch <= 90 {
				sum += int(ch - 'A' + 27)
				// fmt.Println(sum, ch, string(ch), int(ch-'A'+27))
			}
			delete(cache2, ch)
			c++
			if c > 1 {
				for k, v := range cache2 {
					fmt.Print(string(k), ":", v, ",")
				}
				fmt.Println()
				fmt.Println(string(ch))
				log.Fatalln(line)
			}
		}
		// fmt.Println(line[:llen/2], line[llen/2:], sum)
	}
	fmt.Println(sum)
}
