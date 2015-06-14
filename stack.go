package stackgo

type Stack struct {
	slice []interface{}
}

func NewStack() (*Stack) {
	stack := new(Stack)
	stack.slice = make([]interface{}, 0, 0)

	return stack
}

func NewStackWithCapacity(cap int) (*Stack) {
	stack := new(Stack)
	stack.slice = make([]interface{}, 0, cap)

	return stack
}

func (s *Stack) Push(elem interface{}) {
	s.slice = append(s.slice, elem)
}

func (s *Stack) Pop() (elem interface{}) {
	if s.Size() == 0 {
		return nil
	}

	elem, s.slice = s.slice[len(s.slice) - 1], s.slice[:len(s.slice) - 1]

	return
}

func (s *Stack) Size() int {
	return len(s.slice)
}
