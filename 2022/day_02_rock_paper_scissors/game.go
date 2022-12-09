package main

import (
	"bufio"
	"io"
	"os"
)

func calculateRound(a, b rune) int {
	win := 6
	draw := 3
	values := map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'X': 1,
		'Y': 2,
		'Z': 3,
	}

	result := values[b] - values[a]
	if result == 0 {
		return draw + values[b]
	}
	if result == -2 || result == 1 {
		return win + values[b]
	}
	return values[b]
}

func calculateGame(sc *bufio.Scanner) int {
	result := 0
	for sc.Scan() {
		input := []rune(sc.Text())
		result += calculateRound(input[0], input[2])
	}
	return result
}

func playToEnsureResult(input string) rune {
	mappedInput := map[string]rune{
		"A X": 'Z',
		"B X": 'X',
		"C X": 'Y',
		"A Y": 'X',
		"B Y": 'Y',
		"C Y": 'Z',
		"A Z": 'Y',
		"B Z": 'Z',
		"C Z": 'X',
	}
	return mappedInput[string(input)]
}

func runMockedGame(sc *bufio.Scanner) int {
	result := 0
	for sc.Scan() {
		input := []rune(sc.Text())
		result += calculateRound(input[0], playToEnsureResult(sc.Text()))
	}
	return result
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	println("Part I:", calculateGame(bufio.NewScanner(input)))
	_, err = input.Seek(0, io.SeekStart)
	if err != nil {
		panic(err.Error())
	}
	println("Part II:", runMockedGame(bufio.NewScanner(input)))
}
