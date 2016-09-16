package cmds

import (
	"fmt"
	"net"
)

func DnsnameRun(args []string) int {
	for _, ip := range args {
		name, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Println("")
			return 111
		} else {
			fmt.Println(name[0][:len(name[0])-1])
		}
	}
	return 0
}
