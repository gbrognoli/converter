package main

import (
    "fmt"
    "math"
)

func main() {
    var real float64
    fmt.Print("Por favor, insira o valor em real: ")
    fmt.Scanf("%f", &real)

    dollar := real * 0.26
    fmt.Println("O valor em dólar é:", dollar)
}
