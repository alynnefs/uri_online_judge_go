package main
import "fmt"
func main() {
    var a, b, c, maior int
    fmt.Scanln(&a,&b,&c)
    maior = a
    if maior < b && b > c {
        maior = b
    } else if maior < c && c > b {
        maior = c
    }
    fmt.Printf("%d eh o maior\n", maior)
}