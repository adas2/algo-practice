package custom

import (
	"fmt"
	"testing"
)

func TestThreadSafeStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

func TestSimpleStack(t *testing.T) {
	s := initStack(10)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
