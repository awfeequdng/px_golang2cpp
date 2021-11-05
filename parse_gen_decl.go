package main

import (
	"go/ast"
	"go/token"
	"log"
)

func ParseGenDecl(gen_decl *ast.GenDecl) []string {
	var ret []string
	switch gen_decl.Tok {
	case token.IMPORT:
		ret = append(ret, ParseImport(gen_decl)...)
	case token.CONST:
		ret = append(ret, ParseConst(gen_decl)...)
	case token.VAR:
		ret = append(ret, ParseVar(gen_decl)...)
	case token.TYPE:
		ret = append(ret, ParseType(gen_decl)...)
	default:
		log.Fatal("invalid genDecl token type")
	}
	return ret
}