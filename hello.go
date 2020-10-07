package main

import "fmt"

func main() {
	br := []int{7, 9, 8, 9, 9}
	br = append(br, 5)
	for i := 0; i < len(br); i++ {
		fmt.Printf("%d\n", br[i])
	}
}
