// All you need to work with stackgo stacks
//
// For an example usage visit http://github.com/alediaferia/stackgo
package stackgo

type Stack struct {
	size int
	currentPage []interface{}
	pages [][]interface{}
	offset int
	capacity int
	pageSize int
}

const s_DefaultAllocPageSize = 4096

// NewStack Creates a new Stack object with
// an underlying default block allocation size.
// The default is currently 20 but this may vary
// in the future.
// If you want to use a different block size use
//  NewStackWithCapacity()
func NewStack() *Stack {
	stack := new(Stack)
	stack.currentPage = make([]interface{}, s_DefaultAllocPageSize)
	stack.pages = [][]interface{}{stack.currentPage}
	stack.offset = 0
	stack.capacity = s_DefaultAllocPageSize
	stack.pageSize = s_DefaultAllocPageSize
	stack.size = 0

	return stack
}


// NewStackWithCapacity makes it easy to specify
// a custom block size for inner slice backing the
// stack
func NewStackWithCapacity(cap int) *Stack {
	stack := new(Stack)
	stack.currentPage = make([]interface{}, cap)
	stack.pages = [][]interface{}{stack.currentPage}
	stack.offset = 0
	stack.capacity = cap
	stack.pageSize = cap
	stack.size = 0

	return stack
}


// Push pushes a new element to the stack
func (s *Stack) Push(elem interface{}) {
    if s.size == s.capacity {
		s.currentPage = make([]interface{}, s.pageSize)
		s.pages = append(s.pages, s.currentPage)
		s.capacity += s.pageSize
		s.offset = 0
	}
	s.currentPage[s.offset] = elem
	s.offset++
	s.size++
}

// Pop pops the top element from the stack
// If the stack is empty it returns nil
func (s *Stack) Pop() (elem interface{}) {
	if s.size == 0 {
		return nil
	}

	s.offset--
	s.size--
	if s.offset < 0 {
		s.offset = s.pageSize - 1

		s.currentPage, s.pages = s.pages[len(s.pages) - 2], s.pages[:len(s.pages) - 1]
		s.capacity -= s.pageSize
	}

	elem = s.currentPage[s.offset]

	return
}

func (s *Stack) Top() (elem interface{}) {
	if s.size == 0 {
		return nil
	}

	if s.offset == 0 {
		page := s.pages[len(s.pages)-1]
		elem = page[len(page)-1]
		return
	}
	elem = s.currentPage[s.offset - 1]
	return
}

// The current size of the stack
func (s *Stack) Size() int {
	return s.size
}
