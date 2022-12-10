package main

import (
	"bufio"
	"os"

	"github.com/leonests/golinq"
)

func findCommonItem(compartment1, compartment2 golinq.Enumerator[int, rune]) rune {
	return compartment1.Intersect(compartment2).First()
}

func getPriority(item rune) int {
	reference := 'a'
	base := 1
	if item < reference {
		reference = 'A'
		base = 27
	}
	return int(item-reference) + base
}

func extractPriority(items *golinq.Enumerator[int, rune]) int {
	compartment1 := items.Take(items.Count() / 2)
	compartment2 := items.Skip(items.Count() / 2)
	return getPriority(findCommonItem(compartment1, compartment2))
}

type Sack struct {
	items    golinq.Enumerator[int, rune]
	priority int
}

func extractSacks(sc *bufio.Scanner) golinq.Enumerator[int, Sack] {
	var sacks []Sack
	for sc.Scan() {
		sack := &Sack{
			items: golinq.FromString(sc.Text()),
		}
		sack.priority = extractPriority(&sack.items)
		sacks = append(sacks, *sack)
	}
	return golinq.FromSlice(sacks)
}

func prioritiesSum(sacks *golinq.Enumerator[int, Sack]) int64 {
	return sacks.Select(func(i int, s Sack) any {
		return s.priority
	}).Sum2Int()
}

type Group struct {
	sacks    golinq.Enumerator[int, Sack]
	size     int64
	priority int64
}

func groupPrioritiesSum(sacks golinq.Enumerator[int, Sack]) int64 {
	x1 := sacks.ToSlice()
	var groups []Group
	for i := 0; i < int(len(x1)/3); i++ {
		group := &Group{
			sacks: sacks.Take(3),
			size:  3,
		}
		group.priority = extractGroupPriority(group)
		groups = append(groups, *group)
		sacks = sacks.Skip(3)
	}
	x := golinq.FromSlice(groups)
	return x.Select(func(i int, g Group) any {
		return g.priority
	}).Sum2Int()
}

func extractGroupPriority(group *Group) int64 {
	commonItems := group.sacks.First().items
	sacks := group.sacks.Skip(1).ToSlice()
	for i := 0; i < len(sacks); i++ {
		commonItems = commonItems.Intersect(sacks[i].items)
	}
	return int64(getPriority(commonItems.First()))
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	sc := bufio.NewScanner(input)
	sacks := extractSacks(sc)
	println("Part I:", prioritiesSum(&sacks))
	println("Part II:", groupPrioritiesSum(sacks))
}
