package main

import (
	"fmt"
)

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[4])

	var y [3]float64
	y[0] = 94
	y[1] = 95
	y[2] = 96
	var total float64 = 0
	// for i := 0; i < len(y); i++ {
	// 	total += y[i]
	// }
	// With i
	// for i, value := range y {
	// 	total += value
	// }
	//Without i
	for _, value := range y {
		total += value
	}
	fmt.Println("Total : ", total)
	fmt.Println("Average : ", total/float64(len(y)))

	primeNumbers := [5]int{
		2,
		3,
		5,
		7,
		11,
	}
	fmt.Println(primeNumbers)

	//slices instead of arrays

	//var s []float64
	s := make([]float64, 5, 10) // An array with a length of 5 elements and a capacity of 10 elements.
	fmt.Println(s)

	arr := [5]int{1, 2, 3, 4, 5}
	z := arr[0:4] // start index , end index
	fmt.Println(z)

	z[1] = 7
	fmt.Println(z)
	fmt.Println(arr)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)

	fmt.Println(slice3, slice4)

	mapX := make(map[string]int)
	mapX["key"] = 10

	fmt.Println(mapX["key"])

	mapY := make(map[int]int)
	mapY[1] = 10

	fmt.Println(mapY[1])

	contries := map[string]string{
		"tr": "Türkiye",
		"uk": "Unite Kingdom",
		"us": "Unite State",
		"de": "Germany",
	}

	delete(contries, "uk")

	fmt.Println(contries["uk"])

	//country, ok := contries["uk"]

	//fmt.Println(country, ok)
	if country, ok := contries["uk"]; ok {
		fmt.Println(country)
	} else {
		fmt.Println("Country not found.")
	}

}
