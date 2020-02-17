package main

import "fmt"

type PurchaseOrder struct {
	id    string
	value float64
}

// async func to save the PO in a DB
// callback channel
func savePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
	po.id = "ABC"
	// some DB operation here
	// once completed publishes PO in channel
	callback <- po
}

func testCallback() {
	// create PO
	myPO := new(PurchaseOrder)
	myPO.value = 17.23

	ch := make(chan *PurchaseOrder)

	// save it
	go savePO(myPO, ch)

	// blocks till save operation complete
	status := <-ch

	fmt.Printf("%v\n", status)
}
