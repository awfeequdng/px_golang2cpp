package main

import (
	"go/ast"
	"go/token"
	"log"
)

func ParseBranchStmt(stmt *ast.BranchStmt) []string {
	var ret[]string
	//BREAK, CONTINUE, GOTO, FALLTHROUGH
	switch stmt.Tok {
	case token.BREAK:
		log.Fatal("not implemented BREAK")
	case token.GOTO:
		log.Fatal("not implemented GOTO")
	case token.CONTINUE:
		log.Fatal("not implemented CONTINUE")
	case token.FALLTHROUGH:
		ret = append(ret, "[[fallthrough]];")
	default:
		log.Fatal("invalid token type: " + stmt.Tok.String())
	}
	return ret
}
