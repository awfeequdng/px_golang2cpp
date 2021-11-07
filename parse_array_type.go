package main

import "go/ast"

func ParseArrayType(arrayType *ast.ArrayType) string {
	var ret string
	tName := ParseExpr(arrayType.Elt)
	ret = "std::vector<" + tName + ">"
	return ret
}