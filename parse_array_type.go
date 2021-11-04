package main

import "go/ast"

func ParseArrayType(arrayType *ast.ArrayType) string {
	var ret string
	tname := ParseExpr(arrayType.Elt)
	ret = "std::vector<" + tname + ">"
	return ret
}