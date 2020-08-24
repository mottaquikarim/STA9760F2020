package main

import (
	airgo "github.com/mottaquikarim/airgo/airgo"
	log "github.com/sirupsen/logrus"
)

func main() {

	if err := airgo.RebuildSite(airgo.RebuildSiteOpts{
		Datadir:       "data",
		Template:      "templates",
		Output:        "docs",
		TemplateNames: []string{"layout"},
	}); err != nil {
		log.Fatal(err)
	}

	airgo.DevServer("docs", 80)
}
