package main

import "go/ast"

func ParseStarExpr(starExpr *ast.StarExpr) string {
	x := starExpr.X
	// todo: starExpr may be a dereference operation

	includeFileMap["std::shared_ptr"] = "memory"
	return "std::shared_ptr<" + ParseExpr(x) + ">"
}