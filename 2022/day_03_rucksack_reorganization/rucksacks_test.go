package main

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/leonests/golinq"
	"github.com/stretchr/testify/require"
)

func TestFindCommonItem(t *testing.T) {
	tests := []struct {
		input          golinq.Enumerator[int, rune]
		expectedOutput rune
	}{
		{golinq.FromString("vJrwpWtwJgWrhcsFMMfFFhFp"), 'p'},
		{golinq.FromString("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), 'L'},
		{golinq.FromString("PmmdzqPrVvPwwTWBwg"), 'P'},
		{golinq.FromString("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"), 'v'},
		{golinq.FromString("ttgJtRGJQctTZtZT"), 't'},
		{golinq.FromString("CrZsJsPPZsGzwwsLwLmpwMDw"), 's'},
	}

	for _, v := range tests {

		compartment1 := v.input.Take(v.input.Count() / 2)
		compartment2 := v.input.Skip(v.input.Count() / 2)

		require.Equal(t, v.expectedOutput, findCommonItem(compartment1, compartment2))
	}
}

func TestGetPriority(t *testing.T) {
	tests := []struct {
		input          rune
		expectedOutput int
	}{
		{'p', 16},
		{'L', 38},
		{'P', 42},
		{'v', 22},
		{'t', 20},
		{'s', 19},
	}

	for _, v := range tests {
		require.Equal(t, v.expectedOutput, getPriority(v.input))
	}
}

func TestPrioritiesSum(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	expectedOutput := int64(157)
	sc := bufio.NewScanner(bytes.NewBufferString(input))
	sacks := extractSacks(sc)
	require.Equal(t, expectedOutput, prioritiesSum(&sacks))
}

func createSack(input string) *Sack {
	return &Sack{
		items: golinq.FromString(input),
	}
}

func createGroup(sacks []Sack) *Group {
	return &Group{
		sacks: golinq.FromSlice(sacks),
		size:  int64(len(sacks)),
	}
}

func TestExtractGroupPriority(t *testing.T) {
	tests := []struct {
		input          Group
		expectedOutput int64
	}{
		{input: *createGroup([]Sack{
			*createSack("vJrwpWtwJgWrhcsFMMfFFhFp"),
			*createSack("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"),
			*createSack("PmmdzqPrVvPwwTWBwg"),
		}), expectedOutput: 18},
		{input: *createGroup([]Sack{
			*createSack("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"),
			*createSack("ttgJtRGJQctTZtZT"),
			*createSack("CrZsJsPPZsGzwwsLwLmpwMDw"),
		}), expectedOutput: 52},
	}

	for _, v := range tests {
		require.Equal(t, v.expectedOutput, extractGroupPriority(&v.input))
	}
}

func TestGroupPrioritiesSum(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	expectedOutput := int64(70)
	sc := bufio.NewScanner(bytes.NewBufferString(input))
	sacks := extractSacks(sc)
	require.Equal(t, expectedOutput, groupPrioritiesSum(sacks))
}
