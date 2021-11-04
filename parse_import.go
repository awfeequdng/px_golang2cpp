package main

import (
	"go/ast"
)

func ParseImport(decl *ast.GenDecl) []string {
	var ret []string
	var names []string
	for _, spec := range decl.Specs {
		if is, ok := spec.(*ast.ImportSpec); ok {
			names = append(names, is.Path.Value)
		}
	}
	for _, name := range names {
		ret = append(ret, "#include " + name);
	}
	return ret
}