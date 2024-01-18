package main
import "fmt"

func insertion_sort(vetor []int){
    var j int
    for i := 1; i < len(vetor); i++{
        aux := vetor[i]; // 1 i = 3
        for j = i; (j > 0) && (aux < vetor[j - 1]); j--{ 
            vetor[j] = vetor[j - 1];
        }
        vetor[j] = aux;
    }
}

func main() {
    var vetor = []int{3, 5, 12, 9, 0, 1, 3, 1, 6, 5, 4,7, 8, 9, 1, 0}
    fmt.Println("Vetor de entrada")
    fmt.Println(vetor)
    insertion_sort(vetor)
    fmt.Println("\nVetor ordenado")
    fmt.Println(vetor)
}
