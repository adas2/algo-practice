package practice

import (
	"fmt"
	"testing"
)

func TestMaxStack(t *testing.T) {
	st := GetStackHandle(4)
	st.Push(10)
	fmt.Println("max:", st.FindMax())
	st.Push(3)
	fmt.Println("max:", st.FindMax())
	st.Push(5)
	fmt.Println("max:", st.FindMax())
	if err := st.Push(1); err != nil {
		t.Fatalf("PushError")
	}
	fmt.Println("max:", st.FindMax())
	// fmt.Println(st.Peek())
	if err := st.Pop(); err != nil {
		t.Fatalf("Pop error")
	}
	fmt.Println("max:", st.FindMax())
	// fmt.Println(st.Peek())
	if err := st.Pop(); err != nil {
		t.Fatalf("Pop error")
	}
	fmt.Println("max:", st.FindMax())
	// fmt.Println(st.Peek())
	if err := st.Pop(); err != nil {
		t.Fatalf("Pop error")
	}
	fmt.Println("max:", st.FindMax())
	fmt.Println("peek:", st.Peek(), "empty?", st.IsEmpty())
	if err := st.Pop(); err != nil {
		t.Fatalf("Pop error")
	}
	fmt.Println("max:", st.FindMax())
	fmt.Println("peek:", st.Peek(), "empty?", st.IsEmpty())
	if err := st.Pop(); err != nil {
		t.Fatalf("Pop error")
	}
	// st.Pop()
}
