package main

import (
	"fmt"
	"strings"
)

func salidaFormato1(solucion []interface{}, costo int) string {
	var sb strings.Builder
	sb.WriteString(salidaFormato2(solucion, " -> "))
	sb.WriteString(fmt.Sprintf("\nCosto: %d", costo))
	return sb.String()
}

func salidaFormato2(solucion []interface{}, separador string) string {
	var sb strings.Builder
	for i, str := range solucion {
		sb.WriteString(fmt.Sprintf("%s", str))
		if i < len(solucion)-1 {
			sb.WriteString(separador)
		}
	}
	return sb.String()
}