package main

import "fmt"

func main(){

	var n [10]int

	for i:=0; i<10; i++{
		n[i]=i+100;
	}

	for i:=0; i<10; i++{
		fmt.Printf("element[%d]= %d \n",i,n[i])
	}

}