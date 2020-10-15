package util

import "fmt"

type Scope struct {
	Vars  *Vars
	Stack *Stack
}

func NewScope() *Scope {
	return &Scope{
		Vars:  NewVars(),
		Stack: NewStack(),
	}
}

func (s *Scope) String() string {
	return fmt.Sprintf("Scope{Stack: %s, Vars: %s}", s.Stack, s.Vars)
}

func (s *Scope) Clone() *Scope {
	vars := s.Vars.Clone()
	stack := s.Stack.Clone()
	return &Scope{Vars: vars, Stack: stack}
}

func (s *Scope) GetVar(key string) interface{} {
	return s.Vars.Get(key)
}

func (s *Scope) SetVar(key string, value interface{}) *Scope {
	vars := s.Vars.Set(key, value)
	stack := s.Stack.Clone()
	return &Scope{Vars: vars, Stack: stack}
}

func (s *Scope) PopStack() (interface{}, *Scope) {
	vars := s.Vars.Clone()
	value, stack := s.Stack.Pop()
	return value, &Scope{Vars: vars, Stack: stack}
}

func (s *Scope) PushStack(v interface{}) *Scope {
	vars := s.Vars.Clone()
	stack := s.Stack.Push(v)
	return &Scope{Vars: vars, Stack: stack}
}
