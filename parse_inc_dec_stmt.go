package main

import "go/ast"

func ParseIncDecStmt(stmt *ast.IncDecStmt) []string {
	var ret []string
	tok := stmt.Tok.String()
	x := ParseExpr(stmt.X)
	ret = append(ret, x + tok + ";")
	return ret
}
