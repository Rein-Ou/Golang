package stack

import "errors"

type IStack interface {
	Len() int
	IsEmpty() bool
	Push(x interface{})
	Pop() (x interface{}, err error)
}

type Stack []interface{}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("Stack is empty")
	}
	return s[len(s)-1], nil
}

func (s *Stack) Pop() (x interface{}, err error) {
	stk := *s
	if len(stk) == 0 {
		return nil, errors.New("Stack is empty")
	}
	top := stk[len(stk)-1]
	*s = stk[:len(stk)-1]
	return top, nil
}
