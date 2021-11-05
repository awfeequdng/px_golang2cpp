package main

import (
	"go/ast"
	"log"
)

// ParseNoConditionCaseClause use 'if' instead of 'case'
func ParseNoConditionCaseClause(caseClause *ast.CaseClause) []string {
	var ret []string
	var caseList []string
	for _, l := range caseClause.List {
		caseList = append(caseList, ParseExpr(l))
	}
	var body []string
	var bodyCnt = 0
	body = append(body, "{")
	for _, b := range caseClause.Body {
		body = append(body, ParseStmt(&b)...)
		body = append(body, ";")
		bodyCnt++
	}
	body = append(body, "}")

	if len(caseList) == 0 {
		// default
		ret = append(ret, " ")
	} else {
		ret = append(ret, " if (")
		for id, c := range caseList {
			if id == 0 {
				ret = append(ret, c)
			} else {
				ret = append(ret, " || " + c)
			}
		}
		ret = append(ret, ")")
	}
	ret = append(ret, body...)
	return ret
}

func ParseNoConditionSwitchStmt(switchStmt *ast.SwitchStmt) [] string {
	var ret []string
	if switchStmt.Init != nil {
		ret = append(ret, "{")
		ret = append(ret, ParseStmt(&switchStmt.Init)...)
	}

	bodyCnt := len(switchStmt.Body.List)
	for id, l := range switchStmt.Body.List {
		var body []string
		if clause, ok := l.(*ast.CaseClause); ok {
			body = append(body, ParseNoConditionCaseClause(clause)...)
		} else {
			log.Fatal("invalid case clause")
		}
		ret = append(ret, body...)
		if id != bodyCnt - 1 {
			ret = append(ret, " else ")
		}
	}
	
	if switchStmt.Init != nil {
		ret = append(ret, "}")
	}
	return ret
}

func ParseSwitchStmt(switch_stmt *ast.SwitchStmt) []string {
	var ret []string
	var tag string
	if switch_stmt.Tag == nil {
		return ParseNoConditionSwitchStmt(switch_stmt)
	}
	tag = ParseExpr(switch_stmt.Tag)
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
