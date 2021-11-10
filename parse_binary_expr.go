package main

import "go/ast"

func ParseBinaryExpr(binaryExpr *ast.BinaryExpr) string {
	var ret string
	x := binaryExpr.X
	op := binaryExpr.Op
	y := binaryExpr.Y
	ret += ParseExpr(x)
	ret += op.String()
	ret += ParseExpr(y)
	return ret
}