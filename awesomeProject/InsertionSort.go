package main

import (
	"fmt"
)

func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	InsertionSort(ar)
	fmt.Println(ar)
}

func InsertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
