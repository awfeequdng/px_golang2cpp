package main

import (
	"go/ast"
	"strings"
)


// FindColumnInfo finds ColumnInfo in cols by name.
func FindColumnInfo(cols []*ColumnInfo, name string) *ColumnInfo {
	name = strings.ToLower(name)
	for _, col := range cols {
		if col.Name.L == name {
			return col
		}
	}

	return nil
}

// func ParseFuncDecl(decl *ast.FuncDecl) []string {
// 	var ret []string

// 	name := decl.Name.Name
// 	func_type := decl.Type
// 	params := ParseFieldList(func_type.Params)
// 	results := ParseFieldList(func_type.Results)
// 	if len(results) == 0 {
// 		ret = append(ret, "void " + name + "(" + strings.Join(params, " ") + ");")
// 	}
// 	return ret
// }

// func ParseFuncDecl1(decl *ast.FuncDecl) (string, string) {
// 	var ret []string
// 	var v1, v2, v3 string
// 	v1, v2, v3 = GetThreeValue()

// 	if _, st := GenInit(); st {

// 	}
// 	return ret
// }


// func ParseFuncDecl2(decl *ast.FuncDecl) (string, string, int, float64) {
// 	var ret []string

// 	return ret
// }
