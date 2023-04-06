package main

import (
	"fmt"
	"guia3/linkedlist"
)

func main() {
	l := linkedlist.NewLinkedList[int]()
	fmt.Println("Agregamos 1, 2 y 3 al final de la lista")
	l.Append(1)
	l.Append(2)
	l.Append(3)
	fmt.Println(l)
	fmt.Println("Agregamos 0 al inicio de la lista")
	l.Prepend(0)
	fmt.Println(l)
	fmt.Println("Buscamos el numero 3")
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))

	l.Remove(2)
	l.Remove(0)
	l.Remove(3)
	fmt.Println("Buscamos el numero 3 luego de removerlo de la lista")
	fmt.Println(l)
	fmt.Println("Se encuentra en la posicion: ", l.Search(3))
	l.Remove(1)
	l.Remove(1) //No deberia hacer nada
	l.Remove(1) //No deberia hacer nada
	fmt.Println(l)

	fmt.Println("Agregamos 0, 1 , 3, 4 y 5 al final de la lista")
	l.Append(0)
	l.Append(1)
	l.Append(3)
	l.Append(4)
	l.Append(5)

	fmt.Println(l)
	fmt.Println("Agregamos 2 en la posicion 2")
	l.InsertAt(2, 2)

	fmt.Println(l)
	fmt.Println("Tamaño", l.Size()) // by AgusLacomi Punto 1

	/****************Punto Tres*******************************/
	l1 := linkedlist.NewLinkedList[int]()
	l1.Append(10)
	l1.Append(20)
	l1.Append(30)

	l3 := ConcatenarLinkedList[int](l, l1)
	fmt.Println(l3.String())
	/****************Punto Tres*******************************/

	/****************Punto Cuatro*******************************/

	l4 := IntercalarLinkedList[int](l, l1)
	fmt.Println(l4.String())

	/****************Punto Cuatro*******************************/
}

/**
 *	3. Escribir una función que reciba como parámetros dos listas del mismo tipo y devuelva una lista resultante de concatenar la segunda lista al final de la primera.
 *
 *	By AgusLacomi
 */
func ConcatenarLinkedList[T comparable](l1 *linkedlist.LinkedList[T], l2 *linkedlist.LinkedList[T]) *linkedlist.LinkedList[T] {

	aux := linkedlist.NewLinkedList[T]()

	for i := 0; i < l1.Size(); i++ {

		dato, _ := l1.Get(i)

		aux.Append(dato)
	}

	for i := 0; i < l2.Size(); i++ {
		dato, _ := l2.Get(i)

		aux.Append(dato)
	}

	// [0 1 2 3 4 5 10 20 30]

	return aux
}

/**
 *
 * 4. Escribir una función que reciba dos listas del mismo tipo y devuelva la lista que resulta de intercalar uno a uno los elementos de ambas listas.
 *
 * By AgusLacomi
 */
func IntercalarLinkedList[T comparable](l1 *linkedlist.LinkedList[T], l2 *linkedlist.LinkedList[T]) *linkedlist.LinkedList[T] {

	var size int
	var dato, _ T
	var listaMayor int

	paux := *l2

	if l2.Size() >= l1.Size() {
		size = l1.Size()
		listaMayor = l2.Size()
	} else {
		size = l2.Size() - 1
		listaMayor = l1.Size()
		paux = *l1
	}

	aux := linkedlist.NewLinkedList[T]()

	for i := 0; i < size; i++ {
		dato, _ = l1.Get(i)
		aux.Append(dato)
		dato, _ = l2.Get(i)
		aux.Append(dato)
	}

	for i := size; i < listaMayor; i++ {
		dato, _ = paux.Get(i)
		aux.Append(dato)
	}

	return aux
}
