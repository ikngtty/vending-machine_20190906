package vendor

import "testing"

func TestVendingMachine_Push(t *testing.T) {
	vm := New()
	beverage := vm.Push()
	const want = "Cola"
	if beverage != want {
		t.Errorf("want: %s", want)
		t.Errorf("got : %s", beverage)
	}
}
