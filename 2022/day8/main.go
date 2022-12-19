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
	file, err := os.Open("input")
	// file, err := os.Open("example")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	trees := make([][]int, 0, 50)
	for scanner.Scan() {
		line := scanner.Bytes()
		trees = append(trees, make([]int, len(line)))
		for i, tree := range line {
			trees[len(trees)-1][i] = int(tree - '0')
		}
	}
	// fmt.Println(len(trees), len(trees[0]))
	visible := make(map[string]bool)
	for i := range trees {
		tree := fmt.Sprintf("%dx%d", 0, i)
		visible[tree] = true
		tree = fmt.Sprintf("%dx%d", i, 0)
		visible[tree] = true
		tree = fmt.Sprintf("%dx%d", len(trees)-1, i)
		visible[tree] = true
		tree = fmt.Sprintf("%dx%d", i, len(trees)-1)
		visible[tree] = true
	}

	for i := range trees {
		row := trees[i]
		rowMax := row[0]
		colMax := trees[0][i]

		walkRows, walkCols := true, true
		for j, tree := range row {
			if walkRows {
				var walk bool
				rowMax, walk = traverse(i, j, tree, rowMax, visible)
				walkRows = walkRows && walk
			}
			if walkCols {
				var walk bool
				colMax, walk = traverse(j, i, trees[j][i], colMax, visible)
				walkCols = walkCols && walk
			}
			if !walkCols && !walkRows {
				break
			}
		}

		rowMax = row[len(row)-1]
		colMax = trees[len(row)-1][i]
		walkRows, walkCols = true, true
		for j := len(row) - 1; j >= 0; j-- {
			tree := row[j]
			if walkRows {
				var walk bool
				rowMax, walk = traverse(i, j, tree, rowMax, visible)
				walkRows = walkRows && walk
			}
			if walkCols {
				var walk bool
				colMax, walk = traverse(j, i, trees[j][i], colMax, visible)
				walkCols = walkCols && walk
			}
			if !walkCols && !walkRows {
				break
			}
		}

	}
	fmt.Println(len(visible))
	// fmt.Printf("%#v\n", visible)
	// fmt.Println(traverse(3, 2, 5, 3, visible))
}

func traverse(i, j, tree, lMax int, visible map[string]bool) (int, bool) {
	key := fmt.Sprintf("%dx%d", i, j)
	if visible[key] {
		if tree > lMax {
			lMax = tree
		}
		return lMax, true
	}

	if key == "3x3" {
		// fmt.Println(i, j, lMax, visible)
	}

	if tree > lMax {
		visible[key] = true
		lMax = tree
	}
	if tree == 5 && i == 3 && j == 2 {
		// fmt.Println(i, j, lMax, visible)
	}
	if lMax == 9 {
		return lMax, false
	}

	return lMax, true
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
	trees := make([][]int, 0, 50)
	for scanner.Scan() {
		line := scanner.Bytes()
		trees = append(trees, make([]int, len(line)))
		for i, tree := range line {
			trees[len(trees)-1][i] = int(tree - '0')
		}
	}
	// fmt.Println(len(trees), len(trees[0]))
	visible := make(map[string]bool)
	for i := range trees {
		tree := fmt.Sprintf("%dx%d", 0, i)
		visible[tree] = true
		tree = fmt.Sprintf("%dx%d", i, 0)
		visible[tree] = true
		tree = fmt.Sprintf("%dx%d", len(trees)-1, i)
		visible[tree] = true
		tree = fmt.Sprintf("%dx%d", i, len(trees)-1)
		visible[tree] = true
	}

	for i := range trees {
		row := trees[i]
		rowMax := row[0]
		colMax := trees[0][i]

		walkRows, walkCols := true, true
		for j, tree := range row {
			if walkRows {
				var walk bool
				rowMax, walk = traverse(i, j, tree, rowMax, visible)
				walkRows = walkRows && walk
			}
			if walkCols {
				var walk bool
				colMax, walk = traverse(j, i, trees[j][i], colMax, visible)
				walkCols = walkCols && walk
			}
			if !walkCols && !walkRows {
				break
			}
		}

		rowMax = row[len(row)-1]
		colMax = trees[len(row)-1][i]
		walkRows, walkCols = true, true
		for j := len(row) - 1; j >= 0; j-- {
			tree := row[j]
			if walkRows {
				var walk bool
				rowMax, walk = traverse(i, j, tree, rowMax, visible)
				walkRows = walkRows && walk
			}
			if walkCols {
				var walk bool
				colMax, walk = traverse(j, i, trees[j][i], colMax, visible)
				walkCols = walkCols && walk
			}
			if !walkCols && !walkRows {
				break
			}
		}

	}
	fmt.Println(len(visible))
	// fmt.Printf("%#v\n", visible)
	// fmt.Println(traverse(3, 2, 5, 3, visible))
}
