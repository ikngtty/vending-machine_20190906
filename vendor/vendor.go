package vendor

type VendingMachine struct {
	beverage string
}

func New(beverage string) *VendingMachine {
	return &VendingMachine{beverage: beverage}
}

func (vm *VendingMachine) Push() string {
	return vm.beverage
}
