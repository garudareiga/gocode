package collection

type Stack struct {
	data []interface{}
}

// Push add x to the top of the stack.
func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack.
// It's a run-time error to call Pop on an empty stack.
func (s *Stack) Pop() interface{} {
	if s.Size() > 0 {
		i := len(s.data) - 1
		res := s.data[i]
		s.data[i] = nil // to avoid memory leak
		s.data = s.data[:i]
		return res
	}
	return nil
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.data)
}

// Peek returns the top element of the stack
func (s *Stack) Peek() interface{} {
	if s.Size() > 0 {
		return s.data[s.Size()-1]
	}
	return nil
}
