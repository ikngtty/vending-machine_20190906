package vendor

import "testing"

func TestVendingMachine_Push(t *testing.T) {
	const want = "Cola"
	vm := New(want)
	got := vm.Push()
	if got != want {
		t.Errorf("want: %s", want)
		t.Errorf("got : %s", got)
	}
}
