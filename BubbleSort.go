package main

import "fmt"

func swap(el, elem *int) {
	temp := *el
	*el = *elem
	*elem = temp
}
func main() {
	x := []int{2, 44, 5, 11, 43, 6, 55, 74, 113, 8, 1}
	flag := true
	for flag {
		flag = false
		for i := 0; i < len(x)-1; i++ {
			if x[i] < x[i+1] {
				swap(&x[i], &x[i+1])
				flag = true
			}
		}
	}
	fmt.Println(len(x), x)
}
