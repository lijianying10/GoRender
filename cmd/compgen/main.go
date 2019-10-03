package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var compTmpl = `
package {{ .PackageName }}

import (
	"github.com/lijianying10/GR/gorender"
)

// {{ .CompName }} a component
type {{ .CompName }} struct {
    gorender.Component

    // TODO: add component status
}

func ({{ .RecvVerName }} *{{ .CompName }}) constructor() {
	// TODO: init comp
}

func ({{ .RecvVerName }} *{{ .CompName }}) destructor() {
	// TODO: component defer
}

func ({{ .RecvVerName }} *{{ .CompName }}) render(tid string) string {
    // TODO: component render body
	return ""
}
`[1:]

var compGenTmpl = `
// Code generated by compgen
package {{ .PackageName }}

import (
    "github.com/google/uuid"
    "github.com/lijianying10/GR/gorender"
)

// Render component Render
func ({{ .RecvVerName }} *{{ .CompName }}) Render(tid string) string {
    {{ .RecvVerName }}.AppendRenderTraceID(tid)
    return {{ .RecvVerName }}.render(tid)
}

// Destructor Comp DoDestructor
func ({{ .RecvVerName }} *{{ .CompName }}) Destructor() {
    for _, comp := range {{ .RecvVerName }}.SubComponents {
        comp.Destructor()
    }
    {{ .RecvVerName }}.destructor()
}

// Constructor Comp DoDestructor
func ({{ .RecvVerName }} *{{ .CompName }}) Constructor(grender *gorender.GoRender) {
    {{ .RecvVerName }}.Runtime = grender
    {{ .RecvVerName }}.UUID = uuid.New().String()
    {{ .RecvVerName }}.constructor()
}
`[1:]

// Input input
type Input struct {
	CompName    string
	PackageName string
	RecvVerName string
}

func main() {
	var compName, packageName, recvVarName string

	flag.StringVar(&compName, "n", "", "Component name")
	flag.StringVar(&packageName, "p", "", "Package name")
	flag.StringVar(&recvVarName, "r", "", "Recv var name")
	flag.Parse()

	fmt.Println("Createing comp: ", recvVarName, packageName, compName)
	if compName == "" || packageName == "" || recvVarName == "" {
		fmt.Println("error must fill all flag")
		fmt.Println("example: compgen -n Timer -p main -r t")
		os.Exit(1)
	}

	// Checkfile exist first
	checkFileName(compName)

	input := Input{
		CompName:    compName,
		PackageName: packageName,
		RecvVerName: recvVarName,
	}
	writeToFile(compName+".go", compTmpl, input)
	writeToFile(compName+"_gen.go", compGenTmpl, input)
}

func writeToFile(codeFileName, tmpl string, input Input) {
	tmp, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		fmt.Println("comp template error", err.Error())
		os.Exit(1)
	}

	codeFile, err := os.Create(codeFileName)
	if err != nil {
		fmt.Println("create file error: ", err.Error(), codeFileName)
		os.Exit(1)
	}

	err = tmp.Execute(codeFile, input)
	if err != nil {
		fmt.Println("tmpl exec: ", err.Error(), codeFileName)
		os.Exit(1)
	}
	codeFile.Close()
}

func checkFile(fname string) {
	_, err := os.Stat(fname)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		return
	}
	fmt.Println("file already exist")
	os.Exit(1)
}

func checkFileName(compName string) {
	checkFile(compName + ".go")
	checkFile(compName + "_gen.go")
}
