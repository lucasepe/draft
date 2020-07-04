package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasepe/draft"
	"github.com/mitchellh/go-homedir"
)

// go run main.go | dot -Tpng -Gdpi=300 > test.png

const (
	envIconsPath = "DRAFT_ICONS_PATH"
	banner       = `
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
	flagVerbose   bool
	flagImpl      bool
)

func main() {
	if dir := os.Getenv(envIconsPath); len(dir) == 0 {
		setDefaultIconsPath()
	}

	configureFlags()

	if flag.CommandLine.Arg(0) == "" {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	uri := flag.Args()[0]

	cfg := draft.NewConfig(
		draft.Verbose(flagVerbose),
		draft.BottomTop(flagBottomTop),
		draft.IconsPath(os.Getenv(envIconsPath)),
		draft.ShowImpl(flagImpl),
		draft.URI(uri),
	)

	dia, err := draft.Sketch(cfg)
	handleErr(err, uri)

	fmt.Println(dia)
}

// handleErr check for an error and eventually exit
func handleErr(err error, src string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s @ %s\n", err.Error(), src)
		os.Exit(1)
	}
}

func configureFlags() {
	name := appName()

	flag.CommandLine.Usage = func() {
		printBanner()
		fmt.Printf("Generate High Level Cloud Architecture diagrams using YAML.\n\n")

		fmt.Print("USAGE:\n\n")
		fmt.Printf("  %s [options] <yaml file or url>\n\n", name)

		fmt.Print("EXAMPLE(s):\n\n")
		fmt.Printf("  %s input.yml | dot -Tpng -Gdpi=200 > output.png\n", name)
		fmt.Printf("  %s http://a.domain.com/input.yml | dot -Tpng -Gdpi=200 > output.png\n\n", name)

		fmt.Print("OPTIONS:\n\n")
		flag.CommandLine.SetOutput(os.Stdout)
		flag.CommandLine.PrintDefaults()
		flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
		fmt.Print("  -help\n\tprints this message\n")
		fmt.Println()

		fmt.Print("ENVIRONMENT:\n\n")
		fmt.Fprint(os.Stdout, "  DRAFT_ICONS_PATH\n")
		fmt.Fprintf(os.Stdout, "  \tthe base path for custom icons (default %s)\n", os.Getenv(envIconsPath))
		fmt.Println()
	}

	flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError)

	flag.CommandLine.BoolVar(&flagImpl, "impl", false, "if true show the component provider implementation")
	flag.CommandLine.BoolVar(&flagBottomTop, "bottom-top", false, "if true sets layout direction as bottom top")
	flag.CommandLine.BoolVar(&flagVerbose, "verbose", false, fmt.Sprintf("show some extra info as %s is running", name))

	flag.CommandLine.Parse(os.Args[1:])
}

func printBanner() {
	fmt.Print(strings.Trim(strings.Replace(banner, "{{VERSION}}", version, 1), "\n"), "\n\n")
}

// setDefaultIconsPath sets the default icons directory
// creating it if does not exists.
func setDefaultIconsPath() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	workdir := filepath.Join(home, fmt.Sprintf(".%s", appName()), "icons")
	if _, err := os.Stat(workdir); os.IsNotExist(err) {
		err = os.MkdirAll(workdir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	os.Mkdir(filepath.Join(workdir, "default"), os.ModePerm)

	os.Setenv(envIconsPath, workdir)
	return nil
}

func appName() string {
	return filepath.Base(os.Args[0])
}
