package main

import "fmt"

func MaxSequence(arr []int) int {
	return maxSequenceRecursive(arr, 0, len(arr)-1)
}

func maxSequenceRecursive(arr []int, left, right int) int {
	if left == right {
		return arr[left]
	}

	mid := (left + right) / 2

	leftMax := maxSequenceRecursive(arr, left, mid)
	rightMax := maxSequenceRecursive(arr, mid+1, right)
	crossMax := maxCrossingSequence(arr, left, mid, right)

	return max(leftMax, rightMax, crossMax)
}

func maxCrossingSequence(arr []int, left, mid, right int) int {
	leftSum := 0
	leftMax := arr[mid]
	for i := mid; i >= left; i-- {
		leftSum += arr[i]
		if leftSum > leftMax {
			leftMax = leftSum
		}
	}

	rightSum := 0
	rightMax := arr[mid+1]
	for i := mid + 1; i <= right; i++ {
		rightSum += arr[i]
		if rightSum > rightMax {
			rightMax = rightSum
		}
	}

	return leftMax + rightMax
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
