package pkg

import (
	"github.com/jlaffaye/ftp"
	"log"
	"os"
	"time"
)

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
