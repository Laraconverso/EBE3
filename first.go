package main

import (
	"fmt"
)


func main(){

	// imprimit hola mundo 
	fmt.Println("Hola Mundo! :)")

	// declarar variables
	// var nombre string
	// var horas int
	// //asignar valores
	// horas = 20
	// nombre = "Lara"

	// producto1, precio2 := "Jean", 10.5

	// //se puede agrupar en bloques 
	// var (
	// 	producto = "Course" 
	// 	cantidad = 25
	// 	precio = 100
	// 	enStock = true
	// )

	//constantes
	const status = "ok"

	//tipos de datos
	var a [2]string
	a[0] = "Hello"
	a[1] = "world!"
	fmt.Println(a[0],a[1])
	fmt.Println(a)
}