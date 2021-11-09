package main

import "go/ast"

func ParseDeferStmt(stmt *ast.DeferStmt, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	ret = append(ret, "DEFER(")
	ret = append(ret, ParseExpr(stmt.Call) + ";")
	ret = append(ret, ");")
	return ret
}
