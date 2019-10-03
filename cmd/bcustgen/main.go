package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

var tmpl = `
package {{ .PackageName }}

// {{ .BcustName }} a component
type {{ .BcustName }} struct {
    channels     map[string]chan {{ .TargetType }}
    channelsLock sync.RWMutex
}

// New{{ .BcustName }} new {{ .BcustName }}
func New{{ .BcustName }}() *{{ .BcustName }} {
    return &{{ .BcustName }} {
        channels:make(map[string]chan {{ .TargetType }}),
    }
}

// Bind bind a new client
func (bcust *{{ .BcustName }}) Bind(id string) chan {{ .TargetType }} {
    bcust.channelsLock.Lock()
    defer bcust.channelsLock.Unlock()
    bcust.channels[id] = make(chan {{ .TargetType }}, 10)
    return bcust.channels[id]
}

// Unbind a client
func (bcust *{{ .BcustName }}) Unbind(id string) {
    bcust.channelsLock.Lock()
    defer bcust.channelsLock.Unlock()
    close(bcust.channels[id])
    delete(bcust.channels, id)
}

// BCust a event
func (bcust *{{ .BcustName }}) BCust(event {{ .TargetType }}) {
    bcust.channelsLock.RLock()
    defer bcust.channelsLock.RUnlock()
    for _, channel := range bcust.channels {
        channel <- event
    }
}
`[1:]

// Input input
type Input struct {
	PackageName, BcustName, TargetType string
}

func main() {
	var input Input

	flag.StringVar(&input.BcustName, "b", "", "broadcust name")
	flag.StringVar(&input.PackageName, "p", "", "Package name")
	flag.StringVar(&input.TargetType, "t", "", "target broadcust type")
	flag.Parse()

	fmt.Println("Createing bcust module : ", input)
	if input.BcustName == "" || input.PackageName == "" || input.TargetType == "" {
		fmt.Println("error must fill all flag")
		fmt.Println("example: bcustgen -b DataBcust -p main -t 'struct{}'")
		os.Exit(1)
	}

	// Checkfile exist first
	checkFileName(input.BcustName)

	writeToFile(input.BcustName+".go", tmpl, input)
}

func writeToFile(codeFileName, tmpl string, input Input) {
	tmp, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		fmt.Println("module template error", err.Error())
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

func checkFileName(fname string) {
	checkFile(fname + ".go")
}
