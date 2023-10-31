//uno de los casos que puede ocurrir con las subrutinas son las condicion de carrera
//cuando 2 quieren acceder al mismo, y como no se puede acceder explota la go rutina

//ejemplo:

package main

import (
	"fmt"
	"sync"
	"time"
)

var(
	objects= make(map[string]any) //el any es de tipo interface
	mutex sync.Mutex // dentro de una go rutina lo bloquea (lo reserva trabaja y despues lo libera)
)

func extraMutex(){
	go addObject("Key 1", "Valor 1")
	go addObject("Key 2", "Valor 2")
	go addObject("Key 3", "Valor 3")
	go addObject("Key 4", "Valor 4")
	go addObject("Key 5", "Valor 5")
	go addObject("Key 6", "Valor 6")
	go addObject("Key 7", "Valor 7")
	go addObject("Key 8", "Valor 8")
	go addObject("Key 9", "Valor 9")
	go addObject("Key 10", "Valor 10")
	go addObject("Key 11", "Valor 11")
	go addObject("Key 12", "Valor 12")
	go addObject("Key 13", "Valor 1")
	go addObject("Key 14", "Valor 1")
	go addObject("Key 15", "Valor 1")
	go addObject("Key 16", "Valor 1")
	go addObject("Key 17", "Valor 1")
	go printObject("Key 1")
	go addObject("Key 18", "Valor 1")
	go addObject("Key 19", "Valor 1")

	time.Sleep(time.Second *2)
}

func addObject(Key string, Value any){
	mutex.Lock() //cierra el recurso antes de ejecutarlo
	objects[Key] = Value
	mutex.Unlock() //lo desbloquea
}

func printObject(key string){
	mutex.Lock()
	value, ok:= objects[key]
	if ok{
		fmt.Println(value)
	}
	mutex.Unlock()
}

