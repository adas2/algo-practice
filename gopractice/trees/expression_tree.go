package trees

// Exp tree binary tree with leaf nodes as operands and non-leaf nodes as operator
// given 2 exp tree, find if they are equivalent
//  lc # 1612

type extreeNode struct {
	chr   string
	left  *extreeNode
	right *extreeNode
}

func checkEquivalence(root1 *extreeNode, root2 *extreeNode) bool {

	// maintain a map of frequency of operands seen
	oMap := map[string]int{}

	// terminal case
	if root1 == nil && root2 == nil {
		return true
	}

	// populate map from root1
	checkUtil1(root1, oMap)

	// check equivalence
	checkUtil2(root2, oMap)

	// check the map for all entries to have zero count
	for _, val := range oMap {

		// fmt.Println(key, val)

		if val != 0 {
			return false
		}
	}
	return true
}

// util to fill up tje map
func checkUtil1(root *extreeNode, oMap map[string]int) {

	if root == nil {
		return
	}

	// if leaf store the oMap
	if root.left == nil && root.right == nil {
		if _, exists := oMap[root.chr]; !exists {
			oMap[root.chr] = 1
		} else {
			oMap[root.chr]++
		}
	}

	checkUtil1(root.left, oMap)
	checkUtil1(root.right, oMap)

}

func checkUtil2(root *extreeNode, oMap map[string]int) {
	if root == nil {
		return
	}

	// if leaf store the oMap
	if root.left == nil && root.right == nil {
		if _, exists := oMap[root.chr]; !exists {
			// this case means not equivalent, but will be captured during map traversal
			oMap[root.chr] = 1
		} else {
			// exists decrement
			oMap[root.chr]--
		}
	}

	checkUtil2(root.left, oMap)
	checkUtil2(root.right, oMap)
}

// logic:
// (a+b) + c == a + (b+c)
// in this case we just check if the operators are '+' and
// if the operands are same with same frequency of occurence
// pre-order traversal based checks
// T=O(N), S=O(26); for O(1) space hash tree1 and tree2, compare the hash values
