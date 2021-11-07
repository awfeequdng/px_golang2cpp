package main

import (
	"go/ast"
	"go/token"
	"log"
)

func ParseGenDecl(genDecl *ast.GenDecl, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	switch genDecl.Tok {
	case token.IMPORT:
		ret = append(ret, ParseGenDeclImport(genDecl)...)
	case token.CONST:
		ret = append(ret, ParseGenDeclConst(genDecl, objectTypeMap)...)
	case token.VAR:
		ret = append(ret, ParseGenDeclVar(genDecl, objectTypeMap)...)
	case token.TYPE:
		ret = append(ret, ParseGenDeclType(genDecl)...)
	default:
		log.Fatal("invalid genDecl token type")
	}
	return ret
}