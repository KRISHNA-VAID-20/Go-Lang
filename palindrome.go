package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter n : ")
	fmt.Scan(&num)
	og := num
	var rem int
	rev := 0

	for num > 0 {
		rem = num % 10
		rev = (rev * 10) + rem
		num = num / 10
	}
	if og == rev {
		fmt.Print("Palindrome")
	} else {
		fmt.Print("Not")
	}
}
