package main

import (
	"fmt"
	"os"
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
	defer file.Close()

	// reader := strings.NewReader(file.)
	// fmt.Println(reader.Len())
on:
	for i := int64(0); ; i++ {
		markerByte := make([]byte, 4)
		file.ReadAt(markerByte, i)
		marker := string(markerByte)
		// fmt.Println(string(marker))
		for j := 0; j < len(marker)-1; j++ {
			ch := string(marker[j])
			// fmt.Println(marker[j+1:], ch)
			if strings.Contains(marker[j+1:], ch) {
				continue on
			}
		}
		fmt.Println(marker, i+4)
		return
	}
}

func part2() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// reader := strings.NewReader(file.)
	// fmt.Println(reader.Len())
on:
	for i := int64(0); ; i++ {
		markerByte := make([]byte, 14)
		file.ReadAt(markerByte, i)
		marker := string(markerByte)
		// fmt.Println(string(marker))
		for j := 0; j < len(marker)-1; j++ {
			ch := string(marker[j])
			// fmt.Println(marker[j+1:], ch)
			if strings.Contains(marker[j+1:], ch) {
				continue on
			}
		}
		fmt.Println(marker, i+14)
		return
	}
}
