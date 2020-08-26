package main

import (
	"fmt"
	"os"

	airgo "github.com/mottaquikarim/airgo/airgo"
	flag "github.com/namsral/flag"
	log "github.com/sirupsen/logrus"
)

// these are the key variables that store the input
// required to make this service run
var (
	fs    = flag.NewFlagSetWithEnvPrefix(os.Args[0], "STA9760F2020", 0)
	debug = fs.Bool("debug", false, "Turn on debugging output")
)

// helper function to hint to the user
// what the required key variables are
// (in this case, apiKey and baseId)
func usage() {
	fmt.Println("Usage: ./service.go [flags]")
	fs.PrintDefaults()
}

func main() {

	// first, parse the input and extract
	// apiKey and baseId variables
	fs.Usage = usage
	fs.Parse(os.Args[1:])

	if err := airgo.RebuildSite(airgo.RebuildSiteOpts{
		Datadir:       "data",
		Template:      "templates",
		Output:        "docs",
		TemplateNames: []string{"layout"},
	}); err != nil {
		log.Fatal(err)
	}

	if *debug {
		airgo.DevServer("docs", 80)
	}
}
