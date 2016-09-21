package cmds

import (
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func DnsqRun(args []string) int {

	target := args[1]
	server := makeserver(args[2])

	c := dns.Client{}
	m := dns.Msg{}

	k, ok := dns.StringToType[strings.ToUpper(args[0])]
	if !ok {
		return 111
	}

	m.SetQuestion(target+".", k)
	r, _, err := c.Exchange(&m, server)
	if err != nil {
		return 111
	}
	if len(r.Answer) == 0 {
		return 111
	}
	for _, ans := range r.Answer {
		fmt.Println(ans)
	}
	return 0
}

func makeserver(s string) string {
	/*
		we support host, host:port,
		[host]:port
	*/
	server := ""
	fi := strings.Index(s, ":")
	if fi != -1 {
		if fi == strings.LastIndex(s, ":") || strings.HasPrefix(s, "[") {
			h, p, _ := net.SplitHostPort(s)
			server = h + ":" + p
			return server
		}
	}
	server = s + ":53"
	return server

}
