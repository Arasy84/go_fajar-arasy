package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	counts := make(map[string]int)
	recursiveMostAppearItem(items, len(items)-1, counts)

	var pairs []pair
	for name, count := range counts {
		pairs = append(pairs, pair{name, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})

	return pairs
}

func countElement(arr []string, target string, index int) int {
	if index < 0 {
		return 0
	}
	count := 0
	if arr[index] == target {
		count = 1
	}
	return count + countElement(arr, target, index-1)
}

func recursiveMostAppearItem(items []string, index int, counts map[string]int) {
	if index < 0 {
		return
	}

	item := items[index]
	counts[item]++
	recursiveMostAppearItem(items, index-1, counts)
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}
