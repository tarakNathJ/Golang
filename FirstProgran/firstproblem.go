package main

import "fmt"

type Person struct {
	name     string
	status   int
	email    string
	isStatus bool
}

func main() {
	fmt.Println("hello mather facker")
	var firstVariable int = 123
	fmt.Println(firstVariable)
	var secondVariable string = "tari maa ke  chut"
	fmt.Println(secondVariable)

	var thirdVariable [4]string
	thirdVariable[0] = "tari"
	thirdVariable[1] = "chut"
	thirdVariable[2] = "nunu"
	thirdVariable[3] = "gud"

	for i := 0; i < len(thirdVariable); i++ {
		// fmt.Println("this is  - ", thirdVariable[i])
	}

	// if firstVariable >= 100 {
	// 	fmt.Println("this is  data")
	// }

	var firstPerson Person
	firstPerson.name = "tarak"
	firstPerson.email = "tarakjana355@gmail.com"
	firstPerson.isStatus = true
	firstPerson.status = 123

	var secondPerson Person
	secondPerson.name = "takd"
	secondPerson.email = "tarakjana355@gmail.com"
	secondPerson.isStatus = false
	secondPerson.status = 113

	var thirdPerson Person
	thirdPerson.name = "tak"
	thirdPerson.email = "tarakjana355@gmail.com"
	thirdPerson.isStatus = false
	thirdPerson.status = 123

	var forthPerson Person
	forthPerson.name = "takad"
	forthPerson.email = "tarakjana355@gmail.com"
	forthPerson.isStatus = false
	forthPerson.status = 113

	// fmt.Println(firstPerson.isStatus)

	var mordanPerson [4]Person
	mordanPerson[0] = firstPerson
	mordanPerson[1] = secondPerson
	mordanPerson[2] = thirdPerson
	mordanPerson[3] = forthPerson

	for i := 0; i < len(mordanPerson); i++ {
		fmt.Println("this is your index : ", i, "and this is your value : ", mordanPerson[i])
	}

	fitstFunction()

	

}

func fitstFunction() {
	fmt.Println("this is my first  function")
}
