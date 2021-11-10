package main

import (
	"go/ast"
)

func ParseCaseClause(caseClause *ast.CaseClause, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	var caseList []string
	for _, l := range caseClause.List {
		caseList = append(caseList, ParseExpr(l))
	}
	var body []string
	//var bodyCnt = 0
	body = append(body, "{")
	for _, b := range caseClause.Body {
		body = append(body, ParseStmt(&b, objectTypeMap)...)
		//bodyCnt++
	}
	body = append(body, "}")

	//if bodyCnt > 1 {
	//	log.Fatal("body count is : " + strconv.Itoa(bodyCnt))
	//}
	for _, c := range caseList {
		ret = append(ret, "case " + c + ":")
	}
	if len(caseList) == 0 {
		// default
		ret = append(ret, "default:")
	}
	ret = append(ret, body...)
	ret = append(ret, "break;")

	return ret
}
