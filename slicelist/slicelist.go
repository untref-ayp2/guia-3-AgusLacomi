package slicelist

import (
	"errors"
	"fmt"
)

type Slicelist[T comparable] struct {
	list []T
}

// O(1)
func NewSliceList[T comparable]() *Slicelist[T] {
	return &Slicelist[T]{}
}

// O(1)
func (s *Slicelist[T]) Append(value T) {
	s.list = append(s.list, value)
}

// O(1)
func (s *Slicelist[T]) Prepend(value T) {
	s.list = append([]T{value}, s.list...)
}

// O(n)
func (s *Slicelist[T]) InsertAt(value T, position int) {
	if position < 0 || position >= len(s.list) {
		return
	}
	if position == 0 {
		s.Prepend(value)
		return
	}
	aux := s.list[position:]
	s.list = s.list[0 : position-1]
	s.list = append(s.list, value)
	s.list = append(s.list, aux...)
}

// O(n)
func (s *Slicelist[T]) Remove(value T) {
	for i := 0; i < len(s.list); i++ {
		if value == s.list[i] {
			s.list = append(s.list[:i], s.list[i+1:]...)
			return
		}
	}
}

// O(n)
func (s *Slicelist[T]) Search(value T) int {
	for i := 0; i < len(s.list); i++ {
		if value == s.list[i] {
			return i
		}
	}
	return -1
}

// O(1)
func (s *Slicelist[T]) Get(position int) (T, error) {
	if position < 0 || position >= len(s.list) {
		var t T
		return t, errors.New("Posición Inválida")
	}
	return s.list[position], nil
}

// O(1)
func (s *Slicelist[T]) Size() int {
	return len(s.list)
}

// O(n)
func (s *Slicelist[T]) String() string {
	if len(s.list) == 0 {
		return "[]"
	}
	result := "[ "
	for _, value := range s.list {
		result += fmt.Sprintf("%v", value)
		result += " "
	}
	result += "]"
	return result
}
