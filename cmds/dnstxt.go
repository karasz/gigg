package cmds

import (
	"fmt"
	"net"
)

func DnstxtRun(args []string) {
	for _, arg := range args {
		txts, err := net.LookupTXT(arg)
		if err != nil {
			fmt.Println("")
		} else {
			for _, txt := range txts {
				fmt.Println(txt)
			}
		}
	}
}
