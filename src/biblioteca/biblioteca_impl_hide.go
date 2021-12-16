package biblioteca

import (
	"container/heap"
	"fmt"

	"github.com/algo2/tp3/grafo"
)

//Struct utilizado para rastrear a que diccionario pertenece el maximo valor en la funcion de diametro
type padresType struct {
	padres     *map[interface{}]interface{}
	distancias *map[interface{}]int
	vertice    interface{}
}

//Dado un diccionario de distancias, reconstruye el camino hacia el destino
func reconstruirCamino(padres map[interface{}]interface{}, distancias map[interface{}]int, destino interface{}) []interface{} {
	solucion := make([]interface{}, 0, distancias[destino])
	anterior, ok := padres[destino]
	if !ok {
		return nil
	}
	solucion = append(solucion, destino)
	for i := 0; i < distancias[destino]; i++ {
		solucion = append([]interface{}{anterior}, solucion...)
		anterior = padres[anterior]
	}
	return solucion
}

func CaminoMasCorto(grafo *grafo.Grafo, origen interface{}, destino interface{}) ([]interface{}, int) {
	padres, distancias := bfs(grafo, origen, nil, nil)
	solucion := make([]interface{}, 0, distancias[destino])
	anterior, ok := padres[destino]
	if !ok {
		return nil, 0
	}
	solucion = append(solucion, destino)
	for i := 0; i < distancias[destino]; i++ {
		solucion = append([]interface{}{anterior}, solucion...)
		anterior = padres[anterior]
	}

	return solucion, distancias[destino]
}

func PageRank(grafo *grafo.Grafo, n int) {
	pr := calcularPageRank(grafo)
	cp := ColaPrioridad{}
	heap.Init(&cp)
	for vertice, valor := range pr {
		elem := Elemento{dato: vertice, prioridad: valor}
		cp.Push(&elem)
	}
	var top []interface{}
	for i := 0; i < n; i++ {
		max := cp.Pop().(*Elemento)
		top = append(top, *max)
	}
	for _, v := range top {
		fmt.Println(v)
	}

}

func calcularPageRank(grafo *grafo.Grafo) map[interface{}]float64 {

	return nil
}

func Diametro(grafo *grafo.Grafo) ([]interface{}, int) {
	cp := ColaPrioridad{}
	heap.Init(&cp)

	for _, v := range (*grafo).ObtenerVertices() { //O(V)
		padres, orden := bfs(grafo, v, nil, nil) //O(V+E))
		for vertice, distancia := range orden {  //O(V)
			dato := padresType{padres: &padres, vertice: vertice, distancias: &orden}
			elem := Elemento{dato: dato, prioridad: float64(distancia)}
			heap.Push(&cp, &elem) //O(log V)
		}
	}
	maxElemento := heap.Pop(&cp).(*Elemento)
	max := int(maxElemento.prioridad)
	maxDato := maxElemento.dato.(padresType)

	solucion := reconstruirCamino(*maxDato.padres, *maxDato.distancias, maxDato.vertice)

	return solucion, max
}
