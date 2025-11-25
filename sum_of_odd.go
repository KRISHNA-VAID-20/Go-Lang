// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Print("Entr n : ")
// 	fmt.Scan(&n)
// 	sum := 0
// 	fmt.Println("Odd numbers are : ")
// 	for i := 1; i <= n; i++ {
// 		if i%2 != 0 {
// 			fmt.Printf("%d ", i)
// 			sum += i
// 		}
// 	}
// 	fmt.Println()
// 	fmt.Printf("Sum  = %d", sum)
// }

// package  main

// import (
// 	"fmt"
// 	"reflect"

// )

// func compare(arr1,arr2 interface{})bool{
// 	if reflect.TypeOf(arr1).Kind()!=reflect.Array|| reflect.TypeOf(arr2).Kind()!=reflect.Array{
// 		fmt.Print("Not arrays")
// 		return false
// 	}

// 	slice1:=reflect.ValueOf(arr1)
// 	slice2:=reflect.ValueOf(arr2)

// 	if slice1.Len()!=slice2.Len(){
// 		fmt.Print("Not equal ")
// 		return false
// 	}

// 	for i:=0;i<slice1.Len();i++{
// 		if !reflect.DeepEqual(slice1.Index(i).Interface(),slice2.Index(i).Interface()){
// 			return false
// 		}
// 	}
// 	return true


// }

// func main(){
// 	arr1:=[5]int{1,2,3,4,5}
// 	arr2:=[5]int{1,2,3,4,5}

// 	result:=compare(arr1,arr2)
// 	if result{
// 		fmt.Print("Equal")
// 	}else{
// 		fmt.Print("Not equal")
// 	}
// }

func mergearray(arr1 []int,arr2[]int)[]int{
	mertged:=make([]int,len(arr1)+len(arr2))

}