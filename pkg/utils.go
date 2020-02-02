package pkg

import (
	"compress/gzip"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"os"
)

func UnzipFile(what, where *os.File) {
	log.Printf("%v %v %v %v", "Extracting file", what.Name(), "to", where.Name())
	IoCopyPanic(where, GzipReaderPanic(what))
}

func FileCreatePanic(name string) *os.File {
	if file, err := os.Create(name); err == nil {
		return file
	} else {
		log.Fatal(err)
		return nil
	}
}

func FileOpenPanic(name string) *os.File {
	if file, err := os.Open(name); err == nil {
		return file
	} else {
		log.Fatal(err)
		return nil
	}
}

func FtpDownloadPanic(client *ftp.ServerConn, path string) *ftp.Response {
	if resp, err := client.Retr(path); err == nil {
		return resp
	} else {
		log.Fatal(err)
		return nil
	}
}

func IoCopyPanic(dst io.Writer, src io.Reader) (written int64){
	if written, err := io.Copy(dst, src); err == nil {
		return written
	} else {
		log.Fatal(err)
		return -1
	}
}

func GzipReaderPanic(file *os.File) *gzip.Reader {
	if reader, err := gzip.NewReader(file); err == nil {
		return reader
	} else {
		log.Fatalf("%v %v", "Error gzip reader", err)
		return nil
	}
}
