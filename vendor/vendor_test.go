package vendor

import (
	"log"
	"testing"
)

func TestVendingMachine_noCoin(t *testing.T) {
	log.Print("no coin")

	const beverage = "Cola"
	vm := New(beverage)

	got := vm.Push()
	const want = ""

	if got != want {
		t.Errorf("want: %s", want)
		t.Errorf("got : %s", got)
	}
}

func TestVendingMachine_1Coin(t *testing.T) {
	log.Print("1 coin")

	const beverage = "Cola"
	vm := New(beverage)

	gots := make([]string, 3)
	wants := make([]string, 3)

	vm.Insert100Yen()

	gots[0] = vm.Push()
	wants[0] = beverage
	gots[1] = vm.Push()
	wants[1] = ""
	gots[2] = vm.Push()
	wants[2] = ""

	for i := range gots {
		if gots[i] != wants[i] {
			t.Errorf("wants[%d]: %s", i, wants[i])
			t.Errorf("gots [%d]: %s", i, gots[i])
		}
	}
}

func TestVendingMachine_nCoins(t *testing.T) {
	log.Print("n coins")

	const beverage = "Cola"
	vm := New(beverage)

	gots := make([]string, 14)
	wants := make([]string, 14)

	vm.Insert100Yen()
	vm.Insert100Yen()
	vm.Insert100Yen()

	gots[0] = vm.Push()
	wants[0] = beverage
	gots[1] = vm.Push()
	wants[1] = beverage

	vm.Insert100Yen()
	vm.Insert100Yen()
	vm.Insert100Yen()
	vm.Insert100Yen()
	vm.Insert100Yen()

	gots[2] = vm.Push()
	wants[2] = beverage
	gots[3] = vm.Push()
	wants[3] = beverage
	gots[4] = vm.Push()
	wants[4] = beverage
	gots[5] = vm.Push()
	wants[5] = beverage
	gots[6] = vm.Push()
	wants[6] = beverage
	gots[7] = vm.Push()
	wants[7] = beverage
	gots[8] = vm.Push()
	wants[8] = ""
	gots[9] = vm.Push()
	wants[9] = ""

	vm.Insert100Yen()
	vm.Insert100Yen()

	gots[10] = vm.Push()
	wants[10] = beverage
	gots[11] = vm.Push()
	wants[11] = beverage
	gots[12] = vm.Push()
	wants[12] = ""
	gots[13] = vm.Push()
	wants[13] = ""

	for i := range gots {
		if gots[i] != wants[i] {
			t.Errorf("wants[%d]: %s", i, wants[i])
			t.Errorf("gots [%d]: %s", i, gots[i])
		}
	}
}
