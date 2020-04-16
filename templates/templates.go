package templates

import (
	"fmt"
	"strings"
	"text/template"
)

var ChartTemplate = `apiVersion: v1
description: A Helm chart for Kubernetes {{ .ReleaseName }}
name: {{ .ReleaseName }}
version: {{ .Application.Tag }}
`

var ServiceAccountTemplate = `apiVersion: v1
kind: ServiceAccount
metadata:
  name: {{ .ReleaseName }}-{{ .Application.Name }}
  labels:
    app.kubernetes.io/name: {{ .Application.Name }}
	app.kubernetes.io/instance: {{ .ReleaseName }}
	app.kubernetes.io/version: {{ .Application.Tag }}
`

var ServiceTemplate = `apiVersion: v1
kind: Service
metadata:
	name: {{ .ReleaseName }}-{{ .Application.Name }}
	labels:
		app.kubernetes.io/name: {{ .Application.Name }}
		app.kubernetes.io/instance: {{ .ReleaseName }}
		app.kubernetes.io/version: {{ .Application.Tag }}
spec:
	type: ClusterIP
	ports:
	- name: http
	  port: 80
      targetPort: http
      protocol: TCP
	selector:
		app.kubernetes.io/name: {{ .Application.Name }}
    	app.kubernetes.io/instance: {{ .ReleaseName }}
`

var DeploymentTemplate = `apiVersion: v1
kind: Deployment
metadata:
	name: {{ .ReleaseName }}-{{ .Application.Name }}
	labels:
		app.kubernetes.io/name: {{ .Application.Name }}
		app.kubernetes.io/instance: {{ .ReleaseName }}
		app.kubernetes.io/version: {{ .Application.Tag }}
spec:
	replicas: {{ .Application.Replicas }}
	selector:
		matchLabels:
			app.kubernetes.io/name: {{ .Application.Name }}
    		app.kubernetes.io/instance: {{ .ReleaseName }}
    template:
		metadata:
			labels:
				app.kubernetes.io/name: {{ .Application.Name }}
    			app.kubernetes.io/instance: {{ .ReleaseName }}
			annotations:
		spec:
			serviceAccountName: {{ .ReleaseName }}-{{ .Application.Name }}
			securityContext: 
			containers:
				- name: {{ .Application.Name }}
                  image: {{ .Application.Name}}:{{ .Application.Tag}}
                  imagePullPolicy: IfNotPresent

                ports:
				  - name: http
                    containerPort: 8080
                    protocol: TCP

                livenessProbe:
                   httpGet:
                     path: {{ .Application.LivenessProbe }}
                     port: http
                   initialDelaySeconds: 100
                   timeoutSeconds: 100                
                readinessProbe:
                   httpGet:
                     path: {{ .Application.ReadinessProbe }}
                     port: http
                   initialDelaySeconds: 100
                   timeoutSeconds: 100

                resources:
                  limits:
                    cpu: "{{ .Application.Cpu }}"
                    memory:  "{{ .Application.Memory }}"   
                  requests:
                    cpu:  "{{ .Application.Cpu }}"
                    memory:  "{{ .Application.Memory }}"

                env:{{ range $i, $env := .Application.Env }}
                 - name: "{{ $env.Name | ToUpper }}"
                   value: "{{ $env.Value }}"
                 {{end}}
			affinity:
			  nodeSelector:
			  tolerations:
`

//LoadTemplates parse static template to helm chart
func LoadTemplates(tName string) *template.Template {
	switch tName {
	case "ChartTemplate":
		return getTemplate("Chart.yaml", ChartTemplate)
	case "DeploymentTemplate":
		return getTemplate("deployment.yaml", DeploymentTemplate)
	case "ServiceTemplate":
		return getTemplate("service.yaml", ServiceTemplate)
    case "ServiceAccountTemplate":
		return getTemplate("serviceaccount.yaml", ServiceAccountTemplate)
	}
	return nil
}

func getTemplate(name string, serviceTemplate string) *template.Template {
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	tmpl, err := template.New(name).Funcs(funcMap).Parse(serviceTemplate)
	if err != nil {
		fmt.Println("Error parsing ", err)
	}
	return tmpl
}
