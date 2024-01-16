package main
import "fmt"

func Swap(vetor []int, atual int, min int) {
    vetor[atual], vetor[min] = vetor[min], vetor[atual]
}

func Selection_Sort(vetor []int){
    var vetor_size int = len(vetor)
    
    for i := 0; i < vetor_size - 1; i++{
        idx_min := i
        for j := i + 1; j < vetor_size; j++ {
            if vetor[j] < vetor[idx_min] {
                idx_min = j
            }
        }
        Swap(vetor, i, idx_min)
    } 
}
func main() {
    var vetor = []int{3, 5, 12, 9, 0, 1, 3, 1, 6, 5, 4,7, 8, 9, 1, 0}
    fmt.Println("Vetor de entrada")
    fmt.Println(vetor)
    Selection_Sort(vetor)
    fmt.Println("\nVetor ordenado")
    fmt.Println(vetor)
}
