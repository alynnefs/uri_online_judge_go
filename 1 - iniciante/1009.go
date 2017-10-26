package main
import "fmt"
func main() {
    var nome string
    var salario, vendas float64
    fmt.Scanln(&nome)
    fmt.Scanln(&salario)
    fmt.Scanln(&vendas)
    fmt.Printf("TOTAL = R$ %.2f\n", salario+vendas*15/100)
}