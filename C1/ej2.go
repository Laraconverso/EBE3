// Ejercicio 2 - Préstamo
// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
// Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
// Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.

package main

import "fmt"

func main() {
    // Variables
    edad := 25
    empleado := true
    antiguedad := 2
    sueldo := 120000

    // Verificar las condiciones
    if edad > 22 && empleado && antiguedad > 1 {
        if sueldo > 100000 {
            fmt.Println("¡Felicidades! Eres elegible para un préstamo sin intereses.")
        } else {
            fmt.Println("Eres elegible para un préstamo, pero se te aplicará interés.")
        }
    } else {
        fmt.Println("Lo siento, no eres elegible para un préstamo en este momento.")
    }
}
