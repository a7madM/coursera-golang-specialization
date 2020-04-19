package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var filename string
	fmt.Print("Filename: ")
	fmt.Scan(&filename)

	filecontent, err := ioutil.ReadFile(filename)
	check(err)

	namesSlice := strings.Split(string(filecontent), "\n")
	names := make([]Name, 0, len(namesSlice))

	for _, v := range namesSlice {
		n := strings.Split(v, " ")

		name := Name{n[0], n[1]}

		names = append(names, name)
	}

	for _, name := range names {
		fmt.Printf("Firstname: %s, Lastname: %s \n", name.fname, name.lname)
	}
}
