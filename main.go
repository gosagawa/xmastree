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
	usage := ""
	usage += "Usage: %s [OPTION]\n"
	usage += "  -s sample\n"
	usage += "    	sample \n"
	_, err := fmt.Fprintf(os.Stderr, usage)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
