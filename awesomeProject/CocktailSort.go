package main

import "fmt"

// работает как бабл сорт, но она не возвращается в начала как закончила подход, а идет налево
func main() {
	list := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Println("before:", list)
	cocktailSort(list)
	fmt.Println("after: ", list)
}

func cocktailSort(list []int) {
	last := len(list) - 1
	for {
		swapped := false
		for i := 0; i < last; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
		swapped = false
		for i := last - 1; i >= 0; i-- {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
