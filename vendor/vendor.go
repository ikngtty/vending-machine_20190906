package vendor

import (
	"bytes"
	"fmt"
)

type VendingMachine struct {
	products  []string
	coinCount int
}

func New(products []string) *VendingMachine {
	return &VendingMachine{
		products:  products,
		coinCount: 0,
	}
}

func (vm *VendingMachine) Push() string {
	if vm.coinCount <= 0 {
		return ""
	}
	vm.coinCount -= 1
	return vm.products[0]
}

func (vm *VendingMachine) Insert100Yen() {
	vm.coinCount += 1
	return
}

func (vm *VendingMachine) ButtonDescription() string {
	buf := bytes.NewBuffer(nil)

	for i, product := range vm.products {
		buf.WriteString(fmt.Sprintf("%d: %s\n", i, product))
	}

	return buf.String()
}
