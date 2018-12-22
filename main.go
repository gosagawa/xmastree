package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		sample = flag.String("s", "sample", "sampele text")
	)
	flag.Usage = usage
	flag.Parse()

	if *sample == "something invalid" {
		usage()
		os.Exit(2)
	}

	fmt.Println("hello world")

}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  -s sample\n")
	fmt.Fprintf(os.Stderr, "    	sample \n")
}
