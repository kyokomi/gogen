package main

import (
	"flag"
	"fmt"
	"os"

	"text/template"

	"github.com/kyokomi/gogen"
)

var (
	// command line flags
	out  string // output file
	file string // input file (or directory)
	pkg  string // output package name

)

var (
	baseTemplateText = `
// Sample sample code
func ({{.Varname}} *{{.Value.Struct.Name}}) Sample() {
	{{range .Value.Struct.Fields}}
	fmt.Println({{.FieldElem.Varname}})
	{{end}}
}
`
	sampleTemplate = template.Must(template.New("base").Parse(baseTemplateText))
)

func init() {
	flag.StringVar(&out, "o", "", "output file")
	flag.StringVar(&file, "file", os.Getenv("GOFILE"), "input file")
	flag.StringVar(&pkg, "pkg", os.Getenv("GOPACKAGE"), "output package")
}

func main() {
	flag.Parse()

	if file == "" {
		fmt.Println("No file to parse.")
		os.Exit(1)
	}

	g := gogen.NewGenerator(file, out, pkg, "fmt")

	err := g.DoAllTemplate(sampleTemplate)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
