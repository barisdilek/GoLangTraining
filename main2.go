package main

import "fmt"

var greeting, name string = "Hello", "Dilek"

func main() {
	//var age int
	//age = 41
	// only in func so work
	age := 41

	fmt.Printf("%s, %s. You are %d years old.\n", greeting, name, age)

	const pi = 3.14

	fmt.Println(pi)

	var luckyNumber int
	fmt.Println(luckyNumber)

	luckyNumber = 1_234
	fmt.Println(luckyNumber)

	//favoriteBook := "The Brothers\n\"Karamazov\""
	favoriteBook := `The Brothers
	"Karamazov"`
	fmt.Println(favoriteBook)

	var stock uint = 3
	var wealt int = -200
	fmt.Println(stock, wealt)

	var progress float32 //float64
	fmt.Println(progress)

	var x int = 42
	var y float32 = float32(x)
	var z uint = uint(y)

	fmt.Println(x, y, z)
}
