package stackgo

import (
	"testing"
)

func Test_Stack(t *testing.T) {
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

/*
 * Here follows an alternative stack
 * implementation which is not based on
 * slices that we will use as comparison
 * for our benchmarks
 */

type AltStack struct {
	top *Element
	size int
}

type Element struct {
	value interface{}
	next *Element
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

func Benchmark_PushDefaultStack(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	b.Logf("Testing push speed of %d integer values on the default Stack implementation", b.N)
	s := NewStack()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func Benchmark_PushAltStack(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()
	b.Logf("Testing push speed of %d integer values on the alternate Stack implementation", b.N)
	s := new(AltStack)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

func Benchmark_PopDefaultStack(b *testing.B) {
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

func Benchmark_PopAltStack(b *testing.B) {
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
