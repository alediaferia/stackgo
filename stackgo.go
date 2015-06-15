// All you need to work with stackgo stacks
//
// For an example usage visit http://github.com/alediaferia/stackgo
package stackgo

type Stack struct {
	slice []interface{}
	blockSize int
}

const s_DefaultAllocBlockSize = 20;

// NewStack Creates a new Stack object with
// an underlying default block allocation size.
// The default is currently 20 but this may vary
// in the future.
// If you want to use a different block size use
//  NewStackWithCapacity()
func NewStack() (*Stack) {
	stack := new(Stack)
	stack.slice = make([]interface{}, 0, s_DefaultAllocBlockSize)
	stack.blockSize = s_DefaultAllocBlockSize

	return stack
}

// NewStackWithCapacity makes it easy to specify
// a custom block size for inner slice backing the
// stack
func NewStackWithCapacity(cap int) (*Stack) {
	stack := new(Stack)
	stack.slice = make([]interface{}, 0, cap)
	stack.blockSize = cap

	return stack
}

// Push pushes a new element to the stack
func (s *Stack) Push(elem interface{}) {

	if len(s.slice) >= cap(s.slice) {
		slice := make([]interface{}, 0, len(s.slice) + s.blockSize)
		copy(slice, s.slice)
		s.slice = slice
	}

	s.slice = append(s.slice, elem)
}

// Pop pops the top element from the stack
// If the stack is empty it returns nil
func (s *Stack) Pop() (elem interface{}) {
	if s.Size() == 0 {
		return nil
	}

	elem, s.slice = s.slice[len(s.slice) - 1], s.slice[:len(s.slice) - 1]

	return
}

// The current size of the stack
func (s *Stack) Size() int {
	return len(s.slice)
}
