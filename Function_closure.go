package main

import (
	"fmt"
	"strconv"
)

// Function closure
func factorial() func (int) int {
	a := 1;
	return func (i int) int{
		a *= (i); //4 
		return a ;
	}
}

type Physical struct {
	height, weight float64
}

func main(){
	fmt.Print("Enter the number :");
	x := 0;
	fmt.Scan(&x);
	fact := factorial();
	for i := x; i>1; i-- {
		f := fact(i);
		if (i == 2){
			fmt.Print("Factorial of the number is : ",f,"\n");
		}
	}

	//int to float
	fmt.Print("Enter an integer :");
	var a int32;
	fmt.Scan(&a);
	b := float32(a);
	fmt.Printf("Type casting %d as float = %f\n",a,b);
	//float to int
	fmt.Print("Enter an Float :");
	var a1 float32;
	fmt.Scan(&a1);
	b1 := int32(a1);
	fmt.Printf("Type casting %f as float = %d\n",a1,b1);
	//int to string
	fmt.Print("Enter an integer :");
	var a2 int;
	fmt.Scan(&a2);
	b2 := strconv.Itoa(a2);
	fmt.Printf("Type casting %d as string = %q\n",a2,b2);
	//string to int
	a3 := "200";
	b3,_ := strconv.Atoi(a3);
	fmt.Printf("%q is string. Type casted = %T\n",b3,b3);

	//maps
	fmt.Print("Enter name :");
	s := "";
	fmt.Scan(&s);
	m := make(map[string]Physical)
	h,w :=0.0,0.0;
	fmt.Print("Enter height and weight resp. :");
	fmt.Scan(&h, &w);
	m[s] = Physical{
		h, w,
	}
	fmt.Println("Height and weight of ",s," is ",m[s].height,m[s].weight)

}