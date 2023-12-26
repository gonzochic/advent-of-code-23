package day3

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func PartA() int {
	file, err := os.Open("pkg/day3/input.txt")
	if err != nil {
		panic("Input file not found")
	}

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	parts := parseParts(input)
	connectors := parseAllConnectors(input)
	connectedParts := findConnectedParts(parts, connectors)
	sumOfConnectedParts := 0
	for _, cp := range connectedParts {
		sumOfConnectedParts += cp
	}

	return sumOfConnectedParts
}

func PartB() int {
	file, err := os.Open("pkg/day3/input.txt")
	if err != nil {
		panic("Input file not found")
	}

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	parts := parseParts(input)
	connectors := parseGearRatioConnectors(input)
	gearRatios := findGearRatios(parts, connectors)

	sumOfGearRatios := 0

	for _, gr := range gearRatios {
		sumOfGearRatios += gr
	}
	return sumOfGearRatios
}

func parseAllConnectors(input []string) [][]int {
	var connectorIndexes [][]int

	for _, v := range input {
		rx := regexp.MustCompile(`[-#!$%^&*()_+|~={}\[\]:";'<>?,\/@]`)
		connectorIndexes = append(connectorIndexes, parseConnectorsForRegex(v, rx))
	}
	return connectorIndexes
}

func parseGearRatioConnectors(input []string) [][]int {
	var connectorIndexes [][]int

	for _, v := range input {
		rx := regexp.MustCompile(`[*]`)
		connectorIndexes = append(connectorIndexes, parseConnectorsForRegex(v, rx))
	}
	return connectorIndexes
}

func parseConnectorsForRegex(line string, rx *regexp.Regexp) []int {
	indexMatches := rx.FindAllStringIndex(line, -1)
	var conPerLine []int

	for _, m := range indexMatches {
		conPerLine = append(conPerLine, m[0])
	}
	return conPerLine
}

type part struct {
	value      int
	startIndex int
	endIndex   int
}

func parseParts(input []string) [][]part {
	var parts [][]part
	for _, v := range input {
		var partsPerLine []part
		rx := regexp.MustCompile(`\d+`)
		numberMatches := rx.FindAllString(v, -1)
		indexMatches := rx.FindAllStringIndex(v, -1)

		for i, p := range numberMatches {
			partNumber, _ := strconv.Atoi(p)
			newPart := part{
				value:      partNumber,
				startIndex: indexMatches[i][0],
				endIndex:   indexMatches[i][1],
			}
			partsPerLine = append(partsPerLine, newPart)

		}
		parts = append(parts, partsPerLine)
	}
	return parts
}

func findConnectedParts(parts [][]part, connectors [][]int) []int {
	var connectedParts []int

	for plI, pl := range parts {
		for _, p := range pl {
			// Define search window
			luSearchWindow := [2]int{max(0, plI-1), max(0, p.startIndex-1)}
			rlSearchWindow := [2]int{min(len(parts)-1, plI+1), p.endIndex}

			if isConnectorInSearchWindow(connectors, luSearchWindow, rlSearchWindow) {
				connectedParts = append(connectedParts, p.value)
			}
		}
	}
	return connectedParts
}

func isConnectorInSearchWindow(connectors [][]int, luSW, rlSW [2]int) bool {
	for i := luSW[0]; i <= rlSW[0]; i++ {
		connectorLine := connectors[i]
		for _, c := range connectorLine {
			connectorInWindow := c >= luSW[1] && c <= rlSW[1]
			if connectorInWindow {
				return true
			}
		}
	}

	return false
}

func findGearRatios(parts [][]part, connectors [][]int) []int {
	var gearRatios []int

	for cLineI, cLine := range connectors {
		for _, c := range cLine {
			luSearchWindow := [2]int{max(0, cLineI-1), max(0, c-1)}
			rlSearchWindow := [2]int{min(len(parts)-1, cLineI+1), c + 1}

			partsInWindow := findPartInSearchWindow(parts, luSearchWindow, rlSearchWindow)

			if len(partsInWindow) != 2 {
				// skip parts that have no adjacent part
				continue
			}

			var gearRatio int
			for _, pv := range partsInWindow {
				if gearRatio == 0 {
					gearRatio = pv
					continue
				}

				gearRatio *= pv
			}

			gearRatios = append(gearRatios, gearRatio)
		}
	}

	return gearRatios
}

func findPartInSearchWindow(parts [][]part, luSW, rlSW [2]int) []int {
	var partsInWindow []int
	for i := luSW[0]; i <= rlSW[0]; i++ {
		partLine := parts[i]
		for _, p := range partLine {
			partInWindow := p.startIndex >= luSW[1] && p.startIndex <= rlSW[1] || p.endIndex-1 >= luSW[1] && p.endIndex-1 <= rlSW[1]
			if partInWindow {
				partsInWindow = append(partsInWindow, p.value)
			}
		}
	}
	return partsInWindow
}
