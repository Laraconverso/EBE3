// Ejercicio 1 - Descuento
// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. 
// Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
// Crear la aplicación de acuerdo a los requerimientos.

package ej1

import (
    "fmt"
)

func calcularDescuento(precio float64, porcentajeDescuento float64) float64 {
    descuento := (precio * porcentajeDescuento) / 100
    precioConDescuento := precio - descuento
    return precioConDescuento
}

func ej1() {
    var precio, porcentajeDescuento float64

    fmt.Print("Ingrese el precio del producto: ")
    fmt.Scanln(&precio)

    fmt.Print("Ingrese el porcentaje de descuento: ")
    fmt.Scanln(&porcentajeDescuento)

    precioConDescuento := calcularDescuento(precio, porcentajeDescuento)

    fmt.Printf("El precio con el %.2f%% de descuento aplicado es: %.2f\n", porcentajeDescuento, precioConDescuento)
}
