package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

var cow = Animal{"grass", "walk", "moo"}
var bird = Animal{"worms", "fly", "peep"}
var snake = Animal{"mice", "slither", "hiss"}

func (animal *Animal) Eat() {
	fmt.Printf("%v\n", animal.food)
}

func (animal *Animal) Move() {
	fmt.Printf("%v\n", animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Printf("%v\n", animal.noise)
}

func findAnimal(animal string) (Animal, error) {
	switch animal {
	case "cow":
		return cow, nil
	case "bird":
		return bird, nil
	case "snake":
		return snake, nil
	default:
		return Animal{}, errors.New("Animal not found\n")
	}
}

func scan(msg string) (string, string, error) {
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := strings.Fields(strings.ToLower(scanner.Text()))
		if len(input) == 2 {
			return input[0], input[1], nil
		}
	}

	return "", "", errors.New("Please supply one animal and one method, seperated by a space\n")
}

func main() {
	userAnimal, userMethod, error := scan("> ")

	if error != nil {
		fmt.Print(error)
		main()
		return
	}

	animal, error := findAnimal(userAnimal)

	if error != nil {
		fmt.Print(error)
		main()
		return
	}

	switch userMethod {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Print("Method not found\n")
	}

	main()
}
