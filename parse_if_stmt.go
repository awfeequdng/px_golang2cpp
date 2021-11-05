package main

import "go/ast"

func ParseIfStmt(if_stmt *ast.IfStmt) []string {
	var ret []string
	var initStmt []string
	var elseStmt []string

	if if_stmt.Init != nil {
		ret = append(ret, "{")
		// add init statment in '{}'
		initStmt = ParseStmt(&if_stmt.Init)
	}
	cond := ParseExpr(if_stmt.Cond)
	body := ParseBlockStmt(if_stmt.Body)

	ret = append(ret, initStmt...)
	ret = append(ret, "if (" + cond + ")")
	ret = append(ret, body...)
	if if_stmt.Else != nil {
		elseStmt = ParseStmt(&if_stmt.Else)
		ret = append(ret, "else " )
		ret = append(ret, elseStmt...)
	}


	if if_stmt.Init != nil {
		ret = append(ret, "}")
	}
	return ret
}