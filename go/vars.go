package main

import "fmt"

/*
tipos de variables 
impresion de variables junto a texto
diferencia entre constante y variable
*/

func main(){
	var x int; //definicion de entero
	var y float64; //definicion de coma flotante
	var z="Peter"; //define strings. se asigna el valor cuando se crea
	var j byte; //define bytes
	const u int=80; //CONSTANTE su valor se mantiene invariable en el transcurso del programa, se asigna el valor en el momento de su creación

	//var g bool=false --> tipo boolean puede ser falso o verdadero, sobre todo usada en ifs
	
	x=200
	y=32.4
	j=128
	
	
	fmt.Println("x is",x)
	fmt.Println("y is",y)
	fmt.Println("z is",z)
	fmt.Println("j is",j)
	fmt.Println("u is a constant and it is",u)
	
}
