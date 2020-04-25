package main

import "fmt"

func main(){
	var map1 map[string]int

  var map2 =map[string]float64 {
	  "pi" : 3.14,
	  "e" : 2.7,
	  "x" :5,
  }

  fmt.Println(map2);

	fmt.Println(map1)

	if map1==nil {	
		fmt.Println("empty")
	}
}