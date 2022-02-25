package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	//only for

	for i := 1; i < 10; i++ {
		println(i)
	}

	id := 1
	for {
		if id == 10 {
			break
		}
		fmt.Println(id)
		id += 1
	}

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else if i%3 == 0 {
			fmt.Println(i, "divisible by 3")
		} else if i%4 == 0 {
			fmt.Println(i, "divisible by 4")
		} else {
			fmt.Println(i, "out")
		}
	}

	if i := rand.Intn(100); i < 90 {
		fmt.Println(i, "hey")
	}

	os := runtime.GOOS

	// auto break first equal
	switch os {
	case "darvin":
		fmt.Println("MacOs")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("%s\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}

}
