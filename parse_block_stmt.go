package main

import (
	"go/ast"
	"log"
)

func ParseDeclStmt(declStmt *ast.DeclStmt, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	if genDecl, ok := declStmt.Decl.(*ast.GenDecl); ok {
		ret = append(ret, ParseGenDecl(genDecl, objectTypeMap)...)
	} else {
		log.Fatal("not implemented decl otherwise gendecl")
	}
	return ret
}

func ParseStmt(stmt *ast.Stmt, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	switch (*stmt).(type) {
		case *ast.BadStmt:
			log.Fatal("bad statement")
		case *ast.DeclStmt:
			ret = append(ret, ParseDeclStmt((*stmt).(*ast.DeclStmt), objectTypeMap)...)
		case *ast.EmptyStmt:
			// do nothing
		case *ast.LabeledStmt:
			ret = append(ret, ParseLabeledStmt((*stmt).(*ast.LabeledStmt), objectTypeMap)...)
		case *ast.ExprStmt:
			ret = append(ret, ParseExprStmt((*stmt).(*ast.ExprStmt))...)
		case *ast.SendStmt:
		case *ast.IncDecStmt:
			ret = append(ret, ParseIncDecStmt((*stmt).(*ast.IncDecStmt))...)
		case *ast.AssignStmt:
			ret = append(ret, ParseAssignStmt((*stmt).(*ast.AssignStmt), objectTypeMap)...)
		case *ast.GoStmt:
		case *ast.DeferStmt:
		case *ast.ReturnStmt:
			ret = append(ret, ParseReturnStmt((*stmt).(*ast.ReturnStmt))...)
		case *ast.BranchStmt:
			ret = append(ret, ParseBranchStmt((*stmt).(*ast.BranchStmt))...)
		case *ast.BlockStmt:
			ret = append(ret, ParseBlockStmt((*stmt).(*ast.BlockStmt), objectTypeMap)...)
		case *ast.IfStmt:
			ret = append(ret, ParseIfStmt((*stmt).(*ast.IfStmt), objectTypeMap)...)
		case *ast.CaseClause:
			ret = append(ret, ParseCaseClause((*stmt).(*ast.CaseClause), objectTypeMap)...)
		case *ast.SwitchStmt:
			ret = append(ret, ParseSwitchStmt((*stmt).(*ast.SwitchStmt), objectTypeMap)...)
		case *ast.TypeSwitchStmt:
		case *ast.CommClause:
		case *ast.SelectStmt:
		case *ast.ForStmt:
			ret = append(ret, ParseForStmt((*stmt).(*ast.ForStmt), objectTypeMap)...)
		case *ast.RangeStmt:
			ret = append(ret, ParseRangeStmt((*stmt).(*ast.RangeStmt), objectTypeMap)...)
	}
	return ret
}

func ParseBlockStmt(blockStmt *ast.BlockStmt, objectTypeMap *ObjectTypeMap) [] string {
	var ret []string
	ret = append(ret, "{")
	for _, stmt := range blockStmt.List {
		ret = append(ret, ParseStmt(&stmt, objectTypeMap)...)
	}
	ret = append(ret, "}")
	return ret
}