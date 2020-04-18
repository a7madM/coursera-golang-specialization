package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	slice := make([]int, 3)
	fmt.Println("Type X to Quite")
	for {
		var input string
		fmt.Print("Input: ")
		fmt.Scan(&input)
		if input == string('X') || input == string('x') {
			break
		}

		number, err := strconv.Atoi(string(input))
		if err == nil {
			slice = append(slice, number)
		}

		sort.Ints(slice)
		fmt.Println(slice)
	}

}
