package main

import "fmt"

func main() {

	// chck even and odd

	// var number int16 = 12

	// count of n number
	var count = 0
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("this is even number", i)
		} else {
			fmt.Println("this is odd number", i)
		}
		count += i

	}
	fmt.Println("count is ", count)

	var simpleArray = [5]int{12, 34, 12, 54, 24}
	secondarray := [...]string{"som", "bom", "lom", "chum", "chut"}

	var graternumber int = 0
	for i := 0; i < len(simpleArray); i++ {

		if graternumber < simpleArray[i] {
			graternumber = simpleArray[i]
		}
		fmt.Println("this is your second array ; ", secondarray[i])
	}

	fmt.Println("this is your gratest number : ", graternumber)

	fmt.Println("./////////////////////./////////./////////////")

	primemunber := 10
	counts := 0

	for i := 2; i <= primemunber/2; i++ {

		if primemunber%i == 0 {
			fmt.Println("this is not prime number")
			counts++
		}

	}
	fmt.Println("your count is :", counts)

}
