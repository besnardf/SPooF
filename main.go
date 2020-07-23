package main
 
import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
	  return fmt.Sprintf(colorString,
		fmt.Sprint(args...))
	}
	return sprint
}

func main() {

	var domains []string

	sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			domains = append(domains, sc.Text())
		}

		if err := sc.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	for _, domain := range domains {
		vulnerable := true;
		txtrecords, _ := net.LookupTXT(domain)
	
		for _, txt := range txtrecords {

			// naive implementation as what follows could be vulnerable
			if strings.HasPrefix(txt, "v=spf") {
				vulnerable = false
				fmt.Println(Green(domain))
			}
		}

		if vulnerable == true {
			fmt.Println(Red(domain))
		}
	}
}
