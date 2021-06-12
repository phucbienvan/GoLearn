package main

import "fmt"

var (
	text   string = "xin chao"
	number int    = 10

)

func main() {
	var numberFirst int = 10
	fmt.Println(numberFirst)

	var second float32 = 11.1
	fmt.Printf("%v %T", second, second)
	fmt.Println()

	// Ep kieu
	var numberThird int = int(second)
	fmt.Printf("%v, %T", numberThird, numberThird)
	fmt.Println()
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	

}
