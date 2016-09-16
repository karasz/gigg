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
	case "dnsname":
		res = c.DnsnameRun(args)
	case "dnsmx":
		res = c.DnsmxRun(args)
	case "dnstxt":
		res = c.DnstxtRun(args)
	default:
		fmt.Println("Called as ", calledAs, ". I don't recognize that name")
	}
	os.Exit(res)
}
