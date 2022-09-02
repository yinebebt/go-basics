package generics

type Stack struct {
	values []interface{}
}
type StackOfInts = Stack
type StackOfStrings = Stack

func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		var zero interface{}
		return zero, false
	}
	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
