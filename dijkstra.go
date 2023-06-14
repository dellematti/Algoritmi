package main

import (
  "fmt"
  "math"
)


type Grafo map[int][]Arco   //uso direttamente l int come chiave della mappa
type Arco struct {         //per un grafo pesato dove le chiavi dei vertici sono interi
  to int
  costo int
}


func dfs1(g grafo, v int, aux map[int]bool) {
   fmt.Println(v)
   aux[v] = true
   for _, v2 := range g[v] {
      if !aux[v2] {
	 dfs1(g, v2, aux)
	}
    }
}


//altro modo per fare un grafo
// type adjSet []int
// type graph struct {
// 	n   int
// 	adj []adjSet
// }
// func nuovoGrafo(n int) *graph {
// 	var g graph
// 	g.n = n
// 	g.adj = make([]adjSet, n)
// 	stampaGrafo(&g)
// 	return &g
// }

//altro modo
// type vertice struct {
// 	valore int
// 	key      string
// 	adj    []*vertice // insieme dei vertici adiacenti
// }
// type graph map[string]*vertice







func leggiGrafo() Grafo {
  var grafo Grafo
	grafo = make(map[int][]Arco)
  var numeroArchi, numeroVertici int
  fmt.Scan(&numeroVertici)  //non serve con la mappa
   _ = numeroVertici

  fmt.Scan(&numeroArchi)

  for i := 0; i < numeroArchi; i++{
    var from, to, costo int
    fmt.Scan(&from, &to, &costo)
    arcoAndata := Arco{to,costo}   // non è orientato, aggiungo in entrambe le direzioni
    arcoRitorno := Arco{from,costo}
    grafo[from] = append(grafo[from], arcoAndata)
    grafo[to] = append(grafo[to], arcoRitorno)
  }
	return grafo
}


func bfs(g Grafo, partenza int, arrivo int ) bool{
   coda := []int{partenza}
   visitati := make (map[int]bool)
   visitati[partenza] = true

   for len(coda) > 0 {
	vertice := coda[0]
	coda = coda[1:]

    	if vertice == arrivo{
      	   return true
    	}
	for _, arco := range g[vertice] {
	   if !visitati[arco.to]  {   //aggiungo solo i vertici che raggiungo con un arcco povero
		coda = append(coda, arco.to)
		visitati[arco.to] = true
	    }
	}
    }
  return false
}


// se decidessi di fare un minheap dovrei fare l heap con dentro gli stessi elementi di distanze ma con al suo interno, ogni
// elemento sarebbe rappresentato da una struct fatta dalla distanza e dalla chiave, qua se trovo il min in posizione 4 so già
// la posizione, con il miheap troverei sempre il minimo in posizione 1, ma mi serve sapere la chiave del vertice
func minoreDiDistanze(distanze []int, visitati map[int]bool ) int {  //senza l heap devo scorrere tutta la lista per trovare il minore
  min := math.MaxInt64
  posizione := 0
  for i , distanza := range distanze{
    if distanza < min && visitati[i] != true {
      min = distanza
      posizione = i
    }
  }
  return posizione
}//se usassi un heap per la coda con priorità, potrei anche non usare la mappa, farei remove (min) e non mi
// serve la mappa


//dijkstra con liste di adiacenza, senza minheap. Come coda con priorità uso direttamente la lista delle distanze dalla sorgente, e
// ogni volta cerco il vertice con minore distanza minore (che non ho ancora visitato)
// il grafo è implementato come mappa[int][]Archi , uso []Archi e non []int , perchè in struct arco tengo salvato anche il peso
func dijkstra ( grafo Grafo, partenza int, arrivo int )[]int {
  distanze := make([]int , len(grafo))
  for i, _ := range distanze {   //inizio mettendo tutte le distanze ad infinito
    distanze[i] = math.MaxInt64
  }
  distanze[partenza] = 0         //tranne il punto di partenza che è a 0
  visitati := make (map[int]bool)

  for len(visitati) < len(grafo){
    vertice := minoreDiDistanze(distanze, visitati) //cerco il vertice con minore distanza dalla partenza (ancora non visitato)
    visitati[vertice] = true
     for _, arco := range grafo[vertice] {  //guardo gli archi che escono da quel vertice
       destinazione := arco.to
       provenienza := vertice
       if distanze[destinazione] >  distanze[provenienza] + arco.costo { //rilasso
        distanze[destinazione] =  distanze[provenienza] + arco.costo
        }
      }
    }
  return distanze
}
//con la lista archi, la dovrei scorrere tutta guardando m archi per ogni vertice  O(m*n), così invece guardo solo gli archi di ogni
// vertice, e guardando tutti i vertici guarderò quindi tutti gli archi O (m)
//senza l heap devo scorrere tutti i vertici ogni volta quindi tempo totale O ( m*n), con l heap avrei O (m * log n)  //heap e lista adiacenza
// ps il numero di archi ( scritto come m oppure e ) è al massimo il numero di vertici (v oppure n) alla seconda



func main(){

  grafo := leggiGrafo()
	// stampaGrafo(g)

  var partenza, arrivo int

  fmt.Scan(&partenza)
  fmt.Scan(&arrivo)

  fmt.Println(partenza, arrivo)
  fmt.Println(bfs(grafo, partenza, arrivo))



  fmt.Println(dijkstra(grafo, partenza, arrivo))





}
