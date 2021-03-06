package main

import "go/ast"

func ParseReturnStmt(returnStmt *ast.ReturnStmt) []string {
	var exprs []string

	if returnStmt.Results == nil {
		return []string{"return;"}
	}

	for _, res := range returnStmt.Results {
		exprs = append(exprs, ParseExpr(res))
	}

	if len(exprs) == 0 {
		return []string{"return;"}
	} else if len(exprs) == 1 {
		return []string { "return " + exprs[0] + ";"}
	}

	var ret string

	ret = "return {"
	// exprs size >= 2
	for id, expr := range exprs {
		if id == 0 {
			ret += expr
		} else {
			ret += ", " + expr
		}
	}
	ret += "};"
	return []string {ret}

}