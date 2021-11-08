package main

import "go/ast"

func ParseIndexExpr(indexExpr *ast.IndexExpr) string {
	x := ParseExpr(indexExpr.X)
	index := ParseExpr(indexExpr.Index)
	return x + "[" + index + "]"
}
