package main

import "go/ast"

func ParseSwitchStmt(switch_stmt *ast.SwitchStmt) []string {
	var ret []string
	tag := ParseExpr(switch_stmt.Tag)
	body := ParseBlockStmt(switch_stmt.Body)
	if switch_stmt.Init != nil {
		ret = append(ret, "{")
		ret = append(ret, ParseStmt(&switch_stmt.Init)...)
	}
	ret = append(ret, "switch(" + tag + ")")
	ret = append(ret, body...)

	if switch_stmt.Init != nil {
		ret = append(ret, "}")
	}

	return ret
}
