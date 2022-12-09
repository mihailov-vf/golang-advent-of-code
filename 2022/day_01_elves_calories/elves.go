package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	calories int
}

func (e *Elf) TotalCalories() int {
	return e.calories
}

func (e *Elf) AddFood(food int) {
	e.calories += food
}

func parseElves(sc *bufio.Scanner) []*Elf {
	var elves []*Elf
	elf := &Elf{}
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			elves = append(elves, elf)
			elf = &Elf{}
			continue
		}
		foodCalories, _ := strconv.Atoi(line)
		elf.AddFood(foodCalories)
	}
	elves = append(elves, elf)
	return elves
}

func topSum(elves []*Elf, number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += elves[i].TotalCalories()
	}
	return sum
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	sc := bufio.NewScanner(input)
	elves := parseElves(sc)
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories() > elves[j].TotalCalories()
	})

	println("Top 1:", elves[0].TotalCalories())
	println("Top 3:", topSum(elves, 3))
}
