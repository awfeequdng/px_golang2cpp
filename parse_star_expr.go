package main

import "go/ast"

func ParseStarExpr(star_expr *ast.StarExpr) string {
	x := star_expr.X
	// todo: starExpr may be a dereference operation

	includeFileMap["std::shared_ptr"] = "memory"
	return "std::shared_ptr<" + ParseExpr(x) + ">"
}