package practice

import (
	"fmt"
	"testing"
)

func TestInvalidTx(t *testing.T) {
	input := []string{"bob,627,1973,amsterdam", "alex,387,885,bangkok", "alex,355,1029,barcelona", "alex,587,402,bangkok", "chalicefy,973,830,barcelona", "alex,932,86,bangkok", "bob,188,989,amsterdam"}
	fmt.Println(invalidTransactions(input))
}
