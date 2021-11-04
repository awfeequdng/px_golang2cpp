package main

import "go/ast"

func ParseBinaryExpr(binary_expr *ast.BinaryExpr) string {
	var ret string
	x := binary_expr.X
	op := binary_expr.Op
	y := binary_expr.Y
	ret += "(" + ParseExpr(x)
	ret += op.String()
	ret += ParseExpr(y) + ")"
	return ret
}