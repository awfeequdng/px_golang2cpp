package main

import "go/ast"

func ParseExprStmt(exprStmt *ast.ExprStmt) []string {
	var ret []string
	ret = append(ret, ParseExpr(exprStmt.X) + ";")
	return ret
}
