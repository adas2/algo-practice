package custom

import (
	"fmt"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	// ["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]
	// [[],[1],[2],[2],[],[1],[2],[]]

	// ["RandomizedSet","insert","insert","remove","insert","remove","getRandom"]
	// [[],[0],[1],[0],[2],[1],[]]
	obj := InitRandomizedSet()
	param_1 := obj.Insert(0)
	fmt.Println(param_1)

	param_3 := obj.Insert(1)
	fmt.Println(param_3)
	// param_4 := obj.GetRandom()
	// fmt.Println(param_4)
	param_5 := obj.Remove(0)
	fmt.Println(param_5)
	param_6 := obj.Insert(2)
	fmt.Println(param_6)
	param_2 := obj.Remove(1)
	fmt.Println(param_2)
	param_7 := obj.GetRandom()
	fmt.Println(param_7)

	// fmt.Println(param_2, param_3, param_5, param_6, param_7)
}
