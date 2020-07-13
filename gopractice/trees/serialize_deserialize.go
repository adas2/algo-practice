package trees

import (
	"fmt"
	"strconv"
	"strings"
)

// LC: 297

// Codec type
type Codec struct {
	vals []string
}

// Constructor for codec
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	// DFS pre-order traversal
	if root == nil {
		return "Nil "
	}
	// pre-order
	str := strconv.Itoa(root.Val) + " "
	str += c.serialize(root.Left)
	str += c.serialize(root.Right)
	return str
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	// create tree based on pre-order stack
	if len(data) == 0 {
		return nil
	}
	// tokenize the data
	c.vals = c.tokenize(data)
	// use recusive function
	root := c.deserializeUtil()

	return root
}

func (c *Codec) deserializeUtil() *TreeNode {

	if c.vals[0] == "Nil" {
		// pop and return nil
		c.vals = c.vals[1:]
		return nil
	}
	// else if non-Nil value
	// take the first value and create a node
	nodeVal, _ := strconv.Atoi(c.vals[0])
	tNode := &TreeNode{nodeVal, nil, nil}
	// pop the first value
	c.vals = c.vals[1:]

	tNode.Left = c.deserializeUtil()
	tNode.Right = c.deserializeUtil()

	return tNode
}

func (c *Codec) tokenize(data string) []string {
	nums := strings.Fields(data)
	return nums
}

// TreeNode is a bt struct
/*** Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* utils */
func inorderT(root *TreeNode) {
	if root != nil {
		inorderT(root.Left)
		fmt.Println(root.Val)
		inorderT(root.Right)
	}
}
