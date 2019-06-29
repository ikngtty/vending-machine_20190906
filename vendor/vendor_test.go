package vendor

import (
	"errors"
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
					got, err := vm.Push(0)
					if got != want {
						t.Errorf("wants [%d]: %s", pushCount, want)
						t.Errorf("gots  [%d]: %s", pushCount, got)
					}
					if err != nil {
						t.Errorf("errors[%d]: %s", pushCount, err)
					}

					pushCount++
				}
			}
		})
	}
}

func TestVendingMachine_use100YenForVariousBeverages(t *testing.T) {
	const cola = "Cola"
	const oolongTea = "Oolong Tea"
	const drPepper = "Dr.Pepper"
	products := []string{cola, oolongTea, drPepper}

	noButtonError := errors.New("given button does not exist: 3")

	type operation struct {
		insertCount   int
		pushes        []int
		wantBeverages []string
		wantErrors    []error
	}
	testcases := []struct {
		name       string
		operations []operation
	}{
		{
			name: "various beverages",
			operations: []operation{
				{
					insertCount: 4,
					pushes: []int{
						1, 0, 2, 3,
						2, 1, 2, 3,
					},
					wantBeverages: []string{
						oolongTea, cola, drPepper, "",
						drPepper, "", "", "",
					},
					wantErrors: []error{
						nil, nil, nil, noButtonError,
						nil, nil, nil, noButtonError,
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

				for i, button := range ope.pushes {
					wantBeverage := ope.wantBeverages[i]
					wantErr := ope.wantErrors[i]
					beverage, err := vm.Push(button)
					if beverage != wantBeverage {
						t.Errorf("want beverage[%d]: %s", pushCount, wantBeverage)
						t.Errorf("got  beverage[%d]: %s", pushCount, beverage)
					}
					if (wantErr == nil && err != nil) ||
						(wantErr != nil && (err == nil || err.Error() != wantErr.Error())) {
						t.Errorf("want error[%d]: %v", pushCount, wantErr)
						t.Errorf("got  error[%d]: %v", pushCount, err)
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

	testcases := []struct {
		name     string
		products []string
		want     string
	}{
		{
			name:     "no product",
			products: []string{},
			want:     "",
		},
		{
			name:     "various products",
			products: []string{cola, oolong, drPepper},
			want: fmt.Sprintf("0: %s", cola) + "\n" +
				fmt.Sprintf("1: %s", oolong) + "\n" +
				fmt.Sprintf("2: %s", drPepper) + "\n",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			vm := New(tc.products)

			got := vm.ButtonDescription()
			if got != tc.want {
				t.Errorf("want:\n%s", tc.want)
				t.Errorf("got :\n%s", got)
			}
		})
	}
}

func TestVendingMachine_buyVariousBeverages(t *testing.T) {
	const cola = "Cola"
	const oolongTea = "Oolong Tea"
	const drPepper = "Dr.Pepper"
	products := []string{cola, oolongTea, drPepper}
	vm := New(products)
	for i := 0; i < 5; i++ {
		vm.Insert100Yen()
	}

	{
		const button = 0
		const wantBeverage = cola

		beverage, err := vm.Push(button)
		if beverage != wantBeverage {
			t.Errorf("want beverage: %s", wantBeverage)
			t.Errorf("got  beverage: %s", beverage)
		}
		if err != nil {
			t.Errorf("want err: %v", nil)
			t.Errorf("got  err: %v", err)
		}
	}
	{
		const button = 1
		const wantBeverage = oolongTea

		beverage, err := vm.Push(button)
		if beverage != wantBeverage {
			t.Errorf("want beverage: %s", wantBeverage)
			t.Errorf("got  beverage: %s", beverage)
		}
		if err != nil {
			t.Errorf("want err: %v", nil)
			t.Errorf("got  err: %v", err)
		}
	}
	{
		const button = 2
		const wantBeverage = drPepper

		beverage, err := vm.Push(button)
		if beverage != wantBeverage {
			t.Errorf("want beverage: %s", wantBeverage)
			t.Errorf("got  beverage: %s", beverage)
		}
		if err != nil {
			t.Errorf("want err: %v", nil)
			t.Errorf("got  err: %v", err)
		}
	}
	{
		const button = 3
		const wantBeverage = ""
		const wantError = "given button does not exist: 3"

		beverage, err := vm.Push(button)
		if beverage != wantBeverage {
			t.Errorf("want beverage: %s", wantBeverage)
			t.Errorf("got  beverage: %s", beverage)
		}
		if err == nil || err.Error() != wantError {
			t.Errorf("want err: %s", wantError)
			t.Errorf("got  err: %v", err)
		}
	}
}
