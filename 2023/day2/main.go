package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.ReadFile("input")

	fmt.Println(part1(file))
	fmt.Println(part2(file))
}

func part1(file []byte) int {
	redMax, greenMax, blueMax := 12, 13, 14
	total := 0

	// each line
line:
	for i := 0; i < len(file); {
		gameRecord := make([]byte, 0, 9)
		for ; i < len(file); i++ {
			if file[i] == ':' {
				i++
				break
			}
			gameRecord = append(gameRecord, file[i])
		}
		gameID := 0
		fmt.Sscanf(string(gameRecord), "Game %d", &gameID)

		// each game
		for ; i < len(file); i++ {
			whichCube := make([]byte, 0, 10)
			for ; i < len(file) && !(file[i] == ';' || file[i] == ',' || file[i] == '\n'); i++ {
				whichCube = append(whichCube, file[i])
			}

			reds, greens, blues := 0, 0, 0
			cubes := 0
			cubeColor := ""
			fmt.Sscanf(string(whichCube), " %d %s", &cubes, &cubeColor)

			switch cubeColor {
			case "red":
				reds = cubes
			case "green":
				greens = cubes
			case "blue":
				blues = cubes
			default:
				panic("what the color - " + cubeColor)
			}

			if reds > redMax || greens > greenMax || blues > blueMax {
				for ; i < len(file); i++ {
					if file[i] == '\n' {
						i++
						break
					}
				}
				continue line
			}
			if i >= len(file) || file[i] == '\n' {
				i++
				break
			}
		}

		total += gameID
	}

	return total
}

func part2(file []byte) int {
	total := 0

	// each line
	for i := 0; i < len(file); {
		gameRecord := make([]byte, 0, 9)
		for ; i < len(file); i++ {
			if file[i] == ':' {
				i++
				break
			}
			gameRecord = append(gameRecord, file[i])
		}
		gameID := 0
		fmt.Sscanf(string(gameRecord), "Game %d", &gameID)

		redLMax, greenLMax, blueLMax := 0, 0, 0
		// each game
		for ; i < len(file); i++ {
			whichCube := make([]byte, 0, 10)
			for ; i < len(file) && !(file[i] == ';' || file[i] == ',' || file[i] == '\n'); i++ {
				whichCube = append(whichCube, file[i])
			}

			reds, greens, blues := 0, 0, 0
			cubes := 0
			cubeColor := ""
			fmt.Sscanf(string(whichCube), " %d %s", &cubes, &cubeColor)

			switch cubeColor {
			case "red":
				reds = cubes
			case "green":
				greens = cubes
			case "blue":
				blues = cubes
			default:
				panic("what the color - " + cubeColor)
			}

			if reds > redLMax {
				redLMax = reds
			}

			if greens > greenLMax {
				greenLMax = greens
			}

			if blues > blueLMax {
				blueLMax = blues
			}
			/* if reds > redMax || greens > greenMax || blues > blueMax {
				for ; i < len(file); i++ {
					if file[i] == '\n' {
						i++
						break
					}
				}
				continue line
			} */
			if i >= len(file) || file[i] == '\n' {
				i++
				break
			}
		}

		total += redLMax * greenLMax * blueLMax
	}

	return total
}
