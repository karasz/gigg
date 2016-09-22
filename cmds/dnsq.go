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

	m.RecursionDesired = false
	m.SetQuestion(dns.Fqdn(target), k)
	r, _, err := c.Exchange(&m, server)
	if err != nil {
		fmt.Println(err)
		return 111
	}

	fmt.Println("question: ", r.Question[0].String())

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
