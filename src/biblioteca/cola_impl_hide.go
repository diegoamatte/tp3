package biblioteca

import (
	"container/list"
)

type cola struct {
	list *list.List
}

func ColaCrear() Cola {
	cola := cola{list: list.New()}
	return &cola
}

func (cola *cola) Desencolar() interface{} {
	value := cola.list.Front().Value
	cola.list.Remove(cola.list.Front())
	return value
}

func (cola *cola) Encolar(dato interface{}) {
	cola.list.PushBack(dato)
}

func (cola *cola) EstaVacia() bool {
	return cola.list.Len() == 0
}

func (cola *cola) VerPrimero() interface{} {
	return cola.list.Front().Value
}
