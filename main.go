package main

import "fmt"

var (text string = "xin chao"
	number int = 10
)
func main()  {
	m := 2
	fmt.Println( m);
	k := 10
	fmt.Println(k)
	a := k * m
	fmt.Println("k * m = ", a)
	fmt.Printf("%v, %T", a, a)
	fmt.Println()
	fmt.Println(text)
	sum:= 0;
	for i := 0; i < 10; i ++{
		sum += i
	}

	fmt.Println(sum)
	

}
