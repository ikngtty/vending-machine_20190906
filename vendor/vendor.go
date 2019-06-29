package vendor

import (
	"bytes"
	"fmt"
)

type VendingMachine struct {
	products      []Product
	insertedMoney int
}

func New(products []Product) *VendingMachine {
	return &VendingMachine{
		products:      products,
		insertedMoney: 0,
	}
}

func (vm *VendingMachine) Push(button int) (string, PkgError) {
	if button < 0 || button >= len(vm.products) {
		return "", InvalidButtonError{pushed: button}
	}

	targetProduct := vm.products[button]

	if vm.insertedMoney < targetProduct.Price {
		return "", LackingMoneyError{}
	}

	vm.insertedMoney -= targetProduct.Price
	return targetProduct.Name, nil
}

func (vm *VendingMachine) Insert100Yen() {
	vm.insertedMoney += 100
	return
}

func (vm *VendingMachine) ButtonDescription() string {
	buf := bytes.NewBuffer(nil)

	for i, product := range vm.products {
		buf.WriteString(fmt.Sprintf(
			"%d: (%d yen) %s\n",
			i,
			product.Price,
			product.Name,
		))
	}

	return buf.String()
}
