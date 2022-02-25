package main

import "fmt"

func main() {
	greeting("Dilek")
	sum(4, 9)

	fmt.Println(add(3, 5, 5, 7, 8, 8, 4, 5))

	xs := []int{3, 5, 5, 7, 8, 8, 4, 5}
	fmt.Println(add(xs...))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	result := factorial(4)
	fmt.Println(result)

	defer goodbye() //run after main function. finaly
	hello()
}

func hello() {
	fmt.Println("Hello")
}

func goodbye() {
	fmt.Println("Good bye")
}

func greeting(name string) {
	fmt.Printf("Hello %s\n", name)
}

func sum(a int, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

// func add(a, b int) int {
// 	return a + b
// }

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
