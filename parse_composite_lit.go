package main

import (
	"go/ast"
	"log"
)

// return: type and values
func ParseCompositeLit(compLit *ast.CompositeLit) (typ int, tname string, vals []string) {
	if mapType, ok := compLit.Type.(*ast.MapType); ok {
		typ = VAR_MAP_TYPE
		tname = ParseMapType(mapType)
		vals = append(vals, "{")
		for _, elt := range compLit.Elts {
			vals = append(vals, ParseExpr(elt) + ",")
		}
		vals = append(vals, "}")
	} else if arrayType, ok := compLit.Type.(*ast.ArrayType); ok {
		typ = VAR_ARRAY_TYPE
		tname = ParseArrayType(arrayType)
		vals = append(vals, "{")
		for _, elt := range compLit.Elts {
			vals = append(vals, ParseExpr(elt) + ",")
		}
		vals = append(vals, "}")
	} else {
		log.Fatal("not implemented var type")
	}
	return typ, tname, vals
}