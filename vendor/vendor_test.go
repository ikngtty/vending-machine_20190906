package vendor

import (
	"log"
	"testing"
)

func TestVendingMachine(t *testing.T) {
	const beverage = "Cola"

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
					wants:       []string{beverage, "", ""},
				},
			},
		},
		{
			name: "n coins",
			operations: []operation{
				{
					insertCount: 3,
					wants: []string{
						beverage, beverage,
					},
				},
				{
					insertCount: 5,
					wants: []string{
						beverage, beverage, beverage, beverage,
						beverage, beverage, "", "",
					},
				},
				{
					insertCount: 2,
					wants: []string{
						beverage, beverage, "", "",
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			log.Print(tc.name)

			vm := New(beverage)
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
