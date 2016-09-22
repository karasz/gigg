package cmds

import (
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func DnsqrRun(args []string) int {
	conf, err := dns.ClientConfigFromFile("/etc/resolv.conf")
	if err != nil {
		fmt.Println(err)
		return 111
	}
	target := args[1]
	server := conf.Servers[0]
	if server[0] == '[' && server[len(server)-1] == ']' {
		server = server[1 : len(server)-1]
	}

	if i := net.ParseIP(server); i != nil {
		server = net.JoinHostPort(server, "53")
	} else {
		server = dns.Fqdn(server) + ":53"
	}

	c := dns.Client{}
	m := dns.Msg{}

	k, ok := dns.StringToType[strings.ToUpper(args[0])]
	if !ok {
		return 111
	}

	m.RecursionDesired = true
	m.SetQuestion(dns.Fqdn(target), k)
	r, _, err := c.Exchange(&m, server)
	if err != nil {
		fmt.Println(err)
		return 111
	}

	fmt.Println("question:", r.Question[0].String())

	if len(r.Answer) == 0 {
		for _, aut := range r.Ns {
			fmt.Println("authority:", aut)
		}
		for _, ex := range r.Extra {
			fmt.Println("additional:", ex)
		}
		return 0
	}

	for _, ans := range r.Answer {
		fmt.Println("answer:", ans)
	}
	for _, aut := range r.Ns {
		fmt.Println("authority:", aut)
	}

	return 0
}
