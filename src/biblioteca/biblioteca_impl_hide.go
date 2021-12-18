package biblioteca

import (
	"container/heap"
	"math"

	"github.com/algo2/tp3/grafo"
)

//variable para almacenar el pageRank
var pageRank map[interface{}]float64

//Struct utilizado para rastrear a que diccionario pertenece el maximo valor en la funcion de diametro
type padresType struct {
	padres     *map[interface{}]interface{}
	distancias *map[interface{}]int
	vertice    interface{}
}

// utilizado para calculo de CFC
var ordenCFC = 0
var visitadosCFC map[interface{}]bool
var result []interface{}

//Dado un diccionario de distaasnci, reconstruye el camino hacia el destino
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

func CaminoMasCorto(grafo *grafo.Grafo, origen, destino interface{}) ([]interface{}, int) {
	padres, distancias := bfs(grafo, origen)
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

func PageRank(grafo *grafo.Grafo, n int) []interface{} {
	if pageRank == nil {
		pageRank = calcularPageRank(grafo)
	}

	cp := ColaPrioridad{}
	heap.Init(&cp)
	for vertice, valor := range pageRank {
		elem := Elemento{dato: vertice, prioridad: valor}
		heap.Push(&cp, &elem)
	}
	var top []interface{}
	for i := 0; i < n; i++ {
		if cp.Len() > 0 {
			max := heap.Pop(&cp).(*Elemento)
			top = append(top, (*max).dato)
		}
	}
	return top
}

func calcularPageRank(grafo *grafo.Grafo) map[interface{}]float64 {
	d := 0.85
	pr := make(map[interface{}]float64)
	prAnterior := make(map[interface{}]float64)
	vertices := (*grafo).ObtenerVertices()
	cantVertices := float64(len(vertices))
	delta := 1.

	for _, v := range vertices {
		pr[v] = (1 - d) / cantVertices
	}

	for delta > 0.01 {
		for _, v := range vertices {
			prAnterior[v] = pr[v]
		}
		pr = make(map[interface{}]float64)

		for _, v := range vertices {
			adyacentes := (*grafo).ObtenerAdyacentes(v)
			for _, ady := range adyacentes {
				pr[ady] += d * prAnterior[v] / float64(len(adyacentes))
			}
		}
		delta = 0
		//Convergencia
		for _, v := range vertices {
			delta += math.Abs(pr[v] - prAnterior[v])
		}
	}

	return pr
}

func Diametro(grafo *grafo.Grafo) ([]interface{}, int) {
	cp := ColaPrioridad{}
	heap.Init(&cp)

	for _, v := range (*grafo).ObtenerVertices() { //O(V)
		padres, orden := bfs(grafo, v)          //O(V+E))
		for vertice, distancia := range orden { //O(V)
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

func Rango(grafo *grafo.Grafo, v interface{}, n int) int {
	solucion := make([]interface{}, 0)
	_, orden := bfs(grafo, v)
	for i := range orden {
		if orden[i] == n {
			solucion = append(solucion, orden[i])
		}
	}
	return len(solucion)
}

func Navegacion(grafo *grafo.Grafo, origen interface{}, n int) []interface{} {

	padres := make(map[interface{}]interface{})
	orden := make(map[interface{}]int)
	pasos := 0
	padres[origen] = nil
	solucion := navRec(grafo, origen, padres, orden, pasos, n)
	return reconstruirCamino(padres, orden, solucion)
}

func navRec(grafo *grafo.Grafo, origen interface{}, padres map[interface{}]interface{}, orden map[interface{}]int, pasos, n int) interface{} {
	if pasos >= n {
		return origen
	}
	for _, ady := range (*grafo).ObtenerAdyacentes(origen) {
		padres[ady] = origen
		orden[ady] = orden[origen] + 1
		return navRec(grafo, ady, padres, orden, pasos+1, n)
	}
	return origen
}

func Conectividad(g *grafo.Grafo, pagina interface{}) []interface{} {

	if _, pertenece := visitadosCFC[pagina]; !pertenece {
		visitadosCFC = make(map[interface{}]bool)
		pila := PilaCrear()
		ordenCFC = 0
		mb := make(map[interface{}]int)
		orden := make(map[interface{}]int)
		apilados := make(map[interface{}]bool)
		result = conectividad(g, pagina, pagina, pila, orden, visitadosCFC, apilados, mb)
	}
	return result
}

func conectividad(g *grafo.Grafo, pagina, buscado interface{}, pila *Pila, orden map[interface{}]int, visitados, apilados map[interface{}]bool, mb map[interface{}]int) []interface{} {
	visitados[pagina] = true
	orden[pagina] = ordenCFC
	mb[pagina] = orden[pagina]
	ordenCFC += 1
	(*pila).Apilar(pagina)
	apilados[pagina] = true

	for _, ady := range (*g).ObtenerAdyacentes(pagina) {
		if _, ok := visitados[ady]; !ok {
			conectividad(g, ady, buscado, pila, orden, visitados, apilados, mb)
		}
		if _, ok := apilados[ady]; ok {
			mb[pagina] = int(math.Min(float64(mb[pagina]), float64(mb[ady])))

		}
	}

	if orden[pagina] == mb[pagina] && !pila.EstaVacia() {
		cfc := make([]interface{}, 0)
		var w interface{}
		for {
			w = (*pila).Desapilar()
			delete(apilados, w)
			cfc = append(cfc, w)
			if pagina == w {
				break
			}
		}
		if pagina == buscado {
			return cfc
		}
	}
	return nil
}
