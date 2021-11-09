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
			if _, ok := elt.(* ast.KeyValueExpr); ok {
				log.Fatal("not support key-value in array type now")
			}
			vals = append(vals, ParseExpr(elt) + ",")
		}
		vals = append(vals, "}")
		ret = append(ret, vals...)
	} else {
		ret = append(ret, ParseExpr(compLit.Type))
		vals = append(vals, "{")
		for idx, elt := range compLit.Elts {
			if idx == 0 {
				vals = append(vals, ParseExpr(elt))
			} else {
				vals = append(vals, ", " + ParseExpr(elt))
			}
		}
		vals = append(vals, "}")
		ret = append(ret, vals...)
		//p := GetProgram()
		//log.Fatal("not implemented var type: " + ParseExpr(compLit.Type) + ", pos = " + p.fset.Position(compLit.Pos()).String())
	}

	return strings.Join(ret, "\n")
}