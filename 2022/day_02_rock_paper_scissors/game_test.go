package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundResult(t *testing.T) {
	win := 6
	draw := 3
	lose := 0
	var tests = []struct {
		aTurn          rune
		bTurn          rune
		expectedOutput int
	}{
		{aTurn: 'A', bTurn: 'X', expectedOutput: draw + 1}, // 1-1=0
		{aTurn: 'B', bTurn: 'X', expectedOutput: lose + 1}, // 1-2=-1
		{aTurn: 'C', bTurn: 'X', expectedOutput: win + 1},  // 1-3=-2
		{aTurn: 'A', bTurn: 'Y', expectedOutput: win + 2},  // 2-1=1
		{aTurn: 'B', bTurn: 'Y', expectedOutput: draw + 2}, // 2-2=0
		{aTurn: 'C', bTurn: 'Y', expectedOutput: lose + 2}, // 2-3=-1
		{aTurn: 'A', bTurn: 'Z', expectedOutput: lose + 3}, // 3-1=2
		{aTurn: 'B', bTurn: 'Z', expectedOutput: win + 3},  // 3-2=1
		{aTurn: 'C', bTurn: 'Z', expectedOutput: draw + 3}, // 3-3=0
	}

	for _, test := range tests {
		actualOutput := calculateRound(test.aTurn, test.bTurn)
		require.Equal(t, test.expectedOutput, actualOutput)
	}
}

func TestGameResult(t *testing.T) {
	input := `A Y
B X
C Z`
	expectedOutput := 15
	sc := bufio.NewScanner(bytes.NewBufferString(input))
	require.Equal(t, expectedOutput, calculateGame(sc))
}

func TestPlayToEnsureResult(t *testing.T) {
	var tests = []struct {
		input          string
		expectedOutput rune
	}{
		{"A X", 'Z'}, // lose -> scissors
		{"B X", 'X'}, // lose -> rock
		{"C X", 'Y'}, // lose -> paper
		{"A Y", 'X'}, // draw -> rock
		{"B Y", 'Y'}, // draw -> paper
		{"C Y", 'Z'}, // draw -> scissors
		{"A Z", 'Y'}, // win  -> paper
		{"B Z", 'Z'}, // win  -> scissors
		{"C Z", 'X'}, // win  -> rock
	}

	for _, test := range tests {
		require.Equal(t, test.expectedOutput, playToEnsureResult(test.input))
	}
}

func TestMockedGameResult(t *testing.T) {
	input := `A Y
B X
C Z`
	expectedOutput := 12
	sc := bufio.NewScanner(bytes.NewBufferString(input))
	require.Equal(t, expectedOutput, runMockedGame(sc))
}
