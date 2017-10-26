package main
import "fmt"
func main() {
    var A, B, MEDIA float64
    fmt.Scanln(&A)
    fmt.Scanln(&B)
    MEDIA = (A * 3.5 + B * 7.5) / 11
    fmt.Printf("MEDIA = %.5f\n", MEDIA)
}