package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var text string
	fmt.Print("Input: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')

	text = strings.TrimSpace(text)
	var result string

	if meetsCriteria(text) == true {
		result = "Found!"
	} else {
		result = "Not Found!"
	}

	fmt.Println(result)

}

func meetsCriteria(text string) bool {
	return firstCharIsI(text) && lastCharIsN(text) && includesCharacterA(text)
}

func firstCharIsI(text string) bool {
	return (string(text[0]) == string('i') || string(text[0]) == string('I'))
}

func lastCharIsN(text string) bool {
	lastIndex := len(text) - 1
	return (string(text[lastIndex]) == string('n') || string(text[lastIndex]) == string('N'))
}

func includesCharacterA(text string) bool {
	length := len(text)
	for i := 1; i < length-1; i++ {
		if string(text[i]) == string('a') || string(text[i]) == string('A') {
			return true
		}
	}
	return false
}
