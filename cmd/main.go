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
|___/  |_|    \__,_||_|   \__|               v{{VERSION}}`
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	flagBottomTop bool
	flagOrtho     bool
	flagImpl      string
)

func main() {
	configureFlags()

	if flag.CommandLine.Arg(0) == "" {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	fn, err := filepath.Abs(flag.Args()[0])
	handleErr(err, fn)

	file, err := os.Open(fn)
	handleErr(err, fn)

	defer file.Close()

	ark, err := draft.NewDraft(file)
	handleErr(err, fn)

	ark.BottomTop(flagBottomTop)
	ark.Ortho(flagOrtho)
	ark.Provider(flagImpl)

	str, err := ark.Sketch()
	handleErr(err, fn)

	fmt.Println(str)
}

// handleErr check for an error and eventually exit
func handleErr(err error, src string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s @ %s\n", err.Error(), src)
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

	flag.CommandLine.BoolVar(&flagBottomTop, "bottom-top", false, "if true sets layout dir as bottom top")
	flag.CommandLine.BoolVar(&flagOrtho, "ortho", false, "if true edges are drawn as line segments")
	flag.CommandLine.StringVar(&flagImpl, "impl", "", "auto fill the specific provider services (aws, gcp or azure)")

	flag.CommandLine.Parse(os.Args[1:])
}

func printBanner() {
	fmt.Print(strings.Trim(strings.Replace(banner, "{{VERSION}}", version, 1), "\n"), "\n\n")
}
