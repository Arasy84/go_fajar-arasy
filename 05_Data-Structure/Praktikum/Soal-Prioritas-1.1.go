package main

import "fmt"

func ArrayMerge(a, b []string) []string {
	var result []string
	if len(a) == 0 && len(b) == 0 {
		return a
	}

	if len(a) == 0 {
		return b
	}

	if len(b) == 0 {
		return a
	}
	result = append(result, a...)

	for _, v := range b {
		if !contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func contains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
