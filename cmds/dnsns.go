package cmds

import (
	"fmt"
	"net"
)

func DnsnsRun(args []string) int {
	for _, arg := range args {
		nss, err := net.LookupNS(arg)
		if err != nil {
			fmt.Println("")
			return 111
		} else {
			for _, ns := range nss {
				fmt.Println(ns.Host[:len(ns.Host)-1])
			}
		}
	}
	return 0
}
