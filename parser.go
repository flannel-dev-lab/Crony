package Crony

var macros = map[string]string{
	"@daily": "0 0 * * *",
	"@hourly": "0 * * * *",
	"@weekly": "0 0 * * 0",
	"@monthly": "0 0 1 * *",
	"@yearly": "0 0 1 1 *",
}

func parseCron(cron string) {

}