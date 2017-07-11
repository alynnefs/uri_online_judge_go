package main
import "fmt"
func main() {
    const n = 3.14159
    var raio float64
    fmt.Scanln(&raio)
    area := n * (raio * raio)
    fmt.Printf("A=%.4f\n", area)
}