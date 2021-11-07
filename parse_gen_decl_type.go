package main

import (
	"go/ast"
	"log"
)

func ParseGenDeclType(decl *ast.GenDecl) []string {
	var ret []string
	for _, spec := range decl.Specs {
		var name string
		if ts, ok := spec.(*ast.TypeSpec); ok {
			name = ts.Name.Name
			typ := ParseExpr(ts.Type)
			ret = append(ret, "typedef " + typ + " " + name + ";")
		} else {
			log.Fatal("invalid spec type in ParseType function ")
		}
	}

	return ret
}