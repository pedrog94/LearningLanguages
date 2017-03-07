package main

import "fmt"

/*
	sentencias if
	aplicación de boolean type
	comparando variables
*/


func main(){

	var rain bool=false
	var x int=6
	
	if(!rain){
		fmt.Println("not raining")
	}else{
		fmt.Println("take umbrella!!")
	}

	if(x > 7){
		fmt.Println("x is",x)
	}
	fmt.Println("hi, do something")
}
