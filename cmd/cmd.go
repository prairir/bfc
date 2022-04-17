package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/prairir/bfc/pkg/b2c"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("expected input file")
		os.Exit(1)
	}

	inFile := flag.String("in-file", "", "input file in BrainFuck (no default)")
	outFile := flag.String("out-file", "", "output file (default print to stdout)")
	size := flag.Int("memory-size", 30_000, "memory size of BrainFuck machine")

	flag.Parse()

	if *inFile == "" {
		fmt.Printf("in-file required\n")
		os.Exit(1)
	}

	bf, err := ioutil.ReadFile(*inFile)
	if err != nil {
		fmt.Printf("couldn't open file: %q, got error: %s\n", os.Args[1], err)
		os.Exit(1)
	}

	message := b2c.Transpile(string(bf), uint(*size))

	if *outFile != "" {
		os.WriteFile(*outFile, []byte(message), 0666)
	} else {
		fmt.Printf("%s\n", message)
	}

}
