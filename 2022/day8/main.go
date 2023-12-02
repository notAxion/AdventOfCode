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
	// file, err := os.Open("example")
	file, err := os.Open("input")
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
	view := make(map[string]int)
	for i := range trees {
		tree := fmt.Sprintf("%dx%d", 0, i)
		view[tree] = -1
		tree = fmt.Sprintf("%dx%d", i, 0)
		view[tree] = -1
		tree = fmt.Sprintf("%dx%d", len(trees)-1, i)
		view[tree] = -1
		tree = fmt.Sprintf("%dx%d", i, len(trees)-1)
		view[tree] = -1
	}

	for i := range trees {
		row := trees[i]
		rowMax := row[0]
		colMax := trees[0][i]
		visi := newVisi()

		// walkRows, walkCols := true, true
		for j, tree := range row {
			key := fmt.Sprintf("%dx%d", i, j)
			if view[key] == 0 {
				view[key] = 1
			}
			if view[key] != -1 {

				if tree > rowMax {
					rowMax = tree
					view[key] *= j
				} else {
					firstBig := visi.look(tree)
					view[key] *= (j - firstBig)
				}
				// fmt.Println(key, view[key])

				visi.update(j, tree)
			}
		}

		visi = newVisi()
		for j := range row {
			tree := trees[j][i]
			key := fmt.Sprintf("%dx%d", j, i)

			if view[key] == 0 {
				view[key] = 1
			}
			if view[key] != -1 {

				if tree > colMax {
					colMax = tree
					view[key] *= j
				} else {
					firstBig := visi.look(tree)
					view[key] *= (j - firstBig)
				}
				// fmt.Println(key, view[key])

				visi.update(j, tree)
			}
		}
		// rowMax = traversePart2(i, j, tree, rowMax, view)
		// colMax = traversePart2(j, i, trees[j][i], colMax, view)

		rowMax = row[len(row)-1]
		colMax = trees[len(row)-1][i]
		visi = newVisi()
		// walkRows, walkCols = true, true
		// rowMax = traversePart2(i, j, tree, rowMax, view)
		// colMax = traversePart2(j, i, trees[j][i], colMax, view)
		for j := len(row) - 1; j >= 0; j-- {
			tree := row[j]
			key := fmt.Sprintf("%dx%d", i, j)

			if view[key] != -1 {

				if tree > rowMax {
					rowMax = tree
					view[key] *= len(row) - 1 - j
				} else {
					firstBig := visi.look(tree)
					view[key] *= (len(row) - 1 - j - firstBig)
				}

				visi.update(len(row)-1-j, tree)
			}
		}

		visi = newVisi()
		for j := len(row) - 1; j >= 0; j-- {
			tree := trees[j][i]
			key := fmt.Sprintf("%dx%d", j, i)

			if view[key] != -1 {

				if tree > colMax {
					colMax = tree
					view[key] *= len(row) - 1 - j
				} else {
					firstBig := visi.look(tree)
					view[key] *= (len(row) - 1 - j - firstBig)
				}

				visi.update(len(row)-1-j, tree)
			}
		}
	}
	// fmt.Println(len(view))
	// fmt.Println(view)
	max := -2
	for _, val := range view {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
	// fmt.Printf("%#v\n", visible)
	// fmt.Println(traversePart2(3, 2, 5, 3, visible))
}

type visibility [10]*int

func newVisi() visibility {
	var v [10]*int
	return v
}

func (v *visibility) look(height int) int {
	max := -1
	for i := height; i < len(v); i++ {
		if v[i] != nil {
			// return *v[i]
			if *v[i] > max {
				max = *v[i]
			}
		}
	}
	if max != -1 {
		return max
	}

	return 0
}

func (v *visibility) update(distance, tree int) {
	v[tree] = &distance
}

// func traversePart2(i, j, tree, lMax int, view map[string]int, visi visibility){
// 	// tree := row[j]
// 	key := fmt.Sprintf("%dx%d", i, j)
//
// 	if view[key] != -1 {
//
// 		if tree > rowMax {
// 			rowMax = tree
// 			view[key] *= len(row) - 1 - j
// 		} else {
// 			firstBig := visi.look(tree)
// 			view[key] *= (len(row) - 1 - j - firstBig)
// 		}
//
// 		visi.update(len(row)-1-j, tree)
// 	}
// }

// func traversePart2(i, j, tree, lMax int, view map[string]int) int {
// 	key := fmt.Sprintf("%dx%d", i, j)
// 	if view[key] == 0 {
// 		view[key] = 1
// 	}
//
// 	if view[key] == -1 {
// 		return lMax
// 	}
//
// 	if tree > lMax {
// 		panic("tree is higher than max")
// 	}
//
// 	if tree > lMax {
// 		view[key] = true
// 		lMax = tree
// 	}
// 	if tree == 5 && i == 3 && j == 2 {
// 		// fmt.Println(i, j, lMax, visible)
// 	}
// 	if lMax == 9 {
// 		return lMax
// 	}
//
// 	return lMax
// }
