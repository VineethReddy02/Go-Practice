package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type GenericAnimal struct {
	food       string
	locomotion string
	noise      string
}

var cow = GenericAnimal{"grass", "walk", "moo"}
var bird = GenericAnimal{"worms", "fly", "peep"}
var snake = GenericAnimal{"mice", "slither", "hiss"}
var userAnimalMap = map[string]GenericAnimal{}

func scan(msg string) (string, string, string, error) {
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input := strings.Fields(strings.ToLower(scanner.Text()))
		if len(input) == 3 {
			return input[0], input[1], input[2], nil
		}
	}

	return "", "", "", errors.New("Please supply three strings, seperated by a space.\n")
}

func (animal GenericAnimal) Eat() {
	fmt.Printf("%v\n", animal.food)
}

func (animal GenericAnimal) Move() {
	fmt.Printf("%v\n", animal.locomotion)
}

func (animal GenericAnimal) Speak() {
	fmt.Printf("%v\n", animal.noise)
}

func findAnimal(animal string) (GenericAnimal, error) {
	switch animal {
	case "cow":
		return cow, nil
	case "bird":
		return bird, nil
	case "snake":
		return snake, nil
	default:
		return GenericAnimal{}, errors.New("Animal not found\n")
	}
}

func printAndReturn(msg interface{}) {
	fmt.Println(msg)
	main()
	return
}

func main() {
	// Scan for user input
	command, firstInput, secondInput, error := scan("> ")

	if error != nil {
		printAndReturn(error)
	}

	switch command {
	case "newanimal":
		// Find base animal from user input
		animal, error := findAnimal(secondInput)

		if error != nil {
			printAndReturn(error)
		}

		// Find or create new animal
		userAnimalMap[firstInput] = animal
	case "query":
		// Find queried animal
		userAnimal := userAnimalMap[firstInput]

		// Check for non-existent animal
		if (GenericAnimal{}) == userAnimal {
			printAndReturn("User animal not found\n")
		}

		switch secondInput {
		case "eat":
			userAnimalMap[firstInput].Eat()
		case "move":
			userAnimalMap[firstInput].Move()
		case "speak":
			userAnimalMap[firstInput].Speak()
		default:
			fmt.Print("Method not found\n")
		}
	default:
		fmt.Print("Command not recognized\n")
	}

	main()
}
