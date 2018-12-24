package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gosagawa/xmastree/xmastree"
)

func main() {

	var (
		size  = flag.Int("size", 8, "size")
		speed = flag.Int("speed", 100, "speed")
	)
	flag.Usage = usage
	flag.Parse()

	xmastree.Display(*size, *speed)
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
