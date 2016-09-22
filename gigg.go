package main

import (
	"fmt"
	"os"
	"path/filepath"

	c "github.com/karasz/gigg/cmds"
)

func main() {
	_, calledAs := filepath.Split(os.Args[0])
	args := os.Args[1:]
	res := 0
	switch calledAs {
	case "gigg":
		fmt.Println("Called as gigg, that should not happen yet.")
	case "dnsip":
		res = c.DnsipRun(args, false)
	case "dnsip6":
		res = c.DnsipRun(args, true)
	case "dnsipq":
		res = c.DnsipqRun(args, false)
	case "dnsip6q":
		res = c.DnsipqRun(args, true)
	case "dnsname":
		res = c.DnsnameRun(args)
	case "dnsmx":
		res = c.DnsmxRun(args)
	case "dnstxt":
		res = c.DnstxtRun(args)
	case "dnsrandip":
		res = c.DnsrandipRun(false)
	case "dnsrandip6":
		res = c.DnsrandipRun(true)
	case "dnssrv":
		res = c.DnssrvRun(args)
	case "dnsns":
		res = c.DnsnsRun(args)
	case "dnsq":
		res = c.DnsqRun(args)
	case "dnsqr":
		res = c.DnsqrRun(args)
	default:
		fmt.Println("Called as ", calledAs, ". I don't recognize that name")
	}
	os.Exit(res)
}
