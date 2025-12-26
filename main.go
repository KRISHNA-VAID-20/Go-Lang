package main
import "fmt"
func main(){
	n:=1
	fmt.Println(n)
}

package main

import "fmt"

func gcd(a,b int)int{
	for b != 0{
		a,b=b,a%b

	}
	return a
}

func lcm(a,b int)int{
	return (a*b)/(gcd(a,b))
}


func main(){
	// fmt.Println(("krishna"))
	var a,b,c int
	fmt.Println("Enter a,b,c : ")
	fmt.Scan(&a,&b,&c)
	g:=gcd(gcd(a,b),c)
	l:=lcm(lcm(a,b),c)
	sum:=float64(a)+float64(b)+float64(c)
	product:=float64(a*b*c)
	fmt.Println("Gcd : ",g)
	fmt.Println("Lcm : ",l)
	fmt.Println("Sum : ",sum)
	fmt.Println("Product : ",product)
}

package main
import "fmt"

func main(){
	var n int
	fmt.Println("ENter n : ")
	fmt.Scan(&n)

	num:=1
	
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Println(num," ")
			num++
		}
		fmt.Println()

	}
}

package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter Limit : ")
	fmt.Scan(&n)

	a,b:=0,1

	fmt.Printf("Fibonacci of %d numbers : \n",n)
	for a<=n{
		fmt.Println(a," ")
		a,b=b,a+b
	}
}


package main
import "fmt"

func main(){
	var rev,temp,n int

	fmt.Println("Enter n : ")
	fmt.Scan(&n)

	temp=n

	for  n>0 {
		rev=rev*10+n%10
		n/=10

	}

	if(rev==temp){

		fmt.Println("Palindrome")
	}else{
		fmt.Println("Not ")
	}
}


package main
import "fmt"
func main() {
	var n int
	fmt.Println("Enter number of rows: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Println(j, " ")
	}
		fmt.Println()
	}
	}


package main
import "fmt"

func main(){
	var n,sum int

	fmt.Println("Enter n : ")
	fmt.Scan(&n)

	for i:=1;i<=n;i++{
		if i%2 !=0{
			sum+=i
		}
	}

	fmt.Println("Sum : ",sum)
}


package main
import "fmt"

func createstr()string{
	return "Krishna"
}

func strlength(s string)int{
	return len(s)
}




func main(){
	str:=createstr()
	fmt.Println("String : ",str)
	fmt.Println("String Length : ",strlength(str))


}


package main

import "fmt"

func main(){
	a:=[3]int{1,2,3}
	b:=[3]int{1,2,3}

	equal:=true

	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			equal=false
			break
		}
	}

	if equal{
		fmt.Println("Equal")
	}else{
		fmt.Println(("Not Equal"))
	}
}

package main
import "fmt"

func main(){
	a:=[3]int{1,2,3}
	b:=[3]int{4,5,6}

	var c[6]int

	k:=0
	for i:=0;i<len(a);i++{
		c[k]=a[i]
		k++
	}
	for i:=0;i<len(b);i++{
		c[k]=b[i]
		k++
	}

	fmt.Println("Merged Arrays : ",c)

}

package main
import "fmt"

func main(){
	var a int =10
	var p *int=&a

	fmt.Println("value of a : ",a)
	fmt.Println("Address of a : ",p)
	fmt.Println("value of p : ",*p)

	*p=20
	fmt.Println("Modified value of a : ",a)
}


package main
import "fmt"


type student struct{
	id int
	name string
	mark float64
}

func main(){

	var s student
	s.id=20
	s.name="Krishna"
	s.mark=90.3

	fmt.Println("Student details : ")
	fmt.Println("Id : ",s.id)
	fmt.Println("Name : ",s.name)
	fmt.Println("Marks : ",s.mark)

}
