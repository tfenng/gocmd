package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	version = "0.0.1"
)

var (
	ver  bool
	help bool
	ts   = flag.String("timestamp", "", "Unix style system timestamp number")
)

func init() {
	flag.BoolVar(&ver, "v", false, "Print Version Number")
	flag.BoolVar(&ver, "version", false, "Print Version Number")
	flag.BoolVar(&help, "h", false, "Print usage")
	flag.BoolVar(&help, "help", false, "Print usage")
}

func echoNow() {
	fmt.Printf("%v\n", NowTs())
}

func prefix(tss string) string {
	now_str := strconv.Itoa(int(NowTs()))
	return fmt.Sprintf("%s%s", now_str[0:10-len(tss)], tss)
}

func echo(tss string) {
	var fixedTs string
	if len(tss) < 10 {
		fixedTs = prefix(tss)
	} else {
		fixedTs = tss
	}
	ts, err := strconv.Atoi(fixedTs)
	if err != nil {
		log.Fatalln("Invalid timestamp number", err)
		return
	}
	fmt.Printf("%v\n", FormatTs(ts))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		*ts = args[0]
	}
	if ver {
		fmt.Println("Version:", version)
		os.Exit(2)
	}
	if help {
		usage()
	}
	if *ts == "" {
		echoNow()
	} else {
		echo(*ts)
	}
}

var usageinfo string = `vts is a Go implemented CLI tool for display current system timestamp number, or convert someone timestamp into formated string which human read.

Usage:

	vts [flags] [TIMESTAMP]

flags:
  -v, -version=true             Show Version Number
  -h, -help=true                Show usage info

TIMESTAMP:
  The Unix style timestamp, in integer number (eg.1563270875, or just 3270875).

Example:
	vts
	vts 1563270875
	vts     270875
	vts        875
	vts | xargs vts

more help information please refer to https://github.com/tfenng/vts
`

func usage() {
	fmt.Println(usageinfo)
	os.Exit(2)
}
