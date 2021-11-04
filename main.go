package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const version = "golang2cpp 0.1"

func main() {
	inputFilename := ""
	outputFilename := ""
	if len(os.Args) > 1 {
		if os.Args[1] == "--version" {
			fmt.Println(version)
			return
		} else if os.Args[1] == "--help" {
			fmt.Println("supported arguments:")
			fmt.Println(" a .go file as the first argument")
			fmt.Println("supported options:")
			fmt.Println(" -o : Format with clang format")
			fmt.Println(" -O : Don't format with clang format")
			return
		}
		inputFilename = os.Args[1]
	}

	if len(os.Args) > 3 {
		if os.Args[2] != "-o" {
			log.Fatal("The second argument must be -o (format sources with clang-format) or -O (don't format sources with clang-format)")
		}
		outputFilename = os.Args[3]
	}

	var sourceData []byte
	var err error
	if inputFilename != "" {
		sourceData, err = ioutil.ReadFile(inputFilename)
	} else {
		sourceData, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		log.Fatal(err)
	}

	cppSource := golang2cpp(inputFilename, string(sourceData))

	if outputFilename != "" {
		err = ioutil.WriteFile(outputFilename, []byte(cppSource), 0755)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(cppSource)
	}
}

func parseGolang(f *ast.File) []string {
	var ret []string
	for _, decl := range f.Decls {
		if g, ok := decl.(*ast.GenDecl); ok {
			switch g.Tok {
			case token.IMPORT:
				ret = append(ret, ParseImport(g)...)
			case token.CONST:
				ret = append(ret, ParseConst(g)...)
			case token.VAR:
				ret = append(ret, ParseVar(g)...)
				break
			case token.TYPE:
				ret = append(ret, ParseType(g)...)
				break
			case token.FUNC:
				ret = append(ret, ParseFunc(g)...)
				break
			default:
				log.Fatal("invalid genDecl token type")
			}
		}
	}

	return ret
}

func golang2cpp(file, source string) string {
	prog := NewProgram(map[string]string {
		file: source,
	})
	_, f, err := prog.LoadPackage(file)
	if err != nil {
		log.Fatal(err)
	}

	// ast.Print(prog.fset, pkg)
	buf := new(bytes.Buffer)
	ast.Fprint(buf, prog.fset, f, ast.NotNilFilter)
	print(buf.String())

	ret := parseGolang(f)
	return strings.Join(ret, "\n")
}

type Program struct {
	fs   map[string]string
	ast  map[string]*ast.File
	pkgs map[string]*types.Package
	fset *token.FileSet
}

func NewProgram(fs map[string]string) *Program {
	return &Program{
		fs:   fs,
		ast:  make(map[string]*ast.File),
		pkgs: make(map[string]*types.Package),
		fset: token.NewFileSet(),
	}
}

func (p *Program) LoadPackage(path string) (pkg *types.Package, f *ast.File, err error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, p.ast[path], nil
	}

	f, err = parser.ParseFile(p.fset, path, p.fs[path], parser.AllErrors)
	if err != nil {
		return nil, nil, err
	}

	// conf := types.Config{Importer: importer.Default()}
	// pkg, err = conf.Check(path, p.fset, []*ast.File{f}, nil)
	// if err != nil {
		// return nil, nil, err
	// }

	p.ast[path] = f
	p.pkgs[path] = pkg
	return pkg, f, nil
}

func (p *Program) Import(path string) (*types.Package, error) {
	if pkg, ok := p.pkgs[path]; ok {
		return pkg, nil
	}
	pkg, _, err := p.LoadPackage(path)
	return pkg, err
}
