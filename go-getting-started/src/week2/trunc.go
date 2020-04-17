package main

import "fmt"

func main() {
	var floatNumber float64
	fmt.Printf("Float Number: ")

	fmt.Scan(&floatNumber)

	fmt.Println(int(floatNumber))
}
