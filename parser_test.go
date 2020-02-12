package Crony

import "testing"

func TestParseCron(t *testing.T) {
	crony := New()
	crony.parseCron("* * * * *")
}

func TestParseCronWithMacro(t *testing.T) {
	crony := New()
	crony.parseCron("* * * * *")
}

func TestParseCronWithCustomMacro(t *testing.T) {
	crony := New()
	crony.parseCron("* * * * *")
}