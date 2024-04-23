package math

import "errors"

// Linked list
type Stack[T any] struct {
	Data     T
	Previous *Stack[T]
}

func (s *Stack[T]) Append(data T) {
	s = &Stack[T]{
		Data:     data,
		Previous: s,
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if s == nil {
		return *new(T), errors.New("empty stack")
	} else {
		data := s.Data
		s = s.Previous
		return data, nil
	}
}
