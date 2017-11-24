// All you need to work with stackgo stacks
//
// For an example usage visit http://github.com/alediaferia/stackgo
package stackgo

type Stack struct {
	size             int
	currentPage      []interface{}
	pages            [][]interface{}
	offset           int
	capacity         int
	pageSize         int
	currentPageIndex int
}

const s_DefaultAllocPageSize = 4096

// NewStack Creates a new Stack object with
// an underlying default block allocation size.
// The current default allocation size is one page.
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
	stack.currentPageIndex = 0

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
	stack.currentPageIndex = 0

	return stack
}

// Push pushes a new element to the stack
func (s *Stack) Push(elem ...interface{}) {
	if elem == nil || len(elem) == 0 {
		return
	}

	// ensures enough pages are ready to be
	// filled in and then fills them from
	// the provided elem array
	if s.size+len(elem) > s.capacity {
		newPages := len(elem) / s.pageSize
		if len(elem)%s.pageSize != 0 {
			newPages++
		}

		// appending new empty pages
		for newPages > 0 {
			page := make([]interface{}, s.pageSize)
			s.pages = append(s.pages, page)
			s.capacity += len(page)
			newPages--
		}
	}

	// now that we have enough pages
	// we can start copying the elements
	// into the stack
	s.size += len(elem)
	for len(elem) > 0 {
		available := len(s.currentPage) - s.offset
		min := len(elem)
		if available < min {
			min = available
		}
		copy(s.currentPage[s.offset:], elem[:min])
		elem = elem[min:]
		s.offset += min
		if len(elem) > 0 {
			// page fully filled; move to the next one
			s.currentPage = s.pages[s.currentPageIndex+1]
			s.currentPageIndex++
			s.offset = 0
		}
	}
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

		s.currentPage, s.pages = s.pages[s.currentPageIndex-1], s.pages[:s.currentPageIndex+1]
		s.capacity -= s.pageSize
		s.currentPageIndex--
	}

	elem = s.currentPage[s.offset]

	return
}

func (s *Stack) Top() (elem interface{}) {
	if s.size == 0 {
		return nil
	}

	off := s.offset - 1
	if off < 0 {
		page := s.pages[len(s.pages)-1]
		elem = page[len(page)-1]
		return
	}
	elem = s.currentPage[off]
	return
}

// The current size of the stack
func (s *Stack) Size() int {
	return s.size
}
