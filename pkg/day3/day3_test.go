package day3

import (
	"bufio"
	"os"
	"testing"
)

func setupTest() []string {
	file, err := os.Open("./test_input.txt")
	if err != nil {
		panic("Input file not found")
	}

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}

func TestFindConnectedParts(t *testing.T) {
	input := setupTest()

	parts := parseParts(input)
	connectors := parseAllConnectors(input)
	cParts := findConnectedParts(parts, connectors)

	expectedResult := 4361

	sumOfConnectedParts := 0
	for _, cp := range cParts {
		sumOfConnectedParts += cp
	}

	if sumOfConnectedParts != expectedResult {
		t.Errorf("Sum of connected parts is %d, should be %d", sumOfConnectedParts, expectedResult)
	}
}

func TestGearRatios(t *testing.T) {
	input := setupTest()

	parts := parseParts(input)
	connectors := parseGearRatioConnectors(input)

	gRatios := findGearRatios(parts, connectors)

	expectedResult := 467835

	sumOfGearRatios := 0
	for _, cp := range gRatios {
		sumOfGearRatios += cp
	}

	if sumOfGearRatios != expectedResult {
		t.Errorf("Sum of gear ratios is %d, should be %d", sumOfGearRatios, expectedResult)
	}
}
