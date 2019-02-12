package main

import (
	"testing"
)

func TestGrade(t *testing.T) {
	tests := []struct{ a int; b string } {
		{59, "F"},
		{61, "C"},
		{81, "B"},
		{91, "A"},
	}
	for _, test := range tests {
		if i := grade(test.a); i != test.b {
			t.Errorf("grade(%d):" + "%s; expected %s", test.a,  i, test.b)
		}
	}
}

func BenchmarkGrade(t *testing.B) {
	a := 59
	b := "F"
	for i := 0; i < t.N; i++ {
		if ans := grade(a); ans != b {
			t.Errorf("grade(%d):"+"%s; expected %s", a, ans, b)
		}
	}
}
