// package main starts the Planban application
package main

import (
	"fmt"
	"os"

	"github.com/callerobertsson/planban/app"
)

var pbFile = "planban.json" // default board file name

func init() {
	if len(os.Args) > 1 {
		pbFile = os.Args[1]
	}
}

func main() {
	pb, err := app.New(pbFile)
	if err != nil {
		fail(err)
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
