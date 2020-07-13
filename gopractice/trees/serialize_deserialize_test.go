package trees

import (
	"fmt"
	"testing"
)

func TestBTSerDeser(t *testing.T) {
	obj := Constructor()

	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	data := obj.serialize(root)
	fmt.Printf("serailized data: %s\n", data)
	inorderT(root)
	ans := obj.deserialize(data)
	fmt.Printf("de-serailized tree:\n")
	inorderT(ans)
}
