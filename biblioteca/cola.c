#include "cola.h"
#include <stdlib.h>

typedef struct nodo nodo_t;

struct nodo{
    void* dato;
    nodo_t* siguiente;
};

// Devuelve un nodo con el dato ingresado como parametro y el puntero en null
nodo_t* crear_nodo(void* dato){
    nodo_t* nuevo_nodo = malloc(sizeof(nodo_t));
    nuevo_nodo->dato = dato;
    nuevo_nodo->siguiente = NULL;
    return nuevo_nodo;
}

void destruir_nodo(nodo_t* nodo){
    free(nodo);
}

struct cola
{
    nodo_t* primero;
    nodo_t* ultimo;
};

cola_t *cola_crear(void){
    cola_t* cola =  malloc(sizeof(cola_t));
    if(!cola) return NULL;
    cola->primero = NULL;
    cola->ultimo = NULL;
    return cola;
}

void cola_destruir(cola_t *cola, void (*destruir_dato)(void *)){
    while (!cola_esta_vacia(cola))
    {
        void* dato = cola_desencolar(cola);
        if(destruir_dato) destruir_dato(dato);
    }
    free(cola);
}

bool cola_esta_vacia(const cola_t *cola){
    return cola->primero == NULL;
}

bool cola_encolar(cola_t *cola, void *valor){
    nodo_t* nuevo_nodo = crear_nodo(valor);
    if(!nuevo_nodo) return false;
    nodo_t* ultimo_nodo = cola->ultimo;
    cola->ultimo = nuevo_nodo;

    if(cola_esta_vacia(cola)){
       cola->primero = nuevo_nodo;
    }

    if(ultimo_nodo != NULL){
        ultimo_nodo->siguiente = nuevo_nodo;
    }
    return true;
}

void *cola_ver_primero(const cola_t *cola){
    if(cola_esta_vacia(cola)){
        return NULL;
    }
    return cola->primero->dato;
}

void *cola_desencolar(cola_t *cola){
    void* dato = NULL;
    if(!cola_esta_vacia(cola)){
        nodo_t* primer_nodo = cola->primero;
        dato = primer_nodo->dato;
        cola->primero = primer_nodo->siguiente;
        destruir_nodo(primer_nodo);
    }
    if(cola_esta_vacia(cola)) cola->ultimo = NULL;
    
    return dato;
}