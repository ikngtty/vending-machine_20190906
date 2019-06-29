package vendor

import (
	"bytes"
	"fmt"
)

type VendingMachine struct {
	products  []Product
	coinCount int
}

func New(products []Product) *VendingMachine {
	return &VendingMachine{
		products:  products,
		coinCount: 0,
	}
}

func (vm *VendingMachine) Push(button int) (string, PkgError) {
	if button < 0 || button >= len(vm.products) {
		return "", InvalidButtonError{pushed: button}
	}

	if vm.coinCount <= 0 {
		return "", LackingMoneyError{}
	}

	vm.coinCount -= 1
	return vm.products[button].Name, nil
}

func (vm *VendingMachine) Insert100Yen() {
	vm.coinCount += 1
	return
}

func (vm *VendingMachine) ButtonDescription() string {
	buf := bytes.NewBuffer(nil)

	for i, product := range vm.products {
		buf.WriteString(fmt.Sprintf("%d: %s\n", i, product.Name))
	}

	return buf.String()
}
