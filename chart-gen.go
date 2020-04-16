package main

import (
	"fmt"
	"github.com/sujanks/chart-gen/process"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing release name")
		os.Exit(0)
	}
	releaseName := os.Args[1]
	configFile := "sample-app.yaml"

	process.Generator(releaseName, configFile)
	fmt.Println("Helm chart generated under folder tmp for release: ", releaseName)
}
