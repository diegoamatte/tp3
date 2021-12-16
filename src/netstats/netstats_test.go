package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestNetstatDiametro(t *testing.T) {
	
	content := []byte("diametro\n")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) 
    if _, err := tmpfile.Write(content); err != nil {
        log.Fatal(err)
    }

    if _, err := tmpfile.Seek(0, 0); err != nil {
        log.Fatal(err)
    }

    oldStdin := os.Stdin
    defer func() { os.Stdin = oldStdin }() 
    os.Stdin = tmpfile

	net := Crear()
	net.CargarArchivo("../..//wiki-reducido-5000.tsv")
	net.EscucharComandos()


    if err := tmpfile.Close(); err != nil {
        log.Fatal(err)
    }
	
}
