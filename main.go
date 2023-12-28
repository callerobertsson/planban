// package main starts the Planban application
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/callerobertsson/planban/app"
)

var pbFile = "planban.json" // default board file name

var (
	helpFlag  = false
	printFlag = false
)

func init() {
	flag.BoolVar(&helpFlag, "h", false, "print help and exit")
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

	if helpFlag {
		help()
		os.Exit(0)
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
func help() {
	fmt.Printf(`Planban is a simple command line Kanban application.

Usage:
  planban [flags] [<filename>.json]

  If the file doesn't exist it will be created. If no filename is given,
  the default filename "planban.json" will be used.

  The available commands can be shown by pressing enter key in planban.

Options:
`)
	flag.PrintDefaults()

	fmt.Printf(`
Please visit github.com/callerobertsson/planban for more information
`)
}
