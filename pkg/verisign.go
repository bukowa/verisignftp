package pkg

import (
	"bufio"
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

var DomainBadStarts = []string{";", "@", "$"}

func VerisignDownload(login, pass, host, zone string, where *os.File)  {
	log.Printf("%v", "Creating ftp Dial...")
	c, err := ftp.Dial(host, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", "Logging in...")
	err = c.Login(login, pass)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v %v %v %v", "Downloading", zone, "file to", where.Name())
	IoCopyPanic(where, FtpDownloadPanic(c, zone))
	log.Printf("%v", "Done!")
}

func ExtractDomains(zone string, what, where *os.File) {
	var sep string

	if zone == "root.zone.gz" {
		sep = "\t"
	} else if zone == "com.zone.gz" {
		sep = " "
	}
	scanner := bufio.NewScanner(what)
	
	before := ""
	for scanner.Scan() {
		line := scanner.Text()
		domain := strings.Split(line, sep)[0]
		domain = strings.TrimSuffix(domain, ".")
		// skip
		if domain == before {
			continue
		} else {
			before = domain
		}
		// check
		if len(domain) == 0 {
			continue
		}
		if stringInSlice(string(domain[0]), DomainBadStarts) {
			continue
		}
		parsed, err := url.Parse(fmt.Sprintf("http://%v", domain))
		// skip
		if err != nil {
			log.Printf("Error parsing domain %v: %v", domain, err)
			continue
		} else {
			// write domain to file
			if _, err := where.WriteString(fmt.Sprintf("%v\n", parsed.Host)); err != nil {
				log.Printf("Error writing to file: %v", err)
			}
		}
	}
}