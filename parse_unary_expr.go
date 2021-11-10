package main

import "go/ast"

func ParseUnaryExpr(unaryExpr *ast.UnaryExpr) string {
	op := unaryExpr.Op.String()
	val := ParseExpr(unaryExpr.X)
	return  op + val
}