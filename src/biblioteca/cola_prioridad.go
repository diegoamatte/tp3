package biblioteca

type ColaPrioridad []*Elemento

type Elemento struct {
	dato interface{}
	prioridad float64
	indice int
}

//funcion de comparacion
func (cp ColaPrioridad) Less(i, j int) bool {
	return cp[i].prioridad > cp[j].prioridad
}

func (cp ColaPrioridad) Len() int { 
	return len(cp) 
}

func (cp ColaPrioridad) Swap(i, j int) {
	cp[i], cp[j] = cp[j], cp[i]
	cp[i].indice = i
	cp[j].indice = j
}

func (cp *ColaPrioridad) Push(x interface{}) {
	n := len(*cp)
	elemento := x.(*Elemento)
	elemento.indice = n
	*cp = append(*cp, elemento)
}

func (cp *ColaPrioridad) Pop() interface{} {
	anterior := *cp
	n := len(anterior)
	item := anterior[n-1]
	anterior[n-1] = nil  
	item.indice = -1 
	*cp = anterior[0 : n-1]
	return item
}