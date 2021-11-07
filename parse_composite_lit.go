package main

import (
	"go/ast"
	"log"
	"strings"
)

func ParseCompositeLite(compLit *ast.CompositeLit) string {
	var ret []string
	var vals []string
	if mapType, ok := compLit.Type.(*ast.MapType); ok {
		ret = append(ret, ParseMapType(mapType))
		vals = append(vals, "{")
		for _, elt := range compLit.Elts {
			vals = append(vals, ParseExpr(elt) + ",")
		}
		vals = append(vals, "}")
		ret = append(ret, vals...)
	} else if arrayType, ok := compLit.Type.(*ast.ArrayType); ok {
		ret = append(ret, ParseArrayType(arrayType))
		vals = append(vals, "{")
		for _, elt := range compLit.Elts {
			vals = append(vals, ParseExpr(elt) + ",")
		}
		vals = append(vals, "}")
		ret = append(ret, vals...)
	} else {
		log.Fatal("not implemented var type")
	}

	return strings.Join(ret, "\n")
}