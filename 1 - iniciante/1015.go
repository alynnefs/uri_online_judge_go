package main
import "fmt"
import "math"
func main() {
    var x1, x2, y1, y2 float64
    fmt.Scanln(&x1, &y1)
    fmt.Scanln(&x2, &y2)
    fmt.Printf("%.4f\n", math.Sqrt((x2-x1)*(x2-x1)+(y2-y1)*(y2-y1)))
}