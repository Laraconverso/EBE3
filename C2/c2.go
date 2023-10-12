package main

import "fmt"

func main (){
	minFunc, err:=operation(minimum)
	averageFunc, err:= operation(average)
	maxFunc, err:= operation(maximum)


	
}

func operatioon(name string)(func(...int) int, error){

//sigue

}


//for _, -> se saltea un elemento ( en este caso la key)

//Minimo es una funcion que retorna el minimo valor de un conjunto de enteros.
//Parametros ...int
//Retorna: int  
func Min(values...int) int{
	minimum:= 10000
	for _, value :=range values{
		if valor < minimum{
			minimum = value
		}
	}
	return minimum
}

//Maximo
func Max(values...int)int{
	var maximum int
	for_, value :=range values{
		if value > maximum{
			maximum = valor
		}
	}
	return maximum
}

//Promedio
fun Promedio(values...int)int{
	var sum int
	for_, value:= range values{
		sum+=value
	}
	average:=sum/len(values)
	return average
}
