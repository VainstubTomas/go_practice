package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	sec := time.Now().Unix()
	// fmt.Println((sec))
	rand.Seed(sec)
	i := rand.Int31n(10)

	fmt.Println(i)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	}

	fmt.Println("ok")

	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)

	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
		fmt.Println("")
	}

	sum := 0
	for i := 1; i < 100; i++ {
		sum += i
	}
	fmt.Println((sum))

	//while con for - hasta que los valores de "j" sean iguales a 10
	j := 0
	for j != 10 {
		j++
		fmt.Println((j))
	}

}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}
