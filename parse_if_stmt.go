package main

import "go/ast"

func ParseIfStmt(ifStmt *ast.IfStmt, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	var initStmt []string
	var elseStmt []string

	if ifStmt.Init != nil {
		ret = append(ret, "{")
		// add init statment in '{}'
		initStmt = ParseStmt(&ifStmt.Init, objectTypeMap)
	}
	cond := ParseExpr(ifStmt.Cond)
	body := ParseBlockStmt(ifStmt.Body, objectTypeMap)

	ret = append(ret, initStmt...)
	ret = append(ret, "if (" + cond + ")")
	ret = append(ret, "{")
	ret = append(ret, body...)
	ret = append(ret, "}")
	if ifStmt.Else != nil {
		elseStmt = ParseStmt(&ifStmt.Else, objectTypeMap)
		ret = append(ret, "else " )
		ret = append(ret, "{")
		ret = append(ret, elseStmt...)
		ret = append(ret, "}")
	}

	if ifStmt.Init != nil {
		ret = append(ret, "}")
	}
	return ret
}