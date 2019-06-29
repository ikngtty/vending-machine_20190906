package vendor

import (
	"fmt"
	"testing"
)

func TestVendingMachine_use100Yen(t *testing.T) {
	const cola = "Cola"
	products := []string{cola}

	type operation struct {
		insertCount int
		wants       []string
	}
	testcases := []struct {
		name       string
		operations []operation
	}{
		{
			name: "no coin",
			operations: []operation{
				{
					insertCount: 0,
					wants:       []string{""},
				},
			},
		},
		{
			name: "1 coin",
			operations: []operation{
				{
					insertCount: 1,
					wants:       []string{cola, "", ""},
				},
			},
		},
		{
			name: "n coins",
			operations: []operation{
				{
					insertCount: 3,
					wants: []string{
						cola, cola,
					},
				},
				{
					insertCount: 5,
					wants: []string{
						cola, cola, cola, cola,
						cola, cola, "", "",
					},
				},
				{
					insertCount: 2,
					wants: []string{
						cola, cola, "", "",
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			vm := New(products)
			pushCount := 0

			for _, ope := range tc.operations {
				for i := 0; i < ope.insertCount; i++ {
					vm.Insert100Yen()
				}

				for _, want := range ope.wants {
					got := vm.Push()
					if got != want {
						t.Errorf("wants[%d]: %s", pushCount, want)
						t.Errorf("gots [%d]: %s", pushCount, got)
					}

					pushCount++
				}
			}
		})
	}
}

func TestVendingMachine_ButtonDescription(t *testing.T) {
	const cola = "Cola"
	const oolong = "Oolong Tea"
	const drPepper = "Dr.Pepper"
	products := []string{cola, oolong, drPepper}
	vm := New(products)

	want :=
		fmt.Sprintf("0: %s", cola) + "\n" +
			fmt.Sprintf("1: %s", oolong) + "\n" +
			fmt.Sprintf("2: %s", drPepper) + "\n"
	got := vm.ButtonDescription()

	if got != want {
		t.Errorf("want:\n%s", want)
		t.Errorf("got :\n%s", got)
	}
}

func TestVendingMachine_ButtonDescription_emptyProducts(t *testing.T) {
	products := []string{}
	vm := New(products)

	want := ""
	got := vm.ButtonDescription()

	if got != want {
		t.Errorf("want:\n%s", want)
		t.Errorf("got :\n%s", got)
	}
}
