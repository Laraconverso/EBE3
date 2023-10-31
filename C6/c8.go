package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const(
	filename = "data.csv"
)

func main(){


	defer func(){
		if err:=recover(); err != nil{
			log.SetPrefix("main:")
			log.Println(err)
			os.Exit(1)
			//log.Fatal(err)
		}
	}()

	//call function 

	readfile(filename)
}




func readfile(name string){

	file, err:= os.ReadFile(name)

	if(err)!= nil{
		panic(err)
	}

	data:= strings.Split(string(file), ";")
	
	var total float64
	for i:=0; i <len(data)-1;i++{
		line:= strings.Split(string(data[i]), ",")
		fmt.Print(line)
	}

	
		//omitir la cabecera
		if i !=0 {
			precio, err := strconv.ParseFloat(line[1],64)
			if err!=nil{
				log.Panicln("No se pudo convertir el precio")
			}

			cantidad,err := strconv.ParseFloat(line[2],64)
			if err!= nil{
				log.Panicln("No se pudeo convertir la cantidad")
			}

			totalProducto := precio * cantidad

			total += totalProducto
		}

		for i := 0;i<len(line);i++{
			fmt.Printf("%s\t\t", line[i])
			if(i==len(line)-1){
				fmt.Print("\n")
			}
		}
		fmt.Printf("\t\t%s\t\t%.2f\n", "Total:", total)
	}

}