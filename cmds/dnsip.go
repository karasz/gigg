package cmds

import (
	"fmt"
	"net"
	//	"strings"
	//	"github.com/miekg/dns"
)

func DnsipRun(args []string, ipv6 bool) int {
	for _, arg := range args {
		ips := make([]net.IP, 0)
		ip, err := net.LookupIP(arg)
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
		for _, k := range ips {
			fmt.Print(k)
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	return 0
}
