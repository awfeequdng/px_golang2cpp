package main

import "go/ast"

func ParseUnaryExpr(unary_expr *ast.UnaryExpr) string {
	op := unary_expr.Op.String()
	val := ParseExpr(unary_expr.X)
	return "(" + op + val + ")"
}