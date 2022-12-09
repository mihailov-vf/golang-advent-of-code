package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElfCargosShouldBeSeparatedByBlankLine(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	elves := parseElves(bufio.NewScanner(bytes.NewBufferString(input)))
	assert.Len(t, elves, 5)
}

func TestElfTotalCalories(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000	
8000
9000

10000`
	expectedCalories := []int{6000, 4000, 11000, 24000, 10000}
	elves := parseElves(bufio.NewScanner(bytes.NewBufferString(input)))

	for i, v := range elves {
		assert.Equal(t, v.TotalCalories(), expectedCalories[i])
	}
}
