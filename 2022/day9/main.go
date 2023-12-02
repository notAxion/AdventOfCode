package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

type pos struct {
	x, y int
}

func (p pos) distance(p2 pos) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p.x), 2) + math.Pow(float64(p2.y-p.y), 2))
}

type brain struct {
	h, t  pos
	visit map[pos]bool
}

func (rope *brain) moveHead(direction byte, step int) {
	for i := 0; i < step; i++ {
		switch direction {
		case 'U':
			rope.h.y++
		case 'D':
			rope.h.y--
		case 'L':
			rope.h.x--
		case 'R':
			rope.h.x++
		}
		rope.moveTail(direction)

	}
}

func (rope *brain) moveTail(direction byte) {
	dist := rope.h.distance(rope.t)
	if dist < 2 {
		return
	}
	switch direction {
	case 'U':
		rope.t.y++
	case 'D':
		rope.t.y--
	case 'L':
		rope.t.x--
	case 'R':
		rope.t.x++
	}

	switch direction {
	case 'U', 'D':
		rope.t.x = rope.h.x
	case 'L', 'R':
		rope.t.y = rope.h.y
	}
	rope.visit[rope.t] = true
	if dist != 2.0 && dist != math.Sqrt(5) {
		fmt.Printf("\"%#v\"\n", dist)
		log.Fatalln(rope)
	}

}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bheja := brain{
		visit: make(map[pos]bool),
	}
	bheja.visit[bheja.t] = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		step, err := strconv.Atoi(line[2:])
		if err != nil {
			fmt.Println(err)
			return
		}

		bheja.moveHead(line[0], step)
	}

	fmt.Println(len(bheja.visit))
}

type fallingBrain struct {
	h     pos
	knots []pos
	visit map[pos]bool
}

func (rope *fallingBrain) moveHead(direction byte, step int) {
	for i := 0; i < step; i++ {
		switch direction {
		case 'U':
			rope.h.y++
		case 'D':
			rope.h.y--
		case 'L':
			rope.h.x--
		case 'R':
			rope.h.x++
		}
		rope.moveTail(direction, 0)

	}
}

func (rope *fallingBrain) moveTail(direction byte, index int) {
	if index == len(rope.knots) {
		return
	}

	knot := &rope.knots[index]
	prevKnot := pos{}
	if index == 0 {
		prevKnot = rope.h
	} else {
		prevKnot = rope.knots[index-1]
	}
	dist := prevKnot.distance(*knot)
	if dist < 2 {
		return
	}

	switch direction {
	case 'U':
		knot.y++
	case 'D':
		knot.y--
	case 'L':
		knot.x--
	case 'R':
		knot.x++
	}

	// rope.moveTail(direction, index+1)

	switch direction {
	case 'U', 'D':
		if prevKnot.x-knot.x > 0 {
			direction = 'R'
		} else if prevKnot.x-knot.x < 0 {
			direction = 'L'
		}
		knot.x = prevKnot.x
	case 'L', 'R':
		if prevKnot.y-knot.y > 0 {
			direction = 'U'
		} else if prevKnot.y-knot.y < 0 {
			direction = 'D'
		}

		knot.y = prevKnot.y
	}

	if dist != 2.0 && dist != math.Sqrt(5) {
		fmt.Printf("\"%#v\"\n", dist)
		log.Fatalln(rope)
	}

	if index == len(rope.knots)-1 {
		rope.visit[*knot] = true
	}
	rope.moveTail(direction, index+1)

}

func part2() {
	file, err := os.Open("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bheja := fallingBrain{
		knots: make([]pos, 9),
		visit: make(map[pos]bool),
	}
	bheja.visit[pos{}] = true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		step, err := strconv.Atoi(line[2:])
		if err != nil {
			fmt.Println(err)
			return
		}

		bheja.moveHead(line[0], step)
	}

	fmt.Println(len(bheja.visit))
}
