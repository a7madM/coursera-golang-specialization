package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := make([]int, 0, 10)
	fmt.Println("Type x/X to Break")

	for i := 0; i < 10; i++ {

		var input string
		fmt.Print("Input: ")
		fmt.Scan(&input)
		if input == string('X') || input == string('x') {
			break
		}

		number, err := strconv.Atoi(string(input))
		if err == nil {
			numbers = append(numbers, number)
		} else {
			panic(err)
		}
	}
	BubbleSort(numbers)

	fmt.Println(numbers)
}

// BubbleSort Function that sorts array of integer numbers
func BubbleSort(numbers []int) {
	length := len(numbers) - 1
	for i := length; i > 0; i-- {
		for x := 0; x < length; x++ {
			if numbers[x] > numbers[x+1] {
				Swap(x, numbers)
			}
		}
	}
}

// Swap Function to Swap number in index i with number in index i+1
func Swap(i int, sli []int) {
	temp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = temp
}
