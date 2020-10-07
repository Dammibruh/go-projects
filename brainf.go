package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func Interpret(code string) string {
	var cell [50000]int
	var index int
	var out string
	for i := 0; i < len(code)-1; i++ {
		switch code[i] {
		case '+':
			cell[index]++

		case '-':
			cell[index]--

		case '>':
			if index == len(cell)-1 {
				index = 0
			} else {
				index++
			}

		case '<':
			if index == 0 {
				index = len(cell) - 1
			} else {
				index--
			}

		case '.':
			out += string(rune(cell[index]))

		case '[':
			if cell[index] == 0 {
				l := 1
				i++
				for l > 0 {
					i++
					if code[i] == '[' {
						l++
					} else if code[i] == ']' {
						l--
					}
				}
			}

		case ']':
			if cell[index] != 0 {
				l := 1
				i--
				for l > 0 {
					i--
					if code[i] == ']' {
						l++
					} else if code[i] == '[' {
						l--
					}
				}
			}

		}
	}
	return out
}
func main() {
	file := flag.String("file", "", "compile a brianfuck file")
	flag.Parse()
	if len(*file) > 0 {
		content, err := ioutil.ReadFile(*file)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			out := Interpret(string(content))
			fmt.Println(out)
		}
	} else {
		panic("no input files")
	}
}
