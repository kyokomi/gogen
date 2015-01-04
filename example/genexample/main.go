package main

import (
	"flag"
	"fmt"
	"os"

	"text/template"

	"github.com/kyokomi/gogen"
	"github.com/ttacon/chalk"
)

var (
	// command line flags
	out  string // output file
	file string // input file (or directory)
	pkg  string // output package name

)

var (
	baseTemplateText = `
// DecodeMsg implements the msgp.Decodable interface
func ({{.Varname}} *{{.Value.Struct.Name}}) DecodeMsg(dc *msgp.Reader) (err error) {
	{{if not .Value.Struct.AsTuple}}var field []byte; _ = field{{end}}
	{{if .Value.Struct.AsTuple}}
	{ {{/* tuples get their own blocks so that we don't clobber 'ssz'*/}}
		var ssz uint32
		ssz, err = dc.ReadArrayHeader()
		if err != nil {
			return
		}
		if ssz != {{len .Value.Struct.Fields}} {
			err = msgp.ArrayError{Wanted: {{len .Value.Struct.Fields}}, Got: ssz}
			return
		}
		{{range .Value.Struct.Fields}}{{template "ElemTempl" .Value.Struct.FieldElem}}{{end}}
	}
	{{else}}
	var isz uint32
	isz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for xplz:=uint32(0); xplz<isz; xplz++ {
		field, err = dc.ReadMapKey(field)
		if err != nil {
			return
		}

		// TODO:
	}
	{{end}}

	return
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
		fmt.Println(chalk.Red.Color("No file to parse."))
		os.Exit(1)
	}

	g := gogen.NewGenerator(file, out, pkg, "github.com/philhofer/msgp/msgp")

	err := g.DoAllTemplate(sampleTemplate)

	if err != nil {
		fmt.Println(chalk.Red.Color(err.Error()))
		os.Exit(1)
	}
}
