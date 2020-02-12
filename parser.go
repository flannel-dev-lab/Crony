package Crony

import (
	"fmt"
	"strings"
)

func (crony *Crony) parseCron(cron string) {

	if crony.Macros[cron] != "" {
		// Getting cron from default macros
		cron = crony.Macros[cron]
	}

	intervals := strings.Split(cron, " ")
	fmt.Println(intervals)
}
