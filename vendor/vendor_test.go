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

	noMoneyError := errors.New("need more money")
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
						{0, "", noMoneyError},
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
						{0, "", noMoneyError},
						{0, "", noMoneyError},
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
						{0, "", noMoneyError},
						{0, "", noMoneyError},
					},
				},
				{
					insertCount: 2,
					pushes: []pushing{
						{0, cola, nil},
						{0, cola, nil},
						{0, "", noMoneyError},
						{0, "", noMoneyError},
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
						{1, "", noMoneyError},
						{2, "", noMoneyError},
						{3, "", noButtonError},
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			vm := New(products)

			for buyTime, ope := range tc.operations {
				for i := 0; i < ope.insertCount; i++ {
					vm.Insert100Yen()
				}

				for pushTime, p := range ope.pushes {
					beverage, err := vm.Push(p.button)
					t.Logf("buying %d pushing %d", buyTime, pushTime)
					if beverage != p.wantBeverage {
						t.Errorf("want beverage: %s", p.wantBeverage)
						t.Errorf("got  beverage: %s", beverage)
					}
					if (p.wantErr == nil && err != nil) ||
						(p.wantErr != nil && (err == nil || err.Error() != p.wantErr.Error())) {
						t.Errorf("want error: %v", p.wantErr)
						t.Errorf("got  error: %v", err)
					}
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
