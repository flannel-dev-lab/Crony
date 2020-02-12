package Crony

import "testing"

func TestCrony_AddMacro(t *testing.T) {
	crony := New()

	if err := crony.AddMacro("@randomly", "* * * * *"); err != nil {
		t.Errorf("[%s] %v", t.Name(), err)
	}

	if crony.Macros["@randomly"] != "* * * * *" {
		t.Errorf("[%s] macros do not match", t.Name())
	}
}

func TestCrony_AddMacro_EmptyMacros(t *testing.T) {
	crony := New()

	if err := crony.AddMacro("", ""); err == nil {
		t.Errorf("[%s] no macros did not raise any error", t.Name())
	}
}