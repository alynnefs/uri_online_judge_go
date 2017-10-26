package main
import "fmt"
func main() {
    var x int
    var y float64
    fmt.Scanln(&x)
    fmt.Scanln(&y)
    fmt.Printf("%.3f km/l\n", float64(x)/y)
}