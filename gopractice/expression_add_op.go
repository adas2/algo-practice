package practice

import (
	"fmt"
)

// LC # 282
// WIP

func addOperators(num string, target int) []string {

	// candidate := string(r[len(r)-1])
	expList := []string{}
	if len(num) != 0 {
		genExp2(num, 0, "", target, &expList)
	}

	return expList
}

// simple recursive version
func genExp2(nums string, k int, candidate string, target int, expList *[]string) {
	// fmt.Printf("iter: %d, candidate: %s\n", k, candidate)
	if k == len(nums) || len(nums) == 0 {
		// complete
		// fmt.Println(candidate)
		if evalExp(candidate) == target {
			// fmt.Printf("expression  value match: %s\n", candidate)
			*expList = append(*expList, candidate)
		}
		return
	}

	// if first num, just use num without operators
	if k == 0 {
		genExp2(nums, k+1, candidate+string(nums[0]), target, expList)
		return
	}
	ops := []byte{'*', '+', '-', 'c'}

	// for nums[1]...nums[n-1] use all operators
	for _, op := range ops {
		partial := ""
		// if concat or start of str just add the num
		if op == 'c' {
			// skip case like '0x', '00x'
			if candidate[len(candidate)-1] == '0' {
				continue
			}
			partial = string(nums[k])
		} else { //or else add sym+num
			partial = string(op) + string(nums[k])
		}
		genExp2(nums, k+1, candidate+partial, target, expList)
	}
}

func evalExp(exp string) int {
	// declare vals and ops stack
	vals := []int{}
	ops := []byte{}
	// prec for * + -
	prec := map[byte]int{'*': 1, '+': 0, '-': 0}
	i := 0
	for i < len(exp) {
		// if numeric val
		if isNumeric(exp[i]) {
			val := 0
			// while more numeric values are there
			for i < len(exp) && isNumeric(exp[i]) {
				val = 10*val + int(exp[i]-'0')
				i++
			}
			vals = append(vals, val)
		} else {
			// symbol
			// make sure higher prec is calculated first
			// fmt.Printf("ops: %v, vals: %v\n", ops, vals)
			for len(ops) > 0 && prec[ops[len(ops)-1]] >= prec[exp[i]] {
				// pop vals and apply ops
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				n := len(vals)
				arg2 := vals[n-1]
				arg1 := vals[n-2]
				vals = vals[:n-2]
				// calc
				val := applyOp(arg1, arg2, op)
				vals = append(vals, val)
			}
			ops = append(ops, exp[i])
			i++
		}
	}
	// fparse remai remaining ops
	for len(ops) > 0 {
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		n := len(vals)
		arg2 := vals[n-1]
		arg1 := vals[n-2]
		vals = vals[:n-2]

		// calc
		val := applyOp(arg1, arg2, op)
		vals = append(vals, val)
	}

	// lst val
	return vals[0]
}

func applyOp(a, b int, sym byte) int {
	if sym == '*' {
		return a * b
	} else if sym == '+' {
		return a + b
	} else if sym == '-' {
		return a - b
	}
	fmt.Printf("Unknown sym %v\n", sym)
	return 0
}

func isNumeric(b byte) bool {
	if b-'0' >= 0 && b-'0' <= 9 {
		return true
	}
	return false
}

// Logic:
// use recursive backtracking method to compose
// new combinations of expressions
//

// ============
