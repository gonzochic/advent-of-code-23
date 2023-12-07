package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartA() int {
	file, err := os.Open("pkg/day1/input.txt")
	if err != nil {
		panic("Input file not found")
	}

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += parseCVinLine(scanner.Text())
	}

	return result
}

func PartB() int {
	file, err := os.Open("pkg/day1/input.txt")
	if err != nil {
		panic("Input file not found")
	}

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += parceCVinLineWithText(scanner.Text())
	}

	return result
}

var NumberText = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func parceCVinLineWithText(line string) int {
	rx := regexp.MustCompile(`\d+`)
	matches := rx.FindAllStringIndex(line, -1)

	firstNumber := 0
	firstIndex := 99999
	lastNumber := 0
	lastIndex := -9999

	if matches != nil {
		firstIndex = matches[0][0]
		firstNumber, _ = strconv.Atoi(line[firstIndex : firstIndex+1])
		lastIndex = matches[len(matches)-1][0]
		lastNumber, _ = strconv.Atoi(line[lastIndex:matches[len(matches)-1][1]])
	}

	for i, v := range NumberText {
		fIndex := strings.Index(line, i)
		lIndex := strings.LastIndex(line, i)

		if fIndex >= 0 {
			if fIndex < firstIndex {
				firstIndex = fIndex
				firstNumber = v
			}

			if lIndex > lastIndex {
				lastIndex = lIndex
				lastNumber = v
			}
		}
	}

	return firstNumber*10 + lastNumber%10
}

func parseCVinLine(line string) int {
	rx := regexp.MustCompile(`\d+`)
	matches := strings.Join(rx.FindAllString(line, -1), "")
	firstAndLast := fmt.Sprintf("%c%c", matches[0], matches[len(matches)-1])
	res, err := strconv.Atoi(firstAndLast)
	if err != nil {
		return 0
	}

	return res
}
