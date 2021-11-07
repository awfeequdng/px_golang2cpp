package main

import "go/ast"

func ParseForStmt(stmt *ast.ForStmt, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	if stmt.Init != nil {
		ret = append(ret, "{")
		ret = append(ret, ParseStmt(&stmt.Init, objectTypeMap)...)
	}
	ret = append(ret, "for (;")
	if stmt.Cond != nil {
		ret = append(ret, ParseExpr(stmt.Cond))
	}
	ret = append(ret, ";")

	if stmt.Post != nil {
		ret = append(ret, ParseStmt(&stmt.Post, objectTypeMap)...)
	}
	ret = append(ret, ")")
	if stmt.Body == nil {
		ret = append(ret, " {}")
	} else {
		ret = append(ret, ParseBlockStmt(stmt.Body, objectTypeMap)...)
	}

	if stmt.Init != nil {
		ret = append(ret, "}")
	}
	return ret
}
