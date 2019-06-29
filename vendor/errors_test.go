package vendor

import (
	"fmt"
	"testing"
)

func TestErrorsAreSame(t *testing.T) {
	testcases := []struct {
		left  PkgError
		right PkgError
		want  bool
	}{
		{nil, nil, true},
		{nil, LackingMoneyError{}, false},
		{LackingMoneyError{}, nil, false},
		{LackingMoneyError{}, LackingMoneyError{}, true},
		{LackingMoneyError{}, InvalidButtonError{}, false},
		{InvalidButtonError{pushed: 42}, InvalidButtonError{pushed: 42}, true},
		{InvalidButtonError{pushed: 42}, InvalidButtonError{pushed: 24}, false},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			got := errorsAreSame(tc.left, tc.right)
			if got != tc.want {
				t.Errorf("want: %v", tc.want)
				t.Errorf("got : %v", got)
			}
		})
	}
}
