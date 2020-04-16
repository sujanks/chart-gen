package process

import (
	"github.com/sujanks/chart-gen/model"
	"github.com/sujanks/chart-gen/templates"
	"log"
	"os"
)

func Generator(rName string, configFile string) {

	appYaml := Read(configFile)

	app := model.Application{}
	app.UnmarshalYaml(appYaml)

	chart := model.Chart{
		ReleaseName: rName,
		Application: app,
	}
	run(&chart)
}

func run(chart *model.Chart) {
	tmplArray := []string{"ChartTemplate", "ServiceTemplate", "DeploymentTemplate", "ServiceAccountTemplate"}

	dir := "tmp/templates"
	os.MkdirAll(dir, os.ModePerm)
	os.Chdir(dir)

	for _, tName := range tmplArray {
		tmpl := templates.LoadTemplates(tName)

		file, er := os.Create(tmpl.Name())
		if er != nil {
			log.Fatal("error ", er)
		}

		err := tmpl.Execute(file, chart)
		if err != nil {
			log.Fatal("error ", err)
		}
	}

}