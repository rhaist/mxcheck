# mxcheck

mxcheck is an info scanner for e-mail servers.

It checks 
  * DNS records: A, MX, PTR, SPF, MTA-STS, DKIM
  * for support of StartTLS and the certificate
  * open ports: 25, 465, 587
  * and if the server is an open relay

You can set mailFrom, mailTo, the DNS server, DKIM selector and output a report in tsv format.


    -d, --dnsserver string   The dns server to be requested (default "8.8.8.8")
    -f, --mailfrom string    Set the mailFrom address (default "info@foo.wtf")
    -t, --mailto string      Set the mailTo address (default "info@baz.wtf")
    -n, --no-prompt          Answer yes to all questions
    -s, --service string     The service host to check (default "localhost")
    -S, --dkim-selector      The DKIM selector. If set a dkim check is performed on the provided service domain
    -v, --version            Version and license
    -w, --write-tsv          Write tsv formated report to file
    


# Version

v1.2.4

[![Go Report Card](https://goreportcard.com/badge/github.com/steffenfritz/mxcheck)](https://goreportcard.com/report/github.com/steffenfritz/mxcheck) 
[![Go Reference](https://pkg.go.dev/badge/github.com/steffenfritz/mxcheck.svg)](https://pkg.go.dev/github.com/steffenfritz/mxcheck)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Build status](https://ci.appveyor.com/api/projects/status/l6d32n4ax02f6ku2?svg=true)](https://ci.appveyor.com/project/steffenfritz/mxcheck)

# Installation


    go install github.com/steffenfritz/mxcheck
    
or

    download a pre-compiled binary.

# Usage Example

    ./mxcheck -s 2600.com
    ./mxcheck -s 2600.com -v
    ./mxcheck -s 2600.com -v -d 8.8.8.8
    ./mxcheck -s 2600.com -v -n -f info@baz.com -t boss@foo.org -w -S default
    
   [![asciicast](https://asciinema.org/a/471229.svg)](https://asciinema.org/a/471229)
    

There is no check whether the server needs authentication. However, you can do two runs:

The first one uses a from and to address outside the mail server's scope, e.g.:

    ./mxcheck -s example.com -f info@baz.com -t boss@foo.org

The second one uses a from and a to address from the mail server's scope, e.g.:

    ./mxcheck -s example.com -f info@example.com -t boss@example.com

If the first one returns ``Server is not an open relay`` and the second one returns `Server is probably an open relay` the server is not an open relay, but you can send mails from local to local addresses.

