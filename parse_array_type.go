package main

import "go/ast"

func ParseArrayType(arrayType *ast.ArrayType) string {
	var ret string
	tName := ParseExpr(arrayType.Elt)
	if typ, ok := typeMap[tName]; ok {
		tName = typ
	}
	ret = "std::vector<" + tName + ">"
	return ret
}