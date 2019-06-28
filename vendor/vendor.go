package vendor

type VendingMachine struct{}

func New() *VendingMachine {
	return &VendingMachine{}
}

func (vm *VendingMachine) Push() string {
	return "Cola"
}
