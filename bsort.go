package main

import "fmt"

func bubblesort(arr []int) []int {
	for index, item := range arr {
		for indx, sitem := range arr {
			if item <= sitem {
				arr[index], arr[indx] = arr[indx], arr[index]
			}
		}

	}
	return arr
}
func main() {
	bruh := []int{6, 5, 7, 8, 4, 7}
	fmt.Println("old array:", bruh)
	newar := bubblesort(bruh)
	fmt.Println("sorted arr:", newar, "\n")
}
