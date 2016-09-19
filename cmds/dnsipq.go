package cmds

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func DnsipqRun(args []string, ipv6 bool) int {
	var suff []string
	dnsq := os.Getenv("DNSQUALIFY")
	if len(dnsq) != 0 {
		for _, k := range strings.Fields(dnsq) {
			suff = append(suff, k)
		}
	}
	if len(suff) == 0 {
		_, suff, _ = parseResolv()
	}
	if len(suff) != 0 {
		for _, suf := range suff {
			for _, arg := range args {
				ips := make([]net.IP, 0)
				ip, err := net.LookupIP(arg + suf)
				if err != nil {
					ip = []net.IP{}
					return 111
				}
				if !ipv6 {
					for _, i := range ip {
						if i.To4() != nil {
							ips = append(ips, i.To4())
						}
					}
				} else {
					for _, i := range ip {
						ips = append(ips, i.To16())
					}
				}
				fmt.Print(arg + suf)
				fmt.Print(" ")
				for _, k := range ips {
					fmt.Print(k)
					fmt.Print(" ")
				}
				fmt.Println("")
			}
		}
	}
	return 0
}
