package main

import (
	"fmt"
	"flag"
	"os"
	"path"

	"github.com/caltechlibrary/patentfetch"
)

func main() {
	appName := path.Base(os.Args[0])
	licenseText, version, releaseDate, releaseHash := patentfetch.LicenseText, patentfetch.Version, patentfetch.ReleaseDate, patentfetch.ReleaseHash
	helpText := patentfetch.HelpText
	showHelp, showLicense, showVersion := false, false, false	
	fmtHelp := patentfetch.FmtHelp

	flag.BoolVar(&showHelp, "help", showHelp, "display help")
	flag.BoolVar(&showLicense, "license", showLicense, "display license")
	flag.BoolVar(&showVersion, "version", showVersion, "display version")
	flag.Parse()
	args := flag.Args()

	//in := os.Stdin
	out := os.Stdout
	eout := os.Stderr

	if showHelp {
		fmt.Fprintf(out, "%s\n", fmtHelp(helpText, appName, version, releaseDate, releaseHash))
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(out, "%s\n", licenseText)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(out, "%s %s\n", version, releaseHash)
		os.Exit(0)
	}
	if len(args) != 1 {
		fmt.Fprintf(eout, "expected the CSV filename downloaded from Google Patent Search")
		os.Exit(1)
	}

	src, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(eout, "error reading %s, %s\n", args[0], err)
		os.Exit(1)

	}
	data, err := patentfetch.Parse(src)
	if err != nil {
		fmt.Fprintf(eout, "failed to parse %s, %s\n", args[0], err)
		os.Exit(1)
	}

	if err := patentfetch.Process(data); err != nil {
		fmt.Fprintf(eout, "failed to process %s, %s\n", args[0], err)
		os.Exit(1)
	}
}