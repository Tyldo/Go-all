package main

import (
	"fmt"
	"math/rand"
)

func swap(el, elem *int) {
	temp := *el
	*el = *elem
	*elem = temp
}
func main() {
	arr := make([]int, 20)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(200)
	}
	fmt.Printf("Unsorted array: %v \n", arr)
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] < arr[i+1] {
				swap(&arr[i], &arr[i+1])
				flag = true
			}
		}
	}
	fmt.Printf("Sorted array: %v \n", arr)
}
