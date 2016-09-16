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
	switch calledAs {
	case "gigg":
		fmt.Println("Called as gigg, that should not happen yet.")
	case "dnsip":
		c.DnsipRun(args, false)
	case "dnsip6":
		c.DnsipRun(args, true)
	case "dnsname":
		c.DnsnameRun(args)
	case "dnsmx":
		c.DnsmxRun(args)
	default:
		fmt.Println("Called as ", calledAs, ". I don't recognize that name")
	}
}
