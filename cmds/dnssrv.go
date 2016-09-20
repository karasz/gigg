package cmds

import (
	"fmt"
	"net"
)

func DnssrvRun(args []string) int {
	// we always need exactly 3 parameters or just one
	switch len(args) {
	case 1:
		cname, addrs, err := net.LookupSRV("", "", args[0])
		if err == nil {
			printsrv(cname, addrs)
			return 0
		} else {
			fmt.Println("")
			return 111
		}

	case 3:
		cname, addrs, err := net.LookupSRV(args[0], args[1], args[2])
		if err == nil {
			printsrv(cname, addrs)
			return 0
		} else {
			fmt.Println("")
			return 111
		}
	default:
		fmt.Println("")
		return 111
	}
	return 111
}

func printsrv(cname string, addrs []*net.SRV) {
	for _, k := range addrs {
		fmt.Println("CNAME: ", cname, "Target: ", k.Target, "Port: ", k.Port, "Priority: ", k.Priority, "Weight: ", k.Weight)
	}
}
