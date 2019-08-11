package vendor

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

func (vm *VendingMachine) ButtonsDescription() []ButtonDescription {
	descriptions := make([]ButtonDescription, len(vm.products))

	for i, product := range vm.products {
		descriptions[i] = ButtonDescription{Button: i, Product: product}
	}

	return descriptions
}
