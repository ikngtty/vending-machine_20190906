package vendor

import (
	"bytes"
	"errors"
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

func (vm *VendingMachine) Push(button int) (string, error) {
	if button < 0 || button >= len(vm.products) {
		return "", fmt.Errorf("given button does not exist: %d", button)
	}

	if vm.coinCount <= 0 {
		return "", errors.New("need more money")
	}

	vm.coinCount -= 1
	return vm.products[button], nil
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
