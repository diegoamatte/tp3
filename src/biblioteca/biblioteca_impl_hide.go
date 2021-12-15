package biblioteca

import (
	"container/heap"
	"fmt"
	"math/rand"
	"reflect"

	"github.com/algo2/tp3/grafo"
)

type bfsAccion func(map[interface{}]interface{}, map[interface{}]int) (interface{}, interface{})

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
	pad,dist:= bfs(grafo,origen,nil)
	distancias := dist.(map[interface{}]int)
	padres := pad.(map[interface{}]interface{})
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
	pr := calcularPageRank(grafo, 100000, 0.85)
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
	fmt.Printf("%v+\n", top)
}

func calcularPageRank(grafo *grafo.Grafo, iteraciones int, dampingFactor float64) map[interface{}]float64 {
	pageRank := make(map[interface{}]float64)
	vertices := (*grafo).ObtenerVertices()
	for i := 0; i < iteraciones; i++ {
		indiceInicio := rand.Intn(len(vertices) - 1)
		calcularPRGrafo(grafo, vertices[indiceInicio], 1., dampingFactor, pageRank, 1)
	}
	return pageRank
}

func calcularPRGrafo(grafo *grafo.Grafo, vertice interface{}, transferencia float64, df float64, pageRank map[interface{}]float64, pasos int) {
	adyacentes := (*grafo).ObtenerAdyacentes(vertice)
	gradoSalida := float64(len(adyacentes))

	if (pasos >= 10000) || gradoSalida == 0. {
		return
	}
	var prAcumulado float64
	for _, ady := range adyacentes {
		_, ok := pageRank[ady]
		if ok {
			prAcumulado = pageRank[ady] * df
		} else {
			prAcumulado = 0
		}
		pageRank[ady] = prAcumulado + (1-df)/gradoSalida
	}
	indiceSiguiente := rand.Intn(len(adyacentes))
	verticeSiguiente := adyacentes[indiceSiguiente]
	calcularPRGrafo(grafo, verticeSiguiente, pageRank[verticeSiguiente], df, pageRank, pasos+1)
}

func Diametro(grafo *grafo.Grafo) ([]interface{}, int) {
	var diametro interface{}
	max := 0
	for _, v := range (*grafo).ObtenerVertices() {
		recorrido, localmax := bfs(grafo, v, bfsAccion(maximo))
		if localmax.(int) > max {
			max = localmax.(int)
			diametro = recorrido
		}
	}
	vertices := reflect.ValueOf(diametro)
	solucion := make([]interface{}, 0,vertices.Len())
	for i := 0; i < vertices.Len(); i++ {
		solucion = append(solucion, vertices.Index(i))
	}

	return solucion, max
}

func maximo(padres map[interface{}]interface{}, orden map[interface{}]int) (interface{}, interface{}) {
	maximo := 0
	var claveMax interface{}
	for clave, valor := range orden {
		if maximo < valor {
			maximo = valor
			claveMax = clave
		}
	}
	camino := reconstruirCamino(padres, orden, claveMax)
	return camino, maximo
}

//Devuelve el mapa de padres y de orden
func bfs(grafo *grafo.Grafo, origen interface{}, fn bfsAccion) (interface{}, interface{}) {
	cola := ColaCrear()
	visitados := make(map[interface{}]int)
	padres := make(map[interface{}]interface{})
	orden := make(map[interface{}]int)

	visitados[origen] = 1
	padres[origen] = nil
	orden[origen] = 0

	cola.Encolar(origen)

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, ady := range (*grafo).ObtenerAdyacentes(v) {
			if _, ok := visitados[ady]; !ok { //Si no fue visitado
				cola.Encolar(ady)
				orden[ady] = orden[v] + 1
				padres[ady] = v
				visitados[ady] = 1
			}
		}
	}
	if fn == nil{
		return padres,orden
	}
	return fn(padres, orden)
}
