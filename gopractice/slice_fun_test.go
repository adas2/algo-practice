package practice

import "testing"

func TestSliceFun(t *testing.T) {

	S := createSlice(3)
	S.printSlice()
	S.addLast(10)
	S.addLast(15)
	S.addLast(20)
	S.addFront(5)
	S.printSlice()
	S.deleteLast()
	S.printSlice()
	S.deleteFront()
	S.deleteLast()
	S.deleteLast()
	S.printSlice()
}
