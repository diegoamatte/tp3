package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/algo2/tp3/biblioteca"
	"github.com/algo2/tp3/grafo"
)

var operacionesStr = "camino\ndiametro\nrango"

type netstatsType struct {
	grafo grafo.Grafo
}

func Crear() Netstats {
	return new(netstatsType)
}

func (netstats *netstatsType) CargarArchivo(path string) {
	file, error := os.Open(path)
	defer file.Close()
	if error != nil {
		fmt.Print("No se pudo abrir el archivo")
		return
	}
	data := formatter('\t', file)
	netstats.crearGrafo(data)
	file.Close()
}

//Crea un grafo a partir de una matriz
func (netstats *netstatsType) crearGrafo(data [][]string) {
	graph := grafo.Crear(true)
	for _, comp := range data {
		for i, dat := range comp {
			graph.AgregarVertice(dat)
			if i > 0 {
				graph.AgregarArista(comp[0], comp[i], 1)
			}
		}
	}
	netstats.grafo = graph
}

// Funcion auxiliar para obtener una matriz con los datos a partir del archivo.
func formatter(delimiter rune, file *os.File) [][]string {
	reader := csv.NewReader(file)
	reader.Comma = delimiter //'\t'
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo.")
		return nil
	}
	return data
}

func (netstat *netstatsType) EscucharComandos() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		texto := scanner.Text()
		if len(texto) != 0 {
			netstat.ejecutar(texto)
		} else {
			break
		}
	}
}

func (netstat *netstatsType) ejecutar(linea string) {
	input := strings.SplitN(linea, " ", 2)
	comando := input[0]
	switch comando {
	case "listar_operaciones":
		listarOperaciones()
	case "camino":
		args := strings.SplitN(input[1], ",", 2)
		netstat.camino(args[0], args[1])
	case "mas_importantes":
		arg, _ := strconv.Atoi(input[1])
		netstat.pageRank(arg)
	case "diametro":
		netstat.diametro()
	case "rango":
		args := strings.SplitN(input[1], ",", 2)
		n, _ := strconv.Atoi(args[1])
		netstat.rango(args[0], n)
	}
}

func listarOperaciones() {
	fmt.Println(operacionesStr)
}

func (netstat *netstatsType) camino(origen string, destino string) {
	result, costo := biblioteca.CaminoMasCorto(&netstat.grafo, origen, destino)
	if costo == 0 {
		fmt.Println("No se encontro recorrido")
		return
	}
	fmt.Println(salidaFormato1(result, costo))
}

func (netstat *netstatsType) pageRank(n int) {
	biblioteca.PageRank(&netstat.grafo, 20)
}

func (netstat *netstatsType) diametro() {
	diametro, costo := biblioteca.Diametro(&netstat.grafo)
	fmt.Println(salidaFormato1(diametro, costo))
}

func (netstat *netstatsType) rango(origen string, n int) {
	resultado := biblioteca.Rango(&netstat.grafo, origen, n)
	fmt.Println(resultado)
}
func salidaFormato1(solucion []interface{}, costo int) string {
	var sb strings.Builder
	for i, str := range solucion {
		sb.WriteString(fmt.Sprintf("%s", str))
		if i < len(solucion)-1 {
			sb.WriteString(" -> ")
		}
	}
	sb.WriteString(fmt.Sprintf("\nCosto: %d", costo))
	return sb.String()
}
