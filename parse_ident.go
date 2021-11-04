package main

import "go/ast"

func ParseIdent(ident *ast.Ident) string {
	return ident.Name
}