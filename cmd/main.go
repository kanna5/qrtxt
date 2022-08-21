package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kanna5/qrtxt"
)

func help() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] <text to be encoded>\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.Usage = help
}

var (
	levelLookup = map[string]qrtxt.RecoveryLevel{
		"low":     qrtxt.Low,
		"medium":  qrtxt.Medium,
		"high":    qrtxt.High,
		"highest": qrtxt.Highest,
	}
)

func main() {
	levelFlag := flag.String("l", "low", "Recovery level. One of: low, medium, high, highest")
	flag.Parse()

	text := flag.Args()
	if len(text) == 0 {
		help()
		os.Exit(1)
	}

	lvl, ok := levelLookup[*levelFlag]
	if !ok {
		help()
		os.Exit(1)
	}

	qr, err := qrtxt.Encode(strings.Join(text, " "), lvl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
		os.Exit(1)
	}
	fmt.Println(qr)
}
