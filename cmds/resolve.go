package cmds

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseResolv() ([]string, []string, []string) {
	var nameservers, search, domain []string = make([]string, 0), make([]string, 0), make([]string, 0)

	file, err := os.Open("/etc/resolv.conf")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "nameserver") {
			nameservers = append(nameservers, strings.Fields(line)[1])
		}
		if strings.HasPrefix(line, "domain") {
			domain = append(domain, strings.Fields(line)[1])
		}
		if strings.HasPrefix(line, "search") {
			line = strings.TrimPrefix(line, "search")
			for _, v := range strings.Fields(line) {
				if v != "" {
					search = append(search, v)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return nameservers, domain, search
}
