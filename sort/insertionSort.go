package main

import (
	"fmt"
	"math/rand"
)

func sort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			currentIndex := i + 1
			indexAt := findIndex(*arr, (*arr)[currentIndex])
			insertValue(arr, (*arr)[currentIndex], indexAt, currentIndex)
		}
	}
}
func findIndex(arr []int, item int) int {
	for i := 0; i < len(arr); i++ {
		if item < arr[i] {
			return i
		}
	}
	return 0
}

func insertValue(arr *[]int, item int, indexAt int, indexFrom int) {
	temp := (*arr)[indexAt]
	(*arr)[indexAt] = item
	for i := indexFrom; i > indexAt; i-- {
		(*arr)[i] = (*arr)[i-1]
	}
	(*arr)[indexAt+1] = temp
}
func main() {
	x := make([]int, 5)
	for i := 0; i < 5; i++ {
		x[i] = int(rand.Float32() * 10)
	}
	fmt.Println("Unsorted array: ", x)
	sort(&x)
	fmt.Println("Sorted array: ", x)
}
