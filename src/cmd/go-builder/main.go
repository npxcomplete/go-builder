//go:generate go-bindata -o ./bindata.go -pkg main ./builderTemplate.txt
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

func debugRelativePath() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		log.Println(f.Name())
	}

}

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	//debugRelativePath()

	var structName = flag.String("struct", "", "")
	var interfaceName = flag.String("interface", "", "")
	var packageName = flag.String("package", "", "")
	var inputFilePath = flag.String("input", "", "")
	flag.Parse()
	str := generate(*structName, *interfaceName, *packageName, *inputFilePath)

	dirName := path.Dir(*inputFilePath)
	baseName := path.Base(*inputFilePath)
	ext := path.Ext(*inputFilePath)
	fileName := strings.Replace(baseName, ext, "", 1)
	outPath := dirName + "/" + fileName + "_builder_generated"+ ext
	ioutil.WriteFile(outPath, []byte(str), 0755)
}

func generate(structName, interfaceName, packageName, inputFilePath string) string {
	fileset := token.NewFileSet()
	astDat, err := parser.ParseFile(fileset, inputFilePath, nil, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	for _, decl := range astDat.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if s.Name.Name == structName {
						switch structType := s.Type.(type) {
						case *ast.StructType:
							g := generatorFor(packageName, interfaceName, structName, structType)
							buf := &bytes.Buffer{}
							writeFile(buf, g)
							return buf.String()
						default:
							log.Fatalf("%s was not a struct.", structName)
						}
					}
				default:
					continue
				}
			}
		default:
			continue
		}
	}

	log.Fatal("struct not found")
	return ""
}

func writeFile(w io.Writer, g Generator) {
	templateFile, err := Asset("builderTemplate.txt")
	if err != nil {
		log.Fatal(err)
	}
	t := template.New("builder")
	t, err = t.Parse(string(templateFile))
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, g)
}

type BuilderField struct {
	StructAttribute  string
	BuilderAttribute string
	BuilderMethod    string
	StructMethod     string
	TypeName         string
}

type Generator struct {
	Fields []BuilderField
	IName  string
	SName  string
	PName  string
}

func generatorFor(pkg string, inf string, stct string, s *ast.StructType) (g Generator) {
	g = Generator{
		Fields: make([]BuilderField, 0, 8),
		IName:  inf,
		SName:  stct,
		PName:  pkg,
	}
	for _, f := range s.Fields.List {
		var field BuilderField
		field.BuilderAttribute = builderName(f)
		field.StructAttribute = structName(f)
		field.BuilderMethod = methodName(field.BuilderAttribute)
		field.StructMethod = methodName(field.StructAttribute)
		field.TypeName = typeExpToString(f.Type)
		g.Fields = append(g.Fields, field)
	}
	return
}

func typeExpToString(exp ast.Expr) string {
	switch x := exp.(type) {
	case *ast.Ident:
		return x.Name
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", typeExpToString(x.Key), typeExpToString(x.Value))
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", typeExpToString(x.Elt))
	case *ast.ChanType:
		return fmt.Sprintf("chan %s", typeExpToString(x.Value))
	case *ast.FuncType:
		buf := &bytes.Buffer{}
		buf.WriteString("func(")
		for j, p := range x.Params.List {
			if j != 0 {
				buf.WriteString(",")
			}
			for i, n := range p.Names {
				if i != 0 {
					buf.WriteString(",")
				}
				buf.WriteString(n.Name)
			}
			buf.WriteString(typeExpToString(p.Type))
		}
		buf.WriteString(")")
		for j, p := range x.Results.List {
			if j != 0 {
				buf.WriteString(",")
			}
			for i, n := range p.Names {
				if i != 0 {
					buf.WriteString(",")
				}
				buf.WriteString(n.Name)
			}
			buf.WriteString(typeExpToString(p.Type))
		}
		return buf.String()
	default:
		log.Fatalf("Type was %s", reflect.TypeOf(exp).String())
		return ""
	}
}

func methodName(str string) string {
	rs := []rune(str)
	rs[0] = unicode.ToUpper(rs[0])
	return string(rs)
}

func structName(f *ast.Field) string {
	return f.Names[0].Name
}

var r *regexp.Regexp

func init() {
	var err error
	r, err = regexp.Compile("`builder:\"(.*)\"`")
	if err != nil {
		panic(err)
	}
}
func builderName(f *ast.Field) string {
	if f.Tag != nil {
		submatch := r.FindAllStringSubmatch(f.Tag.Value, 1)
		return submatch[0][1]
	} else if len(f.Names) > 0 {
		return structName(f)
	}
	return structName(f)
}
