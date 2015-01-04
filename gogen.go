package gogen

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/philhofer/msgp/gen"
	"github.com/philhofer/msgp/parse"
	"github.com/ttacon/chalk"
)

type Generator struct {
	// input file (or directory)
	inFile string
	// output file
	outFile string
	pkgName string
	imports []string
}

func NewGenerator(inFile, outFile, pkgName string, imports ...string) *Generator {
	return &Generator{
		inFile:  inFile,
		outFile: outFile,
		pkgName: pkgName,
		imports: imports,
	}
}

type GenerateFunc func(w io.Writer, p *gen.Ptr) error

func (g Generator) DoAllTemplate(t *template.Template) error {
	return g.DoAll(func(w io.Writer, p *gen.Ptr) error {
		return execAndFormat(t, w, p)
	})
}

func (g Generator) DoAll(executeFunc GenerateFunc) error {
	var (
		outwr *bufio.Writer // location to write methods
	)

	fmt.Println("gopkg: ", g.pkgName, " gofile: ", g.inFile)

	newfile := createNewFileName(g.pkgName, g.outFile, g.inFile)
	fmt.Println("fileName: ", newfile)

	file, err := os.Create(newfile)
	if err != nil {
		return err
	}
	defer file.Close()

	outwr = bufio.NewWriter(file)

	err = writePkgHeader(outwr, g.pkgName)
	if err != nil {
		return err
	}

	err = writeImportHeader(outwr, g.imports)
	if err != nil {
		return err
	}

	elems, _, err := parse.GetElems(g.inFile)
	if err != nil {
		return err
	}

	for _, el := range elems {

		p, ok := el.(*gen.Ptr)
		fmt.Println(p, ok, !ok || p.Value.Type() != gen.StructType, p.Value.Type(), gen.StructType)
		if !ok || p.Value.Type() != gen.StructType {
			continue
		}

		fmt.Println(p.Value.TypeName())

		if err := executeFunc(outwr, p); err != nil {
			fmt.Println(err)
			continue
		}
	}

	err = outwr.Flush()
	if err != nil {
		return err
	}
	fmt.Print(chalk.Green.Color("\u2713\n"))

	return nil
}

// createNewFileName generateするfile名をいい感じに作成する
// output先を指定した場合は_genをつけない
func createNewFileName(pkgName string, outFile string, gofile string) string {
	var isDir bool
	if fInfo, err := os.Stat(gofile); err == nil && fInfo.IsDir() {
		isDir = true
	}

	var newfile string // new file name
	if outFile != "" {
		newfile = outFile
		if pre := strings.TrimPrefix(outFile, gofile); len(pre) > 0 &&
				!strings.HasSuffix(outFile, ".go") {
			newfile = filepath.Join(gofile, outFile)
		}
	} else {
		// small sanity check if gofile == . or dir
		// let's just stat it again, not too costly
		if isDir {
			gofile = filepath.Join(gofile, pkgName)
		}
		// new file name is old file name + _gen.go
		newfile = strings.TrimSuffix(gofile, ".go") + "_gen.go"
	}

	return newfile
}

// fileの中身作成

func writePkgHeader(w io.Writer, name string) error {
	_, err := io.WriteString(w, fmt.Sprintf("package %s\n\n", name))
	if err != nil {
		return err
	}

	// TODO: option NOTE
	//	_, err = io.WriteString(w, "// NOTE: THIS FILE WAS PRODUCED BY THE\n// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)\n// DO NOT EDIT\n\n")
	return err
}

func writeImportHeader(w io.Writer, imports []string) error {
	_, err := io.WriteString(w, "import (\n")
	if err != nil {
		return err
	}
	for _, im := range imports {
		_, err = io.WriteString(w, fmt.Sprintf("\t%q\n", im))
		if err != nil {
			return err
		}
	}
	_, err = io.WriteString(w, ")\n\n")
	return err
}

// execAndFormat executes a template and formats the output, using buf as temporary storage
func execAndFormat(t *template.Template, w io.Writer, i interface{}) error {
	buf := bytes.NewBuffer(nil)

	err := t.Execute(buf, i)
	if err != nil {
		return fmt.Errorf("template: %s", err)
	}
	bts, err := format.Source(buf.Bytes())
	if err != nil {
		w.Write(buf.Bytes())
		return fmt.Errorf("gofmt: %s", err)
	}
	_, err = w.Write(bts)
	return err
}
