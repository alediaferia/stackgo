package stackgo

import (
	"testing"
)

func Test_Stackgo(t *testing.T) {
	s := NewStack()
	s.Push(75)
	s.Push(124)
	s.Push("Hello")

	if size := s.Size(); size != 3 {
		t.Fatalf("Unexpected stack size: expected 3, got %d", size)
	}
	if s.Pop().(string) != "Hello" {
		t.Fatalf("Unexpected popped value.")
	}
	if size := s.Size(); size != 2 {
		t.Fatalf("Unexpected stack size: expected 2, got %d", size)
	}
	if s.Pop().(int) != 124 {
		t.Fatalf("Unexpected popped value.")
	}
	if size := s.Size(); size != 1 {
		t.Fatalf("Unexpected stack size: expected 3, got %d", size)
	}
	if s.Pop().(int) != 75 {
		t.Fatalf("Unexpected popped value.")
	}
	if size := s.Size(); size != 0 {
		t.Fatalf("Unexpected stack size: expected 3, got %d", size)
	}
}

func Test_NewStackWithCapacity(t *testing.T) {
	s := NewStackWithCapacity(4)
	if c := s.pageSize; c != 4 {
		t.Fatalf("Unexpected stack block size: got %d, expected 42", c)
	}

	for i := 0; i < 15; i++ {
        s.Push(i)
	}

	if v := s.Top().(int); v != 14 {
		t.Fatalf("Unexpected stack block size: got %d, expected 14", v)
	}

	if v := len(s.pages); v != 4 {
		t.Fatalf("Unexpected number of allocated pages: got %d, expected 4", v)
	}

	if v := s.offset; v != 3 {
		t.Fatalf("Unexpected stack page offset: got %d, expected 3", v)
	}

	if s.Size() != 15 {
		t.Fatalf("Unexpected stack size after 14 insertions: %d", s.size)
	}

	for i := 14; i >= 0; i-- {
		if v := s.Pop(); v != i {
			t.Fatalf("Unexpected popped value: got %d, expected %d", v, i)
		}
	}
}

func Test_EmptyStack(t *testing.T) {
	stack := NewStack()

	if stack.Pop() != nil {
		t.Fatalf("Unexpected pop result with empty stack")
	}
}

func Test_Top(t *testing.T) {
	stack := NewStack()

	if stack.Top() != nil {
		t.Fatalf("Unexpected Top value")
	}

	stack.Push(51)
	if stack.Top() != 51 {
		t.Fatalf("Unexpected Top value")
	}
}

func Test_GithubExample(t *testing.T) {
	stack := NewStack()

	// Stack supports any type
	// so we just push whatever
	// we want here
	stack.Push(75)
	stack.Push(124)
	stack.Push("Hello World!")

	for stack.Size() > 0 {
		_ = stack.Pop()
	}

	if stack.Size() != 0 {
		t.Fatalf("Unexpected stack length != 0")
	}
}

/*
 * Here follows an alternative stack
 * implementation which is not based on
 * slices that we will use as comparison
 * for our benchmarks
 */

type AltStack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *AltStack) Size() int {
	return s.size
}

func (s *AltStack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *AltStack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func Benchmark_PushStackgo(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	b.Logf("Testing push speed of %d integer values on the default Stack implementation", b.N)
	s := NewStack()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func Benchmark_PushStandardStack(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	b.Logf("Testing push speed of %d integer values on the alternate Stack implementation", b.N)
	s := new(AltStack)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func Benchmark_PopStackgo(b *testing.B) {
	b.ReportAllocs()
	b.Logf("Testing pop speed of %d integer values on the default Stack implementation", b.N)
	s := NewStack()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.ResetTimer()
	for s.Size() > 0 {
		_ = s.Pop()
	}
}

func Benchmark_PopStandardStack(b *testing.B) {
	b.ReportAllocs()
	b.Logf("Testing pop speed of %d integer values on the alternate Stack implementation", b.N)
	s := new(AltStack)
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for s.Size() > 0 {
		_ = s.Pop()
	}
}
