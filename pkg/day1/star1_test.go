package day1

import "testing"

func TestFindCV(t *testing.T) {
	line := [4]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"134treb7uchet",
	}
	expectedResult := [4]int{
		12,
		38,
		15,
		17,
	}

	for i, v := range line {
		r := parseCVinLine(v)

		if r != expectedResult[i] {
			t.Fatalf("CV for %s = %d, wanted %d", v, r, expectedResult[i])
		}
	}
}

func TestFindCVWithTextNumbers(t *testing.T) {
	line := [9]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
		"pbkprbzvs819threeonekjpk7brkmbqbkgroneightb",
		"bdtwone9fourqdlhsfmstwo",
	}
	expectedResult := [9]int{
		29,
		83,
		13,
		24,
		42,
		14,
		76,
		88,
		22,
	}

	for i, v := range line {
		r := parceCVinLineWithText(v)

		if r != expectedResult[i] {
			t.Fatalf("CV for %s = %d, wanted %d", v, r, expectedResult[i])
		}
	}
}
