package vendor

import (
	"reflect"
	"testing"
)

var cola = Product{Name: "Cola", Price: 100}
var oolongTea = Product{Name: "Oolong Tea", Price: 100}
var drPepper = Product{Name: "Dr.Pepper", Price: 100}
var redBull = Product{Name: "Red Bull", Price: 200}

var products = []Product{cola, oolongTea, drPepper, redBull}

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
		{
			name: "various price",
			operations: []buying{
				{
					insertCount: 10,
					pushes: []pushing{
						{3, redBull.Name, nil},
						{3, redBull.Name, nil},
						{3, redBull.Name, nil},
						{3, redBull.Name, nil},
						{3, redBull.Name, nil},
						{3, "", errLackingMoney},
						{0, "", errLackingMoney},
					},
				},
				{
					insertCount: 5,
					pushes: []pushing{
						{3, redBull.Name, nil},
						{3, redBull.Name, nil},
						{3, "", errLackingMoney},
						{0, cola.Name, nil},
						{0, "", errLackingMoney},
					},
				},
				{
					insertCount: 5,
					pushes: []pushing{
						{0, cola.Name, nil},
						{1, oolongTea.Name, nil},
						{2, drPepper.Name, nil},
						{3, redBull.Name, nil},
						{0, "", errLackingMoney},
					},
				},
				{
					insertCount: 1,
					pushes: []pushing{
						{3, "", errLackingMoney},
					},
				},
				{
					insertCount: 1,
					pushes: []pushing{
						{3, redBull.Name, nil},
						{0, "", errLackingMoney},
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			vm := New(products)

			for _, ope := range tc.operations {
				t.Logf("insert 100 yen %d times", ope.insertCount)
				for i := 0; i < ope.insertCount; i++ {
					vm.Insert100Yen()
				}

				for _, p := range ope.pushes {
					t.Logf("push %d button", p.button)
					beverage, err := vm.Push(p.button)
					if beverage != p.wantBeverage {
						t.Errorf("beverage want: \"%s\", got: \"%s\"",
							p.wantBeverage,
							beverage,
						)
					}
					if !errorsAreSame(err, p.wantErr) {
						t.Errorf("error    want: \"%v\", got: \"%v\"",
							p.wantErr,
							err,
						)
					}
				}
			}
		})
	}
}

func TestVendingMachine_ButtonsDescription(t *testing.T) {
	testcases := []struct {
		name     string
		products []Product
		want     []ButtonDescription
	}{
		{
			name:     "no product",
			products: []Product{},
			want:     []ButtonDescription{},
		},
		{
			name:     "various products",
			products: products,
			want: []ButtonDescription{
				{Button: 0, Product: cola},
				{Button: 1, Product: oolongTea},
				{Button: 2, Product: drPepper},
				{Button: 3, Product: redBull},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			vm := New(tc.products)

			got := vm.ButtonsDescription()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want:\n%#v", tc.want)
				t.Errorf("got :\n%#v", got)
			}
		})
	}
}
