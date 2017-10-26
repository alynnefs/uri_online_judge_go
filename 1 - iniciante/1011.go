package main
import "fmt"
func main() {
    var r float64
    const pi = 3.14159
    fmt.Scanln(&r)
    fmt.Printf("VOLUME = %.3f\n", (4.0/3.0)*pi*r*r*r)
}