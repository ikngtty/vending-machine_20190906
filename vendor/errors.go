package vendor

import "fmt"

type LackingMoneyError struct{}

func (err LackingMoneyError) Kind() string {
	return "lacking_money"
}
func (err LackingMoneyError) Error() string {
	return "need more money"
}

type InvalidButtonError struct {
	pushed int
}

func (err InvalidButtonError) Kind() string {
	return "invalid_button"
}
func (err InvalidButtonError) Error() string {
	return fmt.Sprintf("given button does not exist: %d", err.pushed)
}
