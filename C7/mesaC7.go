package main

import (
	"fmt"
	"os"
	"time"
)


var numbers = []int{0,1,2,3,4,5,6,7,8,9,10}

func main(){
	var input string
	

	parChannel:= make(chan int)
	imparChannel:= make(chan int)


	go par(parChannel)
	go impar(imparChannel)
	go parPrint(parChannel)
	go imparPrint(imparChannel)

	//si ingreso algo por teclado tiene que salir de programa.
	fmt.Scan(&input)
	if input != ""{
		fmt.Println("Saliendo del programa...")
		os.Exit(0)
	}


}

func par(parChannel chan int){
	for _,num := range numbers{
		if num %2 ==0{
			time.Sleep(time.Second *1)
			parChannel <- num
		}
	}
}

func impar(imparChannel chan int){
	for _,num := range numbers{
		if num %2 != 0{
			time.Sleep(time.Second *1)
			imparChannel <- num
		}
	}
}



func parPrint(parChannel chan int){
	for{
		num:= <- parChannel
		fmt.Printf("Numero par: %d\n", num)
	}
}

func imparPrint(imparChannel chan int){
	for{
		num:= <- imparChannel
		fmt.Printf("Numero impar: %d\n", num)
	}
}


