package util

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	inner []interface{}
}

func NewStack() *Stack {
	return &Stack{inner: make([]interface{}, 0)}
}

func (s *Stack) String() string {
	elements := make([]string, len(s.inner))
	for i, element := range s.inner {
		elements[i] = fmt.Sprint(element)
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

func (s *Stack) Clone() *Stack {
	inner := make([]interface{}, 0, len(s.inner))
	for _, v := range s.inner {
		inner = append(inner, v)
	}
	return &Stack{inner: inner}
}

func (s *Stack) Push(v interface{}) *Stack {
	clone := s.Clone()
	clone.inner = append(s.inner, v)
	return clone
}

func (s *Stack) Pop() (interface{}, *Stack) {
	if len(s.inner) == 0 {
		panic("stack is empty")
	}

	top := s.inner[len(s.inner)-1]
	clone := s.Clone()
	clone.inner = s.inner[:len(s.inner)-1]
	return top, clone
}

func (s *Stack) Len() int64 {
	return int64(len(s.inner))
}

func (s *Stack) Peek(idx int64) (interface{}, error) {
	if int64(len(s.inner)) < idx+1 {
		return nil, errors.New("stack idx out of bounds")
	}

	return s.inner[idx], nil
}

func (s *Stack) Get(idx int64) (interface{}, error) {
	if int64(len(s.inner)) < idx+1 {
		return nil, errors.New("stack idx out of bounds")
	}

	return s.inner[idx], nil
}
