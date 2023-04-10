package linkedlist

import (
	"errors"
	"fmt"
)

// node es el nodo de la lista enlazada
// contiene un valor y un puntero al siguiente nodo
// el valor es de tipo genérico, comparable
type node[T any] struct {
	value T
	next  *node[T]
}

// newNode crea un nuevo nodo, con el valor recibido
// y el puntero al siguiente nodo en nil
func newNode[T comparable](value T) *node[T] {
	return &node[T]{value: value, next: nil}
}

/************************************************************/

// LinkedList es la lista enlazada simple
// contiene punteros al primer nodo y al último
type LinkedList[T comparable] struct {
	head *node[T] // puntero al primer nodo
	tail *node[T] // puntero al último nodo

	size int // by AgusLacomi Punto 1
}

// NewLinkedList crea una nueva lista enlazada, vacía
// En nuestra implementación representamos la lista vacía
// con un puntero al primer nodo en nil
// y un puntero al último nodo en nil
// O(1)
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil}
}

// Append agrega un nuevo nodo, con el valor recibido, al final de la lista
// O(1)
func (l *LinkedList[T]) Append(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size++ // by AgusLacomi Punto 1
		return
	}
	l.size++ // by AgusLacomi Punto 1
	l.tail.next = newNode
	l.tail = newNode
}

// Prepend agrega un nuevo nodo, con el valor recibido,
// al inicio de la lista
// O(1)
func (l *LinkedList[T]) Prepend(value T) {
	newNode := newNode(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size++ // by AgusLacomi Punto 1
		return
	}
	l.size++ // by AgusLacomi Punto 1
	newNode.next = l.head
	l.head = newNode
}

// InsertAt agrega un nuevo nodo, con el valor recibido,
// en la posición recibida.
// Si la posición es inválida, no hace nada
// O(n)
func (l *LinkedList[T]) InsertAt(value T, position int) {
	if position < 0 {
		return
	}
	newNode := newNode(value)
	if position == 0 {
		l.Prepend(value)
		return
	}
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}
	if current == nil {
		return
	}
	l.size++ // by AgusLacomi Punto 1
	newNode.next = current.next
	current.next = newNode
}

// Remove elimina el primer nodo que contenga el valor recibido
// O(n)
func (l *LinkedList[T]) Remove(value T) {
	if l.head == nil {
		return // no hay nada que eliminar
	}
	if l.head.value == value {
		l.head = l.head.next
		l.size-- // by AgusLacomi Punto 1
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.size-- // by AgusLacomi Punto 1
			return
		}
		current = current.next
	}
}

// String devuelve una representación en cadena de la lista
// en el formato [1 2 3].
// Se puede usar para imprimir la lista con fmt.Println
// O(n)
func (l *LinkedList[T]) String() string {
	if l.head == nil {
		return "[]"
	}
	current := l.head
	result := "["
	for current != nil {
		result += fmt.Sprintf("%v", current.value)
		if current.next != nil {
			result += " "
		}
		current = current.next
	}
	result += "]"
	return result
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
// O(n)
func (l *LinkedList[T]) Search(value T) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.value == value {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *LinkedList[T]) Get(position int) (T, error) {
	if l.head == nil {
		var t T
		return t, errors.New("Lista vacía")
	}
	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}
	if current == nil {
		var t T
		return t, errors.New("Posición inválida")
	}
	return current.value, nil
}

// Size devuelve la cantidad de nodos en la lista
// O(1)
// By AgusLacomi Punto 1
func (l *LinkedList[T]) Size() int {
	return l.size
}

/************************************************************/

type Stack[T comparable] struct {
	head *node[T] // puntero al primer nodo
	size int
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{head: nil}
}

func (p *Stack[T]) Push(value T) {
	newNode := newNode(value)
	if p.head == nil {
		p.head = newNode
		p.size++
	} else {
		newNode.next = p.head
		p.head = newNode
		p.size++
	}
}

func (p *Stack[T]) Pop() (T, error) {

	if p.head == nil {
		var t T
		return t, errors.New("Nada que eliminar")
	}
	retorno := p.head
	p.head = p.head.next
	p.size--
	return retorno.value, nil
}

func (p *Stack[T]) IsEmpty() bool {

	if p.head != nil {
		return false
	}
	return true
}

func (p *Stack[T]) Top() (T, error) {
	if p.head == nil {
		var t T
		return t, errors.New("Nada en la cabezera")
	}
	return p.head.value, nil
}

func (p *Stack[T]) Size() int {
	return p.size
}

/************************************************************/

type Queue[T comparable] struct {
	head *node[T] // puntero al primer nodo
	tail *node[T] // puntero al último nodo
	size int
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{head: nil, tail: nil}
}

func (p *Queue[T]) Enqueue(value T) {
	newNode := newNode(value)
	if p.tail == nil {
		p.head = newNode
		p.tail = newNode
		p.size++
	} else {
		p.tail.next = newNode
		p.tail = newNode
		p.size++
	}
}

func (p *Queue[T]) Dequeue() (T, error) {
	if p.head == nil {
		var t T
		return t, errors.New("Nada que eliminar")
	}
	retorno := p.head
	p.head = p.head.next
	p.size--
	return retorno.value, nil
}

func (p *Queue[T]) IsEmpty() bool {
	if p.tail != nil {
		return false
	}
	return true
}

func (p *Queue[T]) Front() (T, error) {
	if p.head == nil {
		var t T
		return t, errors.New("Nada que eliminar")
	}
	return p.head.value, nil
}
