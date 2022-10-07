package main

import "fmt"

// Расчёска лучше пузырьковой сортировки, потому что в ней намного меньше операций перестановки
func main() {
	list := []int{3, 4, 1, 2, 5, 7, -1, 0}
	fmt.Println("before:", list)
	combSort(list)
	fmt.Println("after: ", list)
}

func combSort(list []int) {
	if len(list) < 2 {
		return
	}
	for gap := len(list); ; {
		if gap > 1 {
			gap = gap * 4 / 5
		}
		swapped := false
		for i := 0; ; {
			if list[i] > list[i+gap] {
				list[i], list[i+gap] = list[i+gap], list[i]
				swapped = true
			}
			i++
			if i+gap >= len(list) {
				break
			}
		}
		if gap == 1 && !swapped {
			break
		}
	}
}
