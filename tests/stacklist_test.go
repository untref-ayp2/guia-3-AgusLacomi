package tests

import (
	"guia3/linkedlist"
	"testing"
)

func TestPush(t *testing.T) {
	p := linkedlist.NewStack[int]()
	p.Push(1)
	p.Push(2)
	p.Push(3)
	if p.Size() != 3 {
		t.Errorf("Error en TestPush: tamaño incorrecto de la Pila, se esperaba %v, se obtuvo %v", 3, p.Size())
	}
}

func TestPop(t *testing.T) {
	p := linkedlist.NewStack[int]()
	p.Push(1)
	p.Push(2)
	p.Push(3)
	val, err := p.Pop()
	if err != nil {
		t.Errorf("Error en TestPop: se esperaba un valor, se obtuvo error: %v", err)
	}
	if val != 3 {
		t.Errorf("Error en TestPop: valor incorrecto obtenido de la Pila, se esperaba %v, se obtuvo %v", 3, val)
	}
	if p.Size() != 2 {
		t.Errorf("Error en TestPop: tamaño incorrecto de la Pila después de hacer pop, se esperaba %v, se obtuvo %v", 2, p.Size())
	}
}

func TestTop(t *testing.T) {
	p := linkedlist.NewStack[int]()
	p.Push(1)
	p.Push(2)
	p.Push(3)
	val, err := p.Top()
	if err != nil {
		t.Errorf("Error en TestTop: se esperaba un valor, se obtuvo error: %v", err)
	}
	if val != 3 {
		t.Errorf("Error en TestTop: valor incorrecto obtenido de la Pila, se esperaba %v, se obtuvo %v", 3, val)
	}
	if p.Size() != 3 {
		t.Errorf("Error en TestTop: tamaño incorrecto de la Pila después de hacer Top, se esperaba %v, se obtuvo %v", 3, p.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	p := linkedlist.NewStack[int]()
	if !p.IsEmpty() {
		t.Errorf("Error en TestIsEmpty: se esperaba que la Pila esté vacía, se obtuvo que no lo está")
	}
	p.Push(1)
	if p.IsEmpty() {
		t.Errorf("Error en TestIsEmpty: se esperaba que la Pila no esté vacía, se obtuvo que sí lo está")
	}
}

func TestSize(t *testing.T) {
	p := linkedlist.NewStack[int]()
	if p.Size() != 0 {
		t.Errorf("Error en TestSize: tamaño incorrecto de la Pila vacía, se esperaba %v, se obtuvo %v", 0, p.Size())
	}
	p.Push(1)
	if p.Size() != 1 {
		t.Errorf("Error en TestSize: tamaño incorrecto de la Pila después de agregar un elemento, se esperaba %v, se obtuvo %v", 1, p.Size())
	}
	p.Push(2)
	if p.Size() != 2 {
		t.Errorf("Error en TestSize: tamaño incorrecto de la Pila después de agregar dos elementos, se esperaba %v, se obtuvo %v", 2, p.Size())
	}
}
