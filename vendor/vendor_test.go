package vendor

import (
	"errors"
	"fmt"
	"testing"
)

const cola = "Cola"
const oolongTea = "Oolong Tea"
const drPepper = "Dr.Pepper"

func TestVendingMachine_use100Yen(t *testing.T) {
	products := []string{cola, oolongTea, drPepper}

	noButtonError := errors.New("given button does not exist: 3")

	type pushing struct {
		button       int
		wantBeverage string
		wantErr      error
	}
	type buying struct {
		insertCount int
		pushes      []pushing
	}
	testcases := []struct {
		name       string
		operations []buying
	}{
		{
			name: "no coin",
			operations: []buying{
				{
					insertCount: 0,
					pushes: []pushing{
						{0, "", nil},
					},
				},
			},
		},
		{
			name: "1 coin",
			operations: []buying{
				{
					insertCount: 1,
					pushes: []pushing{
						{0, cola, nil},
						{0, "", nil},
						{0, "", nil},
					},
				},
			},
		},
		{
			name: "n coins",
			operations: []buying{
				{
					insertCount: 3,
					pushes: []pushing{
						{0, cola, nil},
						{0, cola, nil},
					},
				},
				{
					insertCount: 5,
					pushes: []pushing{
						{0, cola, nil},
						{0, cola, nil},
						{0, cola, nil},
						{0, cola, nil},
						{0, cola, nil},
						{0, cola, nil},
						{0, "", nil},
						{0, "", nil},
					},
				},
				{
					insertCount: 2,
					pushes: []pushing{
						{0, cola, nil},
						{0, cola, nil},
						{0, "", nil},
						{0, "", nil},
					},
				},
			},
		},
		{
			name: "various beverages",
			operations: []buying{
				{
					insertCount: 4,
					pushes: []pushing{
						{1, oolongTea, nil},
						{0, cola, nil},
						{2, drPepper, nil},
						{3, "", noButtonError},
						{2, drPepper, nil},
						{1, "", nil},
						{2, "", nil},
						{3, "", noButtonError},
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

				for _, p := range ope.pushes {
					beverage, err := vm.Push(p.button)
					if beverage != p.wantBeverage {
						t.Errorf("want beverage[%d]: %s", pushCount, p.wantBeverage)
						t.Errorf("got  beverage[%d]: %s", pushCount, beverage)
					}
					if (p.wantErr == nil && err != nil) ||
						(p.wantErr != nil && (err == nil || err.Error() != p.wantErr.Error())) {
						t.Errorf("want error[%d]: %v", pushCount, p.wantErr)
						t.Errorf("got  error[%d]: %v", pushCount, err)
					}

					pushCount++
				}
			}
		})
	}
}

func TestVendingMachine_ButtonDescription(t *testing.T) {
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
			products: []string{cola, oolongTea, drPepper},
			want: fmt.Sprintf("0: %s", cola) + "\n" +
				fmt.Sprintf("1: %s", oolongTea) + "\n" +
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
