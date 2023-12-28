// package main starts the Planban application
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/callerobertsson/planban/app"
)

var pbFile = "planban.json" // default board file name

var printFlag = false

func init() {
	flag.BoolVar(&printFlag, "p", false, "print board and exit")
	flag.Parse()

	file := flag.Arg(0)
	if file != "" {
		pbFile = file
	}
}

func main() {
	pb, err := app.New(pbFile)
	if err != nil {
		fail(err)
	}

	if printFlag {
		pb.PrintBoard()
		os.Exit(0)
	}

	if err := pb.Run(); err != nil {
		fail(err)
	}

	fmt.Println("bye!")
}

func fail(a ...any) {
	fmt.Fprintf(os.Stderr, "Exit with error: %q\n", a...)
	os.Exit(1)
}
