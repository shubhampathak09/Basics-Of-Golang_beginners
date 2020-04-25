package main

import "fmt"

type Employee struct{
	name string
	designation string
	address string 
	age int
	
}


func main(){

var e1 Employee
//var e2 Employee

e1.name="name"
e1.designation="designation"
e1.address="address"
e1.age=40

fmt.Printf("e1: ",e1,"\n")

//e2=Employee("name1","designation","address2",30)

//fmt.Printf("e2: ",e2,"\n")

	
}