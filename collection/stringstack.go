package collection

type StringStack struct {
	Stack
}

func (s *StringStack) Push(n string) { s.Stack.Push(n) }
func (s *StringStack) Pop() string   { return s.Stack.Pop().(string) }
