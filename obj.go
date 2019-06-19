package main

import "fmt"

type Animal interface {
	Eat()
	Locomotion()
	Sound()
}

type animal struct {
	food       string
	locomotion string
	sound      string
}

func (an *animal) eat1(animal_name string, action string) {
	if animal_name == "cow" {
		an.food = "grass"
	} else if animal_name == "bird" {
		an.food = "worms"
	} else {
		an.food = "mice"
	}

}
func (an *animal) locomotion1(animal_name string, action string) {
	if animal_name == "cow" {
		an.locomotion = "walk"
	} else if animal_name == "bird" {
		an.locomotion = "fly"
	} else {
		an.locomotion = "slither"
	}
}

func (an *animal) sound1(animal_name string, action string) {
	if animal_name == "cow" {
		an.sound = "moo"
	} else if animal_name == "bird" {
		an.sound = "peep"
	} else {
		an.sound = "hsss"
	}
}

type mike struct {
	ba string
	ma string
}

func main() {
	fmt.Println("***************************************")
	fmt.Println("Below are the Format in interacting..!! ")
	fmt.Println("newanimal samplename cow/bird/snake")
	fmt.Println("query samplename eat/move/speak")
	fmt.Println("****************************************")
	fmt.Println()
	a1 := make([]mike, 1)
	// m = make(map[string]string)
	var s1 = ""
	var s2 = ""
	var s = ""
	var p animal
loop:

	fmt.Print(">")
	fmt.Scan(&s)
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	if s == "query" {
		if s1 != "cow" && s1 != "bird" && s1 != "snake" {
			for _, num := range a1 {
				if num.ba == s1 {
					if num.ma == "cow" {
						s1 = "cow"
						break
					} else if num.ma == "bird" {
						s1 = "bird"
						break
					} else if num.ma == "snake" {
						s1 = "snake"
						break
					}
				}
			}
		}
		if s2 == "eat" {
			p.eat1(s1, s2)
			fmt.Println(p.food)
		} else if s2 == "move" {
			p.locomotion1(s1, s2)
			fmt.Println(p.locomotion)
		} else if s2 == "speak" {
			p.sound1(s1, s2)
			fmt.Println(p.sound)
		}
		//s1 = ""
	} else if s == "newanimal" {
		var k mike
		k.ba = s1
		k.ma = s2

		a1 = append(a1, k)
		fmt.Println("Created it!")

	}
	goto loop
}
