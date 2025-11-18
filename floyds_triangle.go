package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter n : ")
	fmt.Scan(&n)

	num := 1
	sum := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			sum += num
			num += 1
		}
		fmt.Println()
	}

	fmt.Printf("Sum ofall numbers : %d", sum)
	fmt.Println()

	if sum%2 == 0 {
		fmt.Print("Even")
	} else {
		fmt.Print("Odd")
	}

}
