package main
import "fmt"
func main() {
    var num, horas, valor float64
    fmt.Scanln(&num)
    fmt.Scanln(&horas)
    fmt.Scanln(&valor)
    fmt.Printf("NUMBER = %.0f\n", num)
    fmt.Printf("SALARY = U$ %.2f\n", valor*horas)
}