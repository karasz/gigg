package cmds

import (
	"fmt"
	"net"
)

func DnsnameRun(args []string) {
	for _, ip := range args {
		name, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Println("")
		} else {
			fmt.Println(name[0][:len(name[0])-1])
		}
	}
}
