package main

import "fmt"

func main() {
	var n int
	fmt.Print("Entr n : ")
	fmt.Scan(&n)
	sum := 0
	fmt.Println("Odd numbers are : ")
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Printf("%d ", i)
			sum += i
		}
	}
	fmt.Println()
	fmt.Printf("Sum  = %d", sum)
}
