package main
import "fmt"
func main() {
    var A, B, PROD int // PROD foi declarado para diminuir o tempo de execução
    fmt.Scanln(&A)
    fmt.Scanln(&B)
    PROD = A * B // se PROD não estivesse declarado, a linha poderia ser PROD := A * B
    fmt.Printf("PROD = %d\n", PROD)
}