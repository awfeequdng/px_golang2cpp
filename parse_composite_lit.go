package main

import (
	"go/ast"
	"log"
)

// return: type and values
func ParseCompositeLit(compLit *ast.CompositeLit) (typ int, tname string, vals []string) {
	if mapType, ok := compLit.Type.(*ast.MapType); ok {
		typ = VAR_MAP_TYPE
		tname, vals = ParseMapType(mapType, compLit.Elts)
	} else if arrayType, ok := compLit.Type.(*ast.ArrayType); ok {
		typ = VAR_ARRAY_TYPE
		tname, vals = ParseArrayType(arrayType, compLit.Elts)
	} else {
		log.Fatal("not implemented var type")
	}
	return typ, tname, vals
}