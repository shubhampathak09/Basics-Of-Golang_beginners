package main

import "fmt"

// func main()
// {

// 	var value_1=10;

// 	var value_2=20;

// 	fmt.Println("Before swap..val1: %d,val2: %d",value_1,value_2)

// }


func swap_val(value_1*int ,value_2*int){

	var temp=*value_1;
	*value_1=*value_2;
	*value_2=temp;
}


func main(){ 


var value_1 = 10

var value_2 = 20


fmt.Println("Before Swap:")

fmt.Println("value_1: %d, value_2: %d", value_1, value_2)

fmt.Println("After Swap")

swap_val(&value_1,&value_2);
 

fmt.Printf("value_1: %d, value_2: %d", value_1, value_2)



}