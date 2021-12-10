package biblioteca

// #cgo CFLAGS: -g -Wall
// #include "cola.h"
import "C"
import (
	"grafo"
	"unsafe"
)

type bibliotecaType struct {
	grafo grafo.Grafo
}

func Crear(grafo grafo.Grafo) Biblioteca {
	biblioteca := bibliotecaType{grafo: grafo}
	return &biblioteca
}

func (biblio *bibliotecaType) CaminoMasCorto(origen interface{}, destino interface{}) ([]interface{}, bool) {
	visitados := make(map[interface{}]bool)
	padres := make(map[interface{}]interface{})
	distancias := make(map[interface{}]int)
	visitados[origen] = true
	padres[origen] = nil
	distancias[origen] = 0
	cola := C.cola_crear()
	defer C.cola_destruir(cola, nil)
	C.cola_encolar(cola, unsafe.Pointer(&origen))
	for !C.cola_esta_vacia(cola) {
		v := *(*interface{})(C.cola_desencolar(cola))
		for _, adyacente := range biblio.grafo.ObtenerAdyacentes(v) {
			_, visitado := visitados[adyacente]
			if !visitado {
				padres[adyacente] = v
				distancias[adyacente] = distancias[v] + 1
				visitados[adyacente] = true
				C.cola_encolar(cola, unsafe.Pointer(&adyacente))
			}
		}
	}
	solucion := make([]interface{}, 0, distancias[destino])
	anterior, ok := padres[destino]
	if !ok {
		return nil, false
	}
	solucion = append(solucion, destino)
	for i := 0; i < distancias[destino]; i++ {
		solucion = append([]interface{}{anterior}, solucion...)
		anterior = padres[anterior]
	}

	return solucion, true
}
