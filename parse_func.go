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
	signature := name + "(" + strings.Join(params, ",") + ")"
	if len(results) == 0 {
		ret = append(ret, "void " + signature)
	} else if len(results) == 1 {
		ret = append(ret, results[0] + " " + signature)
	} else if len(results) == 2 {
		includeFileMap["std::pair"] = "utility"
		ret = append(ret, "std::pair<" + results[0] + "," + results[1] + ">" + signature)
	} else {
		includeFileMap["std::tuple"] = "tuple"
		tuple := "std::tuple<"
		for id, r := range results {
			if id == 0 {
				tuple += r
			} else {
				tuple += ", " + r
			}
		}
		tuple += ">" + signature
		ret = append(ret, tuple)
	}
	body := ParseBlockStmt(decl.Body)
	ret = append(ret, body...)

	return ret
}