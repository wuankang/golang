package main 

import(
	"fmt"
)
/*func main(){
	fmt.Println("hello world")
}*/

func main(){
	var num1 int = 1199
	var num2 int = 1200
	var d1 float64
	d1 = float64(num1)/float64(num2)
	fmt.Printf("%.2f\n",d1)
}
