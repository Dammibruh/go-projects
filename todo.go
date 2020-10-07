package main

import (
	"fmt"
	"strconv"
)

func main() {
	todos := []string{}
	fmt.Print("lmao todo app w/golang\n[1] - add todo\n[2] - list todos\n")
	for true {
		var choice string
		fmt.Scanln(&choice)
		_ch, _ := strconv.Atoi(choice)
		if _ch == 1 {
			var todo string
			fmt.Scanln(&todo)
			todos = append(todos, todo)
			fmt.Print("done\n")
		} else if _ch == 2 {
			for index, item := range todos {
				fmt.Println(index, "-", item, "\n")
			}
		}
		else if _ch == 3 {
			
		}
	}

}
