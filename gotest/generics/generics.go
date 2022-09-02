package generics

//
// type Stack struct {
// values []interface{}
// }
// type StackOfInts = Stack
// type StackOfStrings = Stack
//
// func (s *Stack) Push(value interface{}) {
// s.values = append(s.values, value)
// }
//
// func (s *Stack) IsEmpty() bool {
// return len(s.values) == 0
// }
//
// func (s *Stack) Pop() (interface{}, bool) {
// if s.IsEmpty() {
// var zero interface{}
// return zero, false
// }
// index := len(s.values) - 1
// el := s.values[index]
// s.values = s.values[:index]
// return el, true
// }
//
// /*V-2*/

// Declaring a generic-type stack, constrained to accept any concrete type
type Stack[T any] struct {
	values []T //slice of any,
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T // zero of an interface ; since it will intialized to the sero of its type-(generic parameter{0 for int, "" for string etc})
		return zero, false
	}

	index := len(s.values) - 1 // since slice is zero indexed
	el := s.values[index]
	s.values = s.values[:index] //since this sliceing exclude the last element
	return el, true             //el's type is generics of type T{int, string or any that passed}
}
