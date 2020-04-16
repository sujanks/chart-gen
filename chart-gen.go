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

	appYaml := process.Read("sample-app.yaml")
	releaseName := os.Args[1]

	fmt.Printf("Preparing for release: %s", releaseName)
	process.Generator(appYaml, releaseName)
}
