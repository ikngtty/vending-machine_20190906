package vendor

import "log"

type VendingMachine struct {
	beverage  string
	coinCount int
}

func New(beverage string) *VendingMachine {
	return &VendingMachine{
		beverage:  beverage,
		coinCount: 0,
	}
}

func (vm *VendingMachine) Push() string {
	log.Print("Pushed")

	// if vm.coinCount <= 0 {
	// 	return ""
	// }
	// vm.coinCount -= 1
	return vm.beverage
}

func (vm *VendingMachine) Insert100Yen() {
	log.Printf("Inserted")

	vm.coinCount += 1
	return
}
