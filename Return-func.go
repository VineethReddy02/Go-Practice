package main

import "fmt"

func GenDispfn(acc, vel, dis float64) func(t float64) float64 {
	return func(t float64) float64 {
		res := (0.5 * acc * t * t) + (vel * t) + dis
		return res
	}

}

func main() {
	var acc, vel, dis, time float64
	fmt.Println("Enter the acceleration")
	fmt.Scanln(&acc)
	fmt.Println("Enter the velocity")
	fmt.Scanln(&vel)
	fmt.Println("Enter the displacement")
	fmt.Scanln(&dis)
	f := GenDispfn(acc, vel, dis)
	fmt.Println("Enter the time ")
	fmt.Scanln(&time)
	fmt.Println(f(time))

}
