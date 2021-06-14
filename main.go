package main

import (
	"fmt"
	"math"
	"time"
)


func main() {
	// var numberFirst int = 10
	// fmt.Println(numberFirst)

	// var second float32 = 11.1
	// fmt.Printf("%v %T", second, second)
	// fmt.Println()

	// // Ep kieu
	// var numberThird int = int(second)
	// fmt.Printf("%v, %T", numberThird, numberThird)
	// fmt.Println()
	// sum := 0
	// for i := 0; i <= 10; i++ {
	// 	sum += i
	// 	fmt.Printf("phan tu thu: %v\n", i)
	// 	fmt.Printf("Sum = %v \n", sum)
	// 	fmt.Println()
	// }
	fmt.Println(pow(5, 4, 1000))
	t := time.Now()
	switch {
	case t.Hour() +10 < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	sum := 1
	for sum < 1000{
		sum += sum
	}
	fmt.Println(sum)
	
	
}

func convert(num1 int, num2 int){
	var tmp int = num1
	num1 = num2
	num2 = tmp
	fmt.Println("new", num1, num2)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
