package main

import (
	"go/ast"
	"strings"
)

var globalObjectMap map[string]string

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