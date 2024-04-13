package main

import (
	"testing"
)

func TestArithOperand(t *testing.T) {
	r1 := 4 + 3*4 + 1
	r2 := (4+3)*4 + 1
	r3 := 4 + 3*(4+1)
	r4 := 4 + 3*(4+1) + 2
	r5 := 4 + 3*(4+1) + 2*5
	r6 := 4 + 3*(4+1) + 2*5/2
	t.Log(r1, r2, r3, r4, r5, r6)
	if Evaluate("4+3*4+1") != r1 {
		t.Fail()
	}
	if Evaluate("(4+3)*4+1") != r2 {
		t.Fail()
	}
	if Evaluate("4+3*(4+1)") != r3 {
		t.Fail()
	}
	if Evaluate("4+3*(4+1)+2") != r4 {
		t.Fail()
	}
	if Evaluate("4+3*(4+1)+2*5") != r5 {
		t.Fail()
	}
	if Evaluate("4+3*(4+1)+2*5/2") != r6 {
		t.Fail()
	}

}
