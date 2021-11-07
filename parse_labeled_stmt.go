package main

import (
	"go/ast"
	"log"
)

func ParseLabeledStmt(stmt *ast.LabeledStmt, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	if stmt.Label == nil {
		log.Fatalln("label name can not empty")
	}

	ret = append(ret, stmt.Label.Name + ":")
	ret = append(ret, ParseStmt(&stmt.Stmt, objectTypeMap)...)

	return ret
}
