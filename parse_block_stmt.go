package main

import (
	"go/ast"
	"log"
)

func ParseDeclStmt(decl_stmt *ast.DeclStmt) []string {
	var ret []string
	if gen_decl, ok := decl_stmt.Decl.(*ast.GenDecl); ok {
		ret = append(ret, ParseGenDecl(gen_decl)...)
	} else {
		log.Fatal("not implemented decl otherwise gendecl")
	}
	return ret
}

func ParseStmt(stmt *ast.Stmt) []string {
	var ret []string
	switch (*stmt).(type) {
		case *ast.BadStmt:
			log.Fatal("bad statement")
		case *ast.DeclStmt:
			ret = append(ret, ParseDeclStmt((*stmt).(*ast.DeclStmt))...)
		case *ast.EmptyStmt:
			// do nothing
		case *ast.LabeledStmt:
			ret = append(ret, ParseLabeledStmt((*stmt).(*ast.LabeledStmt))...)
		case *ast.ExprStmt:
			ret = append(ret, ParseExprStmt((*stmt).(*ast.ExprStmt))...)
		case *ast.SendStmt:
		case *ast.IncDecStmt:
		case *ast.AssignStmt:
			ret = append(ret, ParseAssignStmt((*stmt).(*ast.AssignStmt))...)
		case *ast.GoStmt:
		case *ast.DeferStmt:
		case *ast.ReturnStmt:
			ret = append(ret, ParseReturnStmt((*stmt).(*ast.ReturnStmt))...)
		case *ast.BranchStmt:
			ret = append(ret, ParseBranchStmt((*stmt).(*ast.BranchStmt))...)
		case *ast.BlockStmt:
			ret = append(ret, ParseBlockStmt((*stmt).(*ast.BlockStmt))...)
		case *ast.IfStmt:
			ret = append(ret, ParseIfStmt((*stmt).(*ast.IfStmt))...)
		case *ast.CaseClause:
			ret = append(ret, ParseCaseClause((*stmt).(*ast.CaseClause))...)
		case *ast.SwitchStmt:
			ret = append(ret, ParseSwitchStmt((*stmt).(*ast.SwitchStmt))...)
		case *ast.TypeSwitchStmt:
		case *ast.CommClause:
		case *ast.SelectStmt:
		case *ast.ForStmt:
		case *ast.RangeStmt:
	}
	return ret
}

func ParseBlockStmt(blockStmt *ast.BlockStmt) [] string {
	var ret []string
	ret = append(ret, "{")
	for _, stmt := range blockStmt.List {
		ret = append(ret, ParseStmt(&stmt)...)
	}
	ret = append(ret, "}")
	return ret
}