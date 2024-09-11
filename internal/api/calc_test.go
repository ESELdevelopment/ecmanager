package calc_test

import (
	calc "ecmanager/internal/api"
	"testing"
)

func TestMain(t *testing.T) {
	test := calc.Calculate()
	if test != 2 {
		t.Errorf("Result was incorrect")
	}
}
