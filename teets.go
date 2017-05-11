package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var format string
	var appendflg bool
	var addstdout bool
	flag.StringVar(&format, "f", "03:04:05", "timestamp format. https://golang.org/src/time/format.go")
	flag.BoolVar(&appendflg, "a", false, "append flag")
	flag.BoolVar(&addstdout, "o", false, "add timestamp in stdout")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  %s [options] filepath\n\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	filepath := flag.Args()[0]
	var fo *os.File
	var err error

	_, err = os.Stat(filepath)
	if err == nil && appendflg {
		fo, err = os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
	} else {
		fo, err = os.Create(filepath)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	stdout := os.Stdout

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		timestamp := time.Now().Format(format)
		ptime := []byte(fmt.Sprintf("\n[%s] ", timestamp))
		pbuffer := bytes.Replace(buf, []byte("\n"), ptime, -1)

		if addstdout {
			buf = pbuffer
		}

		if _, err = stdout.Write(buf); err != nil {
			log.Fatal(err)
		}

		if _, err = fo.Write(pbuffer); err != nil {
			log.Fatal(err)
		}

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
}
