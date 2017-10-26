package main
import "fmt"
func main() {
    var A, B, C, D, DIFERENCA int
    fmt.Scanln(&A)
    fmt.Scanln(&B)
    fmt.Scanln(&C)
    fmt.Scanln(&D)
    DIFERENCA = A*B-C*D
    fmt.Printf("DIFERENCA = %d\n", DIFERENCA)
}