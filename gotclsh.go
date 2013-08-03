// Niels Widger
// Time-stamp: <02 Aug 2013 at 20:46:12 by nwidger on macros.local>

package main

import (
	"fmt"
	"github.com/nwidger/gotcl"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Printf("%s: FILE\n", os.Args[0])
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", os.Args[1], err)
		os.Exit(1)
	}

	script := string(bytes)
	interp := gotcl.NewInterp()
	interp.Eval(script)
}
