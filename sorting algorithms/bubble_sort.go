package main
import "fmt"

func swap(vetor *[]int, previous int, next int){
    aux := (*vetor)[next]
    (*vetor)[next] = (*vetor)[previous]
    (*vetor)[previous] = aux
}

func BubbleSort(vetor []int) {
    var size int = len(vetor)
    for i := 0; i < size - 1; i++ {
        for j := i + 1; j < size; j++ {
            if vetor[i] > vetor[j] {
                swap(&vetor, i, j)
            }
        }
    }
}


func main() {
    var vetor = []int{3, 5, 2, 7, 1, 8, 0}
    fmt.Println("Vetor de entrada")
    fmt.Println(vetor)
    BubbleSort(vetor)
    fmt.Println("\nVetor ordenado")
    fmt.Println(vetor)
}
