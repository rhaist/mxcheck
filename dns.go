package main

import (
	"log"
	"strings"

	"github.com/miekg/dns"
)

// getMX builds an MX record dns request and sends it to a dns server
// It returns a single mx entry and its status
func getMX(targetHostName *string, dnsServer string) (string, bool) {
	var mx string
	var mxstatus bool

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(*targetHostName), dns.TypeMX)

	c := new(dns.Client)
	in, _, err := c.Exchange(m, dnsServer+":53")
	e(err)


	if len(in.Answer) == 0 {
		mx = *targetHostName
	} else {
		if t, ok := in.Answer[0].(*dns.MX); ok {
			mx = t.Mx
			mxstatus = true
		}
	}

	return mx, mxstatus

}

// getA builds an A record dns request and sends it to a dns server
// It returns a single ip address
func getA(targetHostName string, dnsServer string) string {
	var a string

	m := new(dns.Msg)
	m.SetQuestion(targetHostName, dns.TypeA)

	c := new(dns.Client)
	in, _, err := c.Exchange(m, dnsServer+":53")
	e(err)

	if len(in.Answer) == 0 {
		log.Fatalln("No Answer from DNS")
	}

	if t, ok := in.Answer[0].(*dns.A); ok {
		a = t.A.String()
	}

	return a
}

// getPTR builds a PTR dns request and sends it to a dns server
// It returns a single ptr entry
func getPTR(ipaddr string, dnsServer string) string {
	var ptr string

	ipslice := strings.Split(ipaddr, ".")
	rddapi := ipslice[3] + "." + ipslice[2] + "." + ipslice[1] + "." + ipslice[0]

	rddapi = rddapi + ".in-addr.arpa"

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(rddapi), dns.TypePTR)

	c := new(dns.Client)
	in, _, err := c.Exchange(m, dnsServer+":53")
	e(err)

	if len(in.Answer) == 0 {
		ptr = "No PTR set"
	} else {
		if t, ok := in.Answer[0].(*dns.PTR); ok {
			ptr = t.Ptr
		}
	}

	return ptr
}

// getSPF builds a spf dns request and sends it to a dns server
// It returns a bool if a spf is set and has "v=spf1"
func getSPF(targetHostName string, dnsServer string) (bool, string) {
	var spf bool
	var spfanswer string

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(targetHostName), dns.TypeTXT)

	c := new(dns.Client)
	c.Net = "tcp"
	in, _, err := c.Exchange(m, dnsServer+":53")
	e(err)

	if len(in.Answer) != 0 {
		for n := range in.Answer {
			t := *in.Answer[n].(*dns.TXT)
			for _, v := range t.Txt {
				if strings.HasPrefix(v, "v=spf1") {
					spf = true
					spfanswer = in.Answer[n].String()
				}
			}
		}
	}

	return spf, spfanswer
}
