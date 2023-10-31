package main

import (
	"fmt"
	"os"
	"time"
)

// ej de clase
func c7(){
	
	var input string

	channelDevolucion := make(chan string)
	channelPrestamo := make(chan string)

	//escribo en el channel una go rutina que escribe en un canal 
	go Decoluvion(channelDevolucion) 

	go func(channelPrestamo chan string){ //con el go es designamos que es una go rutina
		defer close(channelPrestamo)
		for{
			time.Sleep(time.Second /2) //cada medio segundo se escribe en el canal 2 prestamos procesados
			channelDevolucion <- "Prestamo procesado"
		}
	}(channelPrestamo)//el parametro o sea 

	//creo 2 gorutinas mas para imprimir en la consola. 
	go func(channelDevolucion chan string){
		for devolucion:= range channelDevolucion{
			fmt.Println(devolucion)

		}
	}(channelDevolucion)


	go func(channelPrestamo chan string){
		for prestamo := range channelPrestamo{
			fmt.Println(prestamo)

		}
	}(channelPrestamo)

	//si ingreso algo por teclado tiene que salir de programa.
	fmt.Scan(&input)
	if input != ""{
		fmt.Println("Saliendo del programa...")
		os.Exit(0)
	}


}

//devolucion...
func Decoluvion(channelDevolucion chan string){ //con el go es designamos que es una go rutina
	defer close(channelDevolucion)
	for{
		time.Sleep(time.Second *1) //por defecto es 1 pero aclaramos x las dudas
		channelDevolucion <- "Decoluvion procesada"
	}
}