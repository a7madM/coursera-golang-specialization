package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Animal struct {
	Name string
	Food string
	LocomotionMethod string
	SpokenSound string
}


func(animal *Animal) Eat() {
	fmt.Println(animal.Food)
}

func(animal *Animal) Speak() {
	fmt.Println(animal.SpokenSound)
}

func(animal *Animal) Move() {
	fmt.Println(animal.LocomotionMethod)
}

func createAnimals() []Animal {
	animals := make([]Animal, 0, 4)

	cow := Animal{"cow", "grass", "walk", "moo"}
	bird := Animal{"bird", "worms", "fly", "peep"}
	snake := Animal{"snake", "mice",	"slither",	"hsss"}

	animals = append(animals, cow)
	animals = append(animals, bird)
	animals = append(animals, snake)

	return animals
}

func main() {
	animals := createAnimals()
	for {
		fmt.Print("Your Request > ")
		request := readRequest()
		animal := getAnimal(animals, request)
		animal.performRequest(request)
	}
}

func(animal *Animal) performRequest(request []string) {
	switch request[1] {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}	
}

func getAnimal(animals []Animal, request []string) Animal {
	var animal Animal
	switch request[0] {
	case "cow":
		animal = animals[0]
	case "bird":
		animal = animals[1]
	case "snake":
		animal = animals[2]
	}

	return animal
}
func readRequest() []string {
	var line string 
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
	    line = scanner.Text()
	}

	return strings.Split(line, " ")
}