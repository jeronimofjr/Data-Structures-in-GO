package main
import "fmt"

func busca_binaria(vetor []int, num int) int {
    pos_ini := 0
    pos_fim := len(vetor) - 1
    for ; pos_ini <= pos_fim; {
        pos_mei := (pos_ini + pos_fim) / 2
        
        if vetor[pos_mei] == num {
            return pos_mei
        } else if vetor[pos_mei] > num{
             pos_fim = pos_mei - 1
        } else if vetor[pos_mei] < num {
            pos_ini = pos_mei + 1
        }
    }
    return -1
}

func main() {
    var vetor = []int{10, 11, 12, 13, 14, 15,16, 17, 18}
    
    // Retorna o índice do vetor onde está o valor que buscamos
    fmt.Println(busca_binaria(vetor, 12))
    // Retorna índice negativo, valor não está contido no vetor
    fmt.Println(busca_binaria(vetor, 1))
}
