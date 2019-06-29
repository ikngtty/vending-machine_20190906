package vendor

import (
	"testing"
)

var cola = Product{Name: "Cola", Price: 100}
var oolongTea = Product{Name: "Oolong Tea", Price: 100}
var drPepper = Product{Name: "Dr.Pepper", Price: 100}
var redBull = Product{Name: "Red Bull", Price: 200}

var products = []Product{cola, oolongTea, drPepper, redBull}

const description = `0: (100 yen) Cola
1: (100 yen) Oolong Tea
2: (100 yen) Dr.Pepper
3: (200 yen) Red Bull
`

func TestVendingMachine_use100Yen(t *testing.T) {
	errLackingMoney := LackingMoneyError{}
	errInvalidButton := InvalidButtonError{pushed: 4}

	type pushing struct {
		button       int
		wantBeverage string
		wantErr      PkgError
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
						{0, "", errLackingMoney},
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
						{0, cola.Name, nil},
						{0, "", errLackingMoney},
						{0, "", errLackingMoney},
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
						{0, cola.Name, nil},
						{0, cola.Name, nil},
					},
				},
				{
					insertCount: 5,
					pushes: []pushing{
						{0, cola.Name, nil},
						{0, cola.Name, nil},
						{0, cola.Name, nil},
						{0, cola.Name, nil},
						{0, cola.Name, nil},
						{0, cola.Name, nil},
						{0, "", errLackingMoney},
						{0, "", errLackingMoney},
					},
				},
				{
					insertCount: 2,
					pushes: []pushing{
						{0, cola.Name, nil},
						{0, cola.Name, nil},
						{0, "", errLackingMoney},
						{0, "", errLackingMoney},
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
						{1, oolongTea.Name, nil},
						{0, cola.Name, nil},
						{2, drPepper.Name, nil},
						{4, "", errInvalidButton},
						{2, drPepper.Name, nil},
						{1, "", errLackingMoney},
						{2, "", errLackingMoney},
						{4, "", errInvalidButton},
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
					if !errorsAreSame(err, p.wantErr) {
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
		products []Product
		want     string
	}{
		{
			name:     "no product",
			products: []Product{},
			want:     "",
		},
		{
			name:     "various products",
			products: products,
			want:     description,
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
