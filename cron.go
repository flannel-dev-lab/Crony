package Crony

import "errors"

type Crony struct {
	Macros map[string]string
}

func New() *Crony {
	crony := Crony{}

	crony.Macros =  map[string]string{
		"@daily":   "0 0 * * *",
		"@hourly":  "0 * * * *",
		"@weekly":  "0 0 * * 0",
		"@monthly": "0 0 1 * *",
		"@yearly":  "0 0 1 1 *",
	}

	return &crony
}

// AddMacro Adds a custom macro to the available set of macros
func (crony *Crony) AddMacro(macroKey, macroValue string) (err error) {
	if macroKey != "" && macroValue != "" {
		crony.Macros[macroKey] = macroValue
		return nil
	} else {
		return errors.New("macros cannot bee empty when defining custom macros")
	}

}