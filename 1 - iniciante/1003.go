package main
import "fmt"
func main() {
    var A, B, SOMA int // SOMA foi declarado para diminuir o tempo de execução
    fmt.Scanln(&A)
    fmt.Scanln(&B)
    SOMA = A + B // se SOMA não estivesse declarada, a linha poderia ser SOMA := A + B
    fmt.Printf("SOMA = %d\n", SOMA)
}