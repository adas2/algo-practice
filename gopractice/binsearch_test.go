package practice

import (
	"testing"
	"fmt"
)
func TestBinarySearchVanilla(t *testing.T){
	arr := []int{2,3,4,5,6,7}
	fmt.Println("Array:", arr)
	target := 6
	index := BinarySearchVanilla(arr, target, 0, 5)
	fmt.Println(target, index)
	target = 11
	index = BinarySearchVanilla(arr, target, 0, 5)
	fmt.Println(target, index)
}

func TestBinarySearchClosest(t *testing.T){
	arr := []int{2,3,4,7,15,23}
	fmt.Println("Array:", arr)
	target := 4
	index := BinarySearchClosest(arr, target, 0, 5)
	fmt.Println(target, index)
	target = 28
	index = BinarySearchClosest(arr, target, 0, 5)
	fmt.Println(target, index)
}

func TestBinarySearchLast(t *testing.T){
	arr := []int{2,4,4,4,6,7}
	fmt.Println("Array:", arr)
	target := 4
	index := BinarySearchLast(arr, target, 0, 5)
	fmt.Println(target, index)
	target = 15
	index = BinarySearchLast(arr, target, 0, 5)
	fmt.Println(target, index)
}

