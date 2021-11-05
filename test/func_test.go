package main

import (
	"go/ast"
	"strings"
)

func ParseFuncDecl(decl *ast.FuncDecl) []string {
	var ret []string

	name := decl.Name.Name
	func_type := decl.Type
	params := ParseFieldList(func_type.Params)
	results := ParseFieldList(func_type.Results)
	if len(results) == 0 {
		ret = append(ret, "void " + name + "(" + strings.Join(params, " ") + ");")
	}
	return ret
}

func ParseFuncDecl1(decl *ast.FuncDecl) (string, string) {
	var ret []string
	var v1, v2, v3 string
	v1, v2, v3 = GetThreeValue()

	return ret
}


// func ParseFuncDecl2(decl *ast.FuncDecl) (string, string, int, float64) {
// 	var ret []string

// 	return ret
// }
