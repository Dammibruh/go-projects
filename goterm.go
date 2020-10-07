package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	-l: count line;

*/
func readF(name string, cl bool) (string, error) {
	content, err := ioutil.ReadFile(name)
	if cl {
		_cn := strings.Split(string(content), "\n")
		var out string
		for i, item := range _cn {
			out += fmt.Sprintf("%d %s\n", (i + 1), item)
		}
		return out, err
	} else {
		return string(content), err
	}
}

func main() {
	count_line := flag.Bool("l", false, "test")
	flag.Parse()
	out, err := readF(flag.Args()[0], *count_line)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(out)
	}
}
