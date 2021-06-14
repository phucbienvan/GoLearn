package main

import (
	"fmt"
	"math"
	// "time"
	"sort"
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
	// fmt.Println(pow(5, 4, 1000))
	// t := time.Now()
	// switch {
	// case t.Hour() +10 < 12:
	// 	fmt.Println("Good morning!")
	// case t.Hour() < 17:
	// 	fmt.Println("Good afternoon.")
	// default:
	// 	fmt.Println("Good evening.")
	// }
	// sum := 1
	// for sum < 1000{
	// 	sum += sum
	// }
	// fmt.Println(sum)

	// numPointer := 9
	// convert(numPointer)
	// fmt.Println(numPointer)
	// convertAddress(&numPointer)
	// fmt.Println(numPointer)

	//Slices
	// numSlice := []int{1, 2, 10}
	// fmt.Println(numSlice)
	// numSlice[2] = 11
	// fmt.Println(numSlice)
	// fmt.Println(numSlice[0])

	var ele rune
    var size int
    var sli = make([]int,0,1)
    size = cap(sli)
    for i:=0; i<=size; i++{
        if i>=len(sli){
            size=size+1
        }
        ele = 0
        fmt.Println("Enter a number to add: ")
        fmt.Scan(&ele)
        if ele==0 {
            fmt.Println("Stopping!")
            break
        }
        sli = append(sli, int(ele))
    }
    fmt.Println(sli)
	bubbleSoft(sli)
	fmt.Println(sli)
}

func bubbleSoft(slide sort.IntSlice){
	for j := 0; j < len(slide)-1; j++ {
		for i := len(slide)-2; i >= j; i-- {
			if slide[i] > slide[i+1] {
				tg := slide[i]
				slide[i] = slide[i+1]
				slide[i+1] = tg
			}
		}
	}
}

func convert(a int){
	a = 100
}
func convertAddress(a *int){
	*a = 200
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
