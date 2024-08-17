package math

import "errors"

type Stack[T any] interface {
	Append(T)
	Pop() (T, error)
}

// Linked list
type LinkedList[T any] struct {
	Data     T
	Previous *LinkedList[T]
}

func (s *LinkedList[T]) Append(data T) {
	s = &LinkedList[T]{
		Data:     data,
		Previous: s,
	}
}

func (s *LinkedList[T]) Pop() (T, error) {
	if s == nil {
		return *new(T), errors.New("empty stack")
	} else {
		data := s.Data
		s = s.Previous
		return data, nil
	}
}
