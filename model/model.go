package model

import (
	"gopkg.in/yaml.v2"
	"log"
)

type Chart struct {
	ReleaseName string
	Application Application
}

type Application struct {
	Name				string			`yaml:"name"`
	Tag					string			`yaml:"tag"`
	Kind				string			`yaml:"kind"`
	Owner				string			`yaml:"owner"`
	Replicas			string			`yaml:"replicas"`
	Cpu					string			`yaml:"cpu"`
	Memory				string			`yaml:"memory"`
	LivenessProbe		string			`yaml:"livenessProbe"`
	ReadinessProbe		string			`yaml:"readinessProbe"`
	ServicePort			string			`yaml:"servicePort",default:"80"`
	Env					[]EnvVars		`yaml:"env"`
}

type EnvVars struct {
	Name		string 			`yaml:"name"`
	Value		string 			`yaml:"value"`
}

func (app *Application) UnmarshalYaml(content *[]byte) {
	err := yaml.Unmarshal(*content, app)
	if err != nil {
		log.Fatal("error parsing it ", err)
	}
}