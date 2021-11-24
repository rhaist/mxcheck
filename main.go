// mxcheck is a security scanner for mail servers
package main

import (
	. "github.com/logrusorgru/aurora"
	flag "github.com/spf13/pflag"
	"log"
)

func main() {

	println()
	println(versionmsg)

	dnsServer := flag.StringP("dnsserver", "d", "8.8.8.8", "The dns server to consult")
	mailFrom := flag.StringP("mailfrom", "f", "infi@foo.wtf", "Set the mail from address")
	mailTo := flag.StringP("mailto", "t", "info@baz.wtf", "Set the mail to address")
	targetHostName := flag.StringP("host","h", "localhost", "The target host to check")
	verbose := flag.BoolP("verbose", "v", false, "verbose")

	flag.Parse()


	log.Println("ii Checking: " + *targetHostName)

	targetHost, mxstatus := getMX(targetHostName, *dnsServer)
	if mxstatus {
		log.Println("ii Found MX: " + targetHost)
	} else {
		log.Println("-- No MX entry found. Using Target Host Name.")
	}

	log.Println("ii Checking for A record")
	ipaddr := getA(targetHost, *dnsServer)
	log.Println("ii IP address MX: " + ipaddr)

	log.Println("ii Checking for PTR record")
	ptrentry := getPTR(ipaddr, *dnsServer)
	log.Println("ii PTR entry: " + ptrentry)

	if ptrentry == targetHost {
		log.Println(Green("++ PTR matches MX record"))
	} else {
		log.Println(Red("-- PTR does not match MX record"))
	}

	log.Println("ii Checking for SPF record")
	spfentry := getSPF(*targetHostName, *dnsServer, *verbose)
	if spfentry {
		log.Println(Green("++ SPF set"))
	} else {
		log.Println(Red("-- No SPF set"))
	}

	log.Println("ii Checking for open mail ports")
	openPorts := portScan(targetHost)
	log.Print("ii Open ports: ", openPorts)

	if len(openPorts) == 0 {
		log.Println(Cyan("ii No open ports to connect to. Quitting."))
		return
	}

	for _, port := range openPorts {
		if port == "25" {
			log.Println("ii Checking for open relay")
			tlsresult, orresult := openRelay(*mailFrom, *mailTo, targetHost)

			if tlsresult {
				log.Println(Green("++ StartTLS supported"))
			} else {
				log.Println("-- StartTLS not supported")
			}

			if orresult {
				log.Println(Red("!! Server is probably an open relay"))
			} else {
				log.Println(Green("++ Server is not an open relay"))
			}
			println()
		}
	}

}
