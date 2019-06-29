package vendor

import (
	"fmt"
	"reflect"
)

type PkgError interface {
	thisIsVendorError() // A marker interface
	Kind() string
	Error() string
}

func errorsAreSame(left, right PkgError) bool {
	return reflect.DeepEqual(left, right)
}

type LackingMoneyError struct{}

func (err LackingMoneyError) thisIsVendorError() {}
func (err LackingMoneyError) Kind() string {
	return "lacking_money"
}
func (err LackingMoneyError) Error() string {
	return "need more money"
}

type InvalidButtonError struct {
	pushed int
}

func (err InvalidButtonError) thisIsVendorError() {}
func (err InvalidButtonError) Kind() string {
	return "invalid_button"
}
func (err InvalidButtonError) Error() string {
	return fmt.Sprintf("given button does not exist: %d", err.pushed)
}
