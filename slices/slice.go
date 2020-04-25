package main

import "fmt"



func printSlice(x []int) {
fmt.Printf("len=%d capacity=%d slice=%v\n",len(x),cap(x),x)
}


func main(){
	
	
	// declare empty slace make

	//var numbers=make([]int,5,10)
	// pointer
	//length
	//capacity


	var numbers=[]int{5,10,15,20,25,30}

	var slice1[] int=numbers[1:4]

	printSlice(slice1)

}