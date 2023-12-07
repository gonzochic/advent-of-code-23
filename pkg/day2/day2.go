package day2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func PartA() int {
	file, err := os.Open("pkg/day2/input.txt")
	if err != nil {
		panic("Input file not found")
	}

	scanner := bufio.NewScanner(file)

	var selections []cubeSelection

	for scanner.Scan() {
		selections = append(selections, readCubesFromLine(scanner.Text()))
	}

	possibleGames := calculatePossibleGames(14, 12, 13, selections)

	return possibleGames
}

func PartB() int {
	file, err := os.Open("pkg/day2/input.txt")
	if err != nil {
		panic("Input file not found")
	}

	scanner := bufio.NewScanner(file)

	var selections []cubeSelection

	for scanner.Scan() {
		selections = append(selections, readCubesFromLine(scanner.Text()))
	}

	powerOfGames := calculatePowerOfGames(selections)

	return powerOfGames
}

type cubeSelection struct {
	red   int
	blue  int
	green int
	id    int
}

func parseMaxCubesForColor(matches [][]string) int {
	cubes := 0
	for _, match := range matches {
		currentCubeNumber, _ := strconv.Atoi(match[1])
		if currentCubeNumber > cubes {
			cubes = currentCubeNumber
		}
	}
	return cubes
}

func readCubesFromLine(line string) cubeSelection {
	idRx := regexp.MustCompile(`Game (\d+):`)
	id, _ := strconv.Atoi(idRx.FindStringSubmatch(line)[1])

	blueRx := regexp.MustCompile(`(\d+) blue`)
	blueCubes := blueRx.FindAllStringSubmatch(line, -1)
	blue := parseMaxCubesForColor(blueCubes)

	redRx := regexp.MustCompile(`(\d+) red`)
	redCubes := redRx.FindAllStringSubmatch(line, -1)
	red := parseMaxCubesForColor(redCubes)

	greenRx := regexp.MustCompile(`(\d+) green`)
	greenCubes := greenRx.FindAllStringSubmatch(line, -1)
	green := parseMaxCubesForColor(greenCubes)

	res := cubeSelection{id: id, blue: blue, red: red, green: green}
	return res
}

func calculatePossibleGames(blueLim, redLim, greenLim int, games []cubeSelection) int {
	possibleGames := 0
	for _, v := range games {
		if v.blue <= blueLim && v.red <= redLim && v.green <= greenLim {
			possibleGames += v.id
		}
	}
	return possibleGames
}

func calculatePowerOfGames(games []cubeSelection) int {
	powerOfGames := 0
	for _, v := range games {
		powerOfGames += v.red * v.blue * v.green
	}
	return powerOfGames
}
