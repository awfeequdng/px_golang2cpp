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

func ParseBlockStmt(blockStmt *ast.BlockStmt) [] string {
	var ret []string
	ret = append(ret, "{")
	for _, stmt := range blockStmt.List {
		switch stmt.(type) {
			case *ast.BadStmt:
			case *ast.DeclStmt:
				ret = append(ret, ParseDeclStmt(stmt.(*ast.DeclStmt))...)
			case *ast.EmptyStmt:
			case *ast.LabeledStmt:
			case *ast.ExprStmt:
			case *ast.SendStmt:
			case *ast.IncDecStmt:
			case *ast.AssignStmt:
				ret = append(ret, ParseAssignStmt(stmt.(*ast.AssignStmt))...)
			case *ast.GoStmt:
			case *ast.DeferStmt:
			case *ast.ReturnStmt:
			case *ast.BranchStmt:
			case *ast.BlockStmt:
			case *ast.IfStmt:
			case *ast.CaseClause:
			case *ast.SwitchStmt:
			case *ast.TypeSwitchStmt:
			case *ast.CommClause:
			case *ast.SelectStmt:
			case *ast.ForStmt:
			case *ast.RangeStmt:
		}
	}
	ret = append(ret, "}")
	return ret
}