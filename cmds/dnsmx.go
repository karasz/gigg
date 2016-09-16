package cmds

import (
	"fmt"
	"net"
)

func DnsmxRun(args []string) {
	for _, arg := range args {
		mxs, err := net.LookupMX(arg)
		if err != nil {
			fmt.Println(0, arg)
		} else {
			for _, mx := range mxs {
				fmt.Println(mx.Pref, mx.Host[:len(mx.Host)-1])
			}
		}
	}
}
