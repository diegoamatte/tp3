package biblioteca

type nodo struct {
	dato      interface{}
	siguiente *nodo
}

type Pila struct {
	tope     *nodo
	cantidad int
}

func PilaCrear() *Pila {
	return &Pila{}
}

func (pila *Pila) EstaVacia() bool {
	return (*pila).cantidad == 0
}

func (pila *Pila) Apilar(dato interface{}) {
	nodo := &nodo{dato: dato}
	if pila.tope == nil {
		pila.tope = nodo
	} else {
		nodo.siguiente = pila.tope
		pila.tope = nodo
	}
	pila.cantidad++
}

func (pila *Pila) Desapilar() interface{} {
	if pila.EstaVacia() {
		return nil
	}
	nodo := *pila.tope
	pila.tope = nodo.siguiente
	pila.cantidad--
	return nodo.dato
}

func (pila *Pila) VerTope() interface{} {
	if pila.EstaVacia() {
		return nil
	}
	return pila.tope.dato
}
