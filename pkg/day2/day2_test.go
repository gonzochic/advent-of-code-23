package day2

import "testing"

func TestPartA(t *testing.T) {
	testReadCubesFromLine(t, "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", cubeSelection{id: 1, blue: 6, red: 4, green: 2})
	testReadCubesFromLine(t, "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", cubeSelection{id: 2, blue: 4, red: 1, green: 3})
	testReadCubesFromLine(t, "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", cubeSelection{id: 3, blue: 6, red: 20, green: 13})
	testReadCubesFromLine(t, "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", cubeSelection{id: 4, blue: 15, red: 14, green: 3})
	testReadCubesFromLine(t, "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", cubeSelection{id: 5, blue: 2, red: 6, green: 3})
}

func testReadCubesFromLine(t *testing.T, input string, expectedResult cubeSelection) {
	result := readCubesFromLine(input)
	if result != expectedResult {
		t.Fatalf("ReadCubess from %s = %v, expected %v", input, result, expectedResult)
	}
}

func testCalculatePossibleGames(t *testing.T) {
	cubeSelection := []cubeSelection{
		{id: 1, blue: 6, red: 4, green: 2},
		{id: 2, blue: 4, red: 1, green: 3},
		{id: 3, blue: 6, red: 20, green: 13},
		{id: 4, blue: 15, red: 14, green: 3},
		{id: 5, blue: 2, red: 6, green: 3},
	}
	possibleGames := calculatePossibleGames(14, 12, 13, cubeSelection)
	expected := 8
	if possibleGames != expected {
		t.Fatalf("Possible Games: %d; expected: %d", possibleGames, expected)
	}
}

func testCalculatePowerOfGames(t *testing.T) {
	cubeSelection := []cubeSelection{
		{id: 1, blue: 6, red: 4, green: 2},
		{id: 2, blue: 4, red: 1, green: 3},
		{id: 3, blue: 6, red: 20, green: 13},
		{id: 4, blue: 15, red: 14, green: 3},
		{id: 5, blue: 2, red: 6, green: 3},
	}
	possibleGames := calculatePowerOfGames(cubeSelection)
	expected := 2286
	if possibleGames != expected {
		t.Fatalf("Possible Games: %d; expected: %d", possibleGames, expected)
	}
}
