package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var totalScore int
	for _, v := range s.score {
		totalScore += v
	}
	return float64(totalScore) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i, v := range s.score {
		if v < min {
			min = v
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i, v := range s.score {
		if v > max {
			max = v
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input  " + " Student`s Name :  ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input  " + name + " Score :  ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is ", a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is :  "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is :  "+nameMin+" (", scoreMin, ")")
}
