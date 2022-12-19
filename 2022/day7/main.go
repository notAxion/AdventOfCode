package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	size int64
	dirs []string
}

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

	// qu = make([]string, 0, 10)
	tree := make(map[string]*Directory)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	folder := "/"
	scanner.Scan()
	dir := &Directory{
		dirs: []string{"/", "/"},
	}
	tree[folder] = dir
	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Fields(line)
		switch ff[0] {
		case "dir":
			listDirectory := dir.dirs[0] + ff[1] + "/"
			dir.dirs = append(dir.dirs, listDirectory)
		case "$":
			if ff[1] == "cd" {
				if ff[2] == ".." {
					folder = dir.dirs[1]
					tree[folder].size += dir.size
					dir = tree[folder]
				} else {
					folder = dir.dirs[0] + ff[2] + "/"
					if _, ok := tree[folder]; ok {
						fmt.Println(folder)
					}
				}
			} else if ff[1] == "ls" {
				dir = &Directory{
					dirs: []string{folder, dir.dirs[0]},
				}
				tree[folder] = dir
			}

		default:
			if len(ff) < 2 {
				continue
			}
			size, err := strconv.Atoi(ff[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			dir.size += int64(size)
		}
	}
	for {

		folder = dir.dirs[1]
		tree[folder].size += dir.size
		dir = tree[folder]
		if folder == "/" {
			break
		}
	}
	fmt.Println(tree["/"].size)
	fmt.Println()
	// c := 0
	workDirs := make([]int, 0, 26)
	for _, dir := range tree {
		if dir.size < 100000 {
			// fmt.Println(c)
			// c++
			workDirs = append(workDirs, int(dir.size))
		}
	}
	sort.Ints(workDirs)
	sum := 0
	for i := len(workDirs) - 1; i >= 0; i-- {
		size := int(workDirs[i])
		sum += size
		fmt.Println(size, sum)
	}

	sum = 0
	for i := 0; i < len(workDirs); i++ {
		size := int(workDirs[i])
		sum += size
		fmt.Println(size, sum)
	}
}

func part2() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	// qu = make([]string, 0, 10)
	tree := make(map[string]*Directory)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	folder := "/"
	scanner.Scan()
	dir := &Directory{
		dirs: []string{"/", "/"},
	}
	tree[folder] = dir
	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Fields(line)
		switch ff[0] {
		case "dir":
			listDirectory := dir.dirs[0] + ff[1] + "/"
			dir.dirs = append(dir.dirs, listDirectory)
		case "$":
			if ff[1] == "cd" {
				if ff[2] == ".." {
					folder = dir.dirs[1]
					tree[folder].size += dir.size
					dir = tree[folder]
				} else {
					folder = dir.dirs[0] + ff[2] + "/"
					if _, ok := tree[folder]; ok {
						fmt.Println(folder)
					}
				}
			} else if ff[1] == "ls" {
				dir = &Directory{
					dirs: []string{folder, dir.dirs[0]},
				}
				tree[folder] = dir
			}

		default:
			if len(ff) < 2 {
				continue
			}
			size, err := strconv.Atoi(ff[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			dir.size += int64(size)
		}
	}
	for {

		folder = dir.dirs[1]
		tree[folder].size += dir.size
		dir = tree[folder]
		if folder == "/" {
			break
		}
	}
	fmt.Println(tree["/"].size)
	freeSpace := 70000000 - tree["/"].size
	needMore := 30000000 - freeSpace
	fmt.Println("need this more", needMore)
	fmt.Println()
	// c := 0
	workDirs := make([]int, 0, 26)
	for _, dir := range tree {
		workDirs = append(workDirs, int(dir.size))
	}
	sort.Ints(workDirs)
	for _, size := range workDirs {
		if size > int(needMore) {
			fmt.Println(size)
			break
		}
	}
	// sum := 0
	// for i := len(workDirs) - 1; i >= 0; i-- {
	// 	size := int(workDirs[i])
	// 	sum += size
	// 	fmt.Println(size, sum)
	// }

	// sum = 0
	// for i := 0; i < len(workDirs); i++ {
	// 	size := int(workDirs[i])
	// 	sum += size
	// 	fmt.Println(size, sum)
	// }
}
