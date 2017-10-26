package main
import "fmt"
func main() {
    var codigo, numero, valor, total float64
    total = 0
    for i:=0; i<2; i++{
        fmt.Scanln(&codigo,&numero,&valor)
        total += numero*valor
    }
    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}