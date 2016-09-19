package cmds

import (
	"fmt"
	"math/rand"
	"time"
)

func DnsrandipRun(ipv6 bool) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if !ipv6 {
		fmt.Print(r.Intn(255), ".", r.Intn(256), ".", r.Intn(256), ".", r.Intn(256), "\n")
	} else {
		fmt.Print(hexfromint(r.Intn(65535)), ":", hexfromint(r.Intn(65535)), ":", hexfromint(r.Intn(65535)), ":", hexfromint(r.Intn(65535)), ":", hexfromint(r.Intn(65535)), ":", hexfromint(r.Intn(65535)), ":", hexfromint(r.Intn(65535)), ":", hexfromint(r.Intn(65535)), "\n")
	}
	return 0
}

func hexfromint(ii int) string {
	return fmt.Sprintf("%x", ii)
}
