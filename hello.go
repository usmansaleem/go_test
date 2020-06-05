package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// name returned values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// variables
// var c, python, java bool

// var with initialization
var j, k int = 1, 2

// var with initialization without type
var c, python, java = true, false, "yes"

const NAME = "Uzi"

var (
	firstName = "Usman"
	lastName  = "Saleem"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(3, 2))
	fmt.Println(add2(3, 2))
	fmt.Println(swap("x", "y"))
	fmt.Println(split(17))

	// variables
	var i int
	fmt.Println(i, j, k, c, python, java)

	// short variable declaration
	m := 10
	fmt.Println(m)

	// type conversion
	amount := 10.15
	amountWhole := uint(amount)
	fmt.Println(amount, amountWhole)

	// type
	fmt.Printf("amount type: %T\n amountWhole Type: %T\n", amount, amountWhole)

	fmt.Println("Name: ", NAME)

	fmt.Println(firstName, lastName)
}
