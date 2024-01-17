package main
import "fmt"

func busca_linear(vetor []int, num int) bool{
    for i := 0; i < len(vetor); i++ {
        if vetor[i] == num {
            return true
        }
    }
    return false
}

func main() {
    var vetor = []int{3, 5, 12, 9, 0, 1, 3, 1, 6, 5, 4,7, 8, 9, 1, 0}
    
    // Retorna true
    fmt.Println(busca_linear(vetor, 9))
    // Retorna false
    fmt.Println(busca_linear(vetor, 15))
}
