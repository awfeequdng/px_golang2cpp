package main

import (
	"go/ast"
	"go/token"
	"log"
)

func ParseConst(decl *ast.GenDecl) []string {
	var ret []string
	for _, spec := range decl.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			var names []string
			for _, name := range vs.Names {
				names = append(names, name.Name)
			}
			var values []string

			for _, value := range vs.Values {
				if v, ok := value.(*ast.BasicLit); ok {
					switch v.Kind {
					case token.STRING:
						values = append(values, v.Value)
						break
					case token.INT:
						values = append(values, v.Value)
						break
					case token.FLOAT:
						values = append(values, v.Value)
						break
					case token.CHAR:
						values = append(values, v.Value)
						break
					default:
						log.Fatal("invalid basicLit type")
					}
				}
			}
			if len(names) != len(values) {
				if len(values) == 1 {
					for _, name := range names {
						ret = append(ret, "#define " + name + " " + values[0])
					}
				} else {
					log.Fatal("invalid names size and values size")
				}
			} else {
				for i, name := range names {
					ret = append(ret, "#define " + name + " " + values[i])
				}
			}
		}
	}
	return ret
}