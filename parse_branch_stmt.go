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
		//log.Fatal("not implemented BREAK")
		if stmt.Label != nil {
			label := stmt.Label.Name
			ret = append(ret, "goto " + label + ";")
		} else {
			ret = append(ret, "break;")
		}
	case token.GOTO:
		if stmt.Label != nil {
			label := stmt.Label.Name
			ret = append(ret, "goto " + label + ";")
		} else {
			log.Fatal("implement GOTO error: no label")
		}
	case token.CONTINUE:
		if stmt.Label != nil {
			log.Fatal("implement CONTINUE error: do not use label in continue")
		} else {
			ret = append(ret, "continue;")
		}
	case token.FALLTHROUGH:
		ret = append(ret, "[[fallthrough]];")
	default:
		log.Fatal("invalid token type: " + stmt.Tok.String())
	}
	return ret
}
