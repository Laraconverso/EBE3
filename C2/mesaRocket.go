package main

import "fmt"

func CalcularSalario(minutosTrabajadosPorMes int, categoria string) float64 {
    var tarifaHora, salarioMensual float64

    switch categoria {
    case "C":
        tarifaHora = 1000.0
    case "B":
        tarifaHora = 1500.0
    case "A":
        tarifaHora = 3000.0
    default:
        fmt.Println("Categoría no válida")
        return 0.0
    }

    salarioMensual = tarifaHora * float64(minutosTrabajadosPorMes) / 60.0

    if categoria == "B" {
        salarioMensual += (salarioMensual * 0.20) // Aumento del 20% para la categoría B
    } else if categoria == "A" {
        salarioMensual += (salarioMensual * 0.50) // Aumento del 50% para la categoría A
    }

    return salarioMensual
}

func main() {
    minutosTrabajados := 1800 // Cambia esto con la cantidad de minutos trabajados
    categoria := "B"         // Cambia esto con la categoría del empleado (C, B o A)

    salario := calcularSalario(minutosTrabajados, categoria)
    fmt.Printf("El salario mensual del empleado de categoría %s es: $%.2f\n", categoria, salario)
}
