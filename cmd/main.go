package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasepe/draft"
)

// go run main.go | dot -Tpng -Gdpi=300 > test.png

const (
	maxFileSize = 500 * 1024
	banner      = `
______                __  _   
|  _  \              / _|| |        Crafted with passion by Luca Sepe
| | | | _ __   __ _ | |_ | |_       
| | | || '__| / _' ||  _|| __|      https://github.com/lucasepe/draft
| |/ / | |   | (_| || |  | |_ 
|___/  |_|    \__,_||_|   \__|                {{VERSION}}`
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	configureFlags()

	if flag.CommandLine.Arg(0) == "" {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	fn, err := filepath.Abs(flag.Args()[0])
	handleErr(err)

	file, err := os.Open(fn)
	handleErr(err)

	defer file.Close()

	ark, err := draft.NewDraft(file)
	handleErr(err)

	str, err := ark.Sketch()
	handleErr(err)

	fmt.Println(str)
}

// handleErr check for an error and eventually exit
func handleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}

func configureFlags() {
	flag.CommandLine.Usage = func() {
		printBanner()
		fmt.Printf("Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax.\n\n")

		name := filepath.Base(os.Args[0])

		fmt.Print("USAGE:\n\n")
		fmt.Printf("  %s [options] /path/to/yaml/file\n\n", name)

		fmt.Print("EXAMPLE:\n\n")
		fmt.Printf("  %s input.yml | dot -Tpng -Gdpi=200 > output.png\n\n", name)

		fmt.Print("OPTIONS:\n\n")
		flag.CommandLine.SetOutput(os.Stdout)
		flag.CommandLine.PrintDefaults()
		flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
		fmt.Print("  -help\n\tprints this message\n")
		fmt.Println()
	}

	flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError)

	flag.CommandLine.Parse(os.Args[1:])
}

func printBanner() {
	fmt.Print(strings.Trim(strings.Replace(banner, "{{VERSION}}", version, 1), "\n"), "\n\n")
}
