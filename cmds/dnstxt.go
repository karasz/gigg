package cmds

import (
	"fmt"
	"net"
)

func DnstxtRun(args []string) int {
	for _, arg := range args {
		txts, err := net.LookupTXT(arg)
		if err != nil {
			fmt.Println("")
			return 111
		} else {
			for _, txt := range txts {
				fmt.Println(txt)
			}
		}
	}
	return 0
}
