package runtime

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
	return fmt.Sprintf("Stack{inner: [%s]}", strings.Join(elements, ", "))
}

func (s *Stack) Push(v interface{}) {
	s.inner = append(s.inner, v)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.inner) == 0 {
		return nil, errors.New("stack is empty")
	}

	top := s.inner[len(s.inner)-1]
	s.inner = s.inner[:len(s.inner)-1]
	return top, nil
}
