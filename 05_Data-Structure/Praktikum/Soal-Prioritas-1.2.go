package main

import "fmt"

func Mapping(slice []string) map[string]int {
	result := make(map[string]int)
	for _, element := range slice {
		if count, ok := result[element]; ok {
			result[element] = count + 1
		} else {
			result[element] = 1
		}
	}
	return result
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
