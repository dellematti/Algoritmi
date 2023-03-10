package main
import "fmt"

// heap (max-heap) e heapsort


// scambia due elementi
func swap(vettorePosizionale *[]int, primo, secondo int) {
	(*vettorePosizionale)[primo], (*vettorePosizionale)[secondo] = (*vettorePosizionale)[secondo], (*vettorePosizionale)[primo]
}


func foglia(vettorePosizionale *[]int, indiceNodo int) bool {
	if indiceNodo >= (len(*vettorePosizionale)/2) && indiceNodo <= len(*vettorePosizionale) {
		return true
	}
	return false
}


// per risistemare (verso il basso) impiego al massimo log n, l altezza dello heap è logaritmica rispetto al numero di nodi
func risistemaVersoIlBasso(vettorePosizionale *[]int, indiceNodo int, size int) {
	if foglia(vettorePosizionale, indiceNodo) {  //l algoritmo funziona anche senza questo case base, ma mi
  	return                                     //permette di evitare l iterazione dell ultima chiamata ricorsiva
	}
  // per fare un min heap mi basterebbe prendere il minore invece del maggiore (sostituisco > con <)
	maggiore := indiceNodo
	indiceFiglioSx := 2 * indiceNodo + 1
	indiceFiglioDx := 2 * indiceNodo + 2
  // controllo se c è un figlio, e poi controllo se è più grande del nodo che sto guardando
	if indiceFiglioSx < size && (*vettorePosizionale)[indiceFiglioSx] > (*vettorePosizionale)[maggiore] {
		maggiore = indiceFiglioSx
	}
	if indiceFiglioDx < size && (*vettorePosizionale)[indiceFiglioDx] > (*vettorePosizionale)[maggiore] {
		maggiore = indiceFiglioDx
	}
  // se uno dei due figlio è maggiore del nodo, allora scambio e controllo che il nodo che ho messo al posto del
  // figlio, rispetti la condizione dell heap nel livello di profondità "sotto" rispetto a dove siamo ora
	if maggiore != indiceNodo {
		swap(vettorePosizionale, indiceNodo, maggiore)
		risistemaVersoIlBasso(vettorePosizionale, maggiore, size)
	}
	return
}


// la prima cosa da fare, prima di iniziare ad ordinare, è creare la struttura dati heap che uso in seguito (max-heap)
// tempo O (n)     questo anche se chiamo n volte risistema che costa O ( log n)
func creaMinHeap(vettorePosizionale *[]int) {
  dimensione := len(*vettorePosizionale)
	for i := (dimensione / 2) - 1; i >= 0; i-- {
		risistemaVersoIlBasso(vettorePosizionale, i, dimensione)
	}
}


// faccio n volte risistema che costa O ( log n )  ---->  O ( n log n)
func heapSort(vettorePosizionale *[]int) {
	creaMinHeap(vettorePosizionale)
  // ogni volta scambio con l ultima foglia e faccio risistema
	for i := len(*vettorePosizionale) - 1; i > 0; i-- {
		swap(vettorePosizionale, 0, i)
		risistemaVersoIlBasso(vettorePosizionale, 0, i)
	}
}


func main() {
	inputArray := []int{6, 5, 3, 7, 2, 8, -1, 34, 1 , 3, 30}
  heapSort( &inputArray )
  // passo l array per riferimento, devo controllare se in go ha senso oppure no; passandolo per valore mi copierebbe l header della
  //  slice ogni volta, il riferimento resta quello ma ogni volta devo copiare: length, capacità e il riferimento (comunque sarebbe tempo costante)
  for _, r := range inputArray{
    fmt.Print(r, " ")
  }
  fmt.Println()

}
