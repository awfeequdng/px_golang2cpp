package main

import (
	"go/ast"
	"go/token"
	"log"
)

func ParseAssignStmt(assignStmt *ast.AssignStmt, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	var names []string
	var values []string
	switch assignStmt.Tok {
		case token.DEFINE:
		case token.ASSIGN:
		case token.ADD_ASSIGN:
		case token.SUB_ASSIGN:
		case token.MUL_ASSIGN:
		case token.QUO_ASSIGN:
		case token.REM_ASSIGN:
		case token.AND_ASSIGN:
		case token.OR_ASSIGN:
		default:
			log.Fatal("do not support, token = " + assignStmt.Tok.String())
	}

	for _, lhs := range assignStmt.Lhs {
		names = append(names, ParseExpr(lhs))
	}
	for _, rhs := range assignStmt.Rhs {
		values = append(values, ParseExpr(rhs))
	}
	nameSize := len(names)
	valueSize := len(values)
	// if name_size != value_size {
	// 	log.Fatalf("name size: %d not equal value size: %d", name_size, value_size)
	// }

	var name string
	var value string

	if nameSize == 0 {
		log.Fatal("name size is 0")
	}
	if valueSize == 0 {
		log.Fatal("value size if 0")
	}

	if nameSize == 1 {
		name = names[0]
		value = values[0]
		switch assignStmt.Tok {
		case token.DEFINE:
			ret = append(ret, "auto " + name + " = " + value + ";")
		case token.ASSIGN:
			fallthrough
		case token.ADD_ASSIGN:
			fallthrough
		case token.SUB_ASSIGN:
			fallthrough
		case token.MUL_ASSIGN:
			fallthrough
		case token.QUO_ASSIGN:
			fallthrough
		case token.REM_ASSIGN:
			fallthrough
		case token.AND_ASSIGN:
			fallthrough
		case token.OR_ASSIGN:
			ret = append(ret, name + assignStmt.Tok.String() + value + ";")
		default:
			log.Fatal("not support assign operation, token = " + assignStmt.Tok.String())
		}
	} else if assignStmt.Tok == token.DEFINE || assignStmt.Tok == token.ASSIGN {
		for id, n := range names {
			if id == 0 {
				name += n
			} else {
				name += ", " + n
			}
		}
		for id, v := range values {
			if id == 0 {
				value += v
			} else {
				value += ", " + v
			}
		}
		switch assignStmt.Tok {
		case token.DEFINE:
			ret = append(ret, "auto [" + name + "] = " + value + ";")
		case token.ASSIGN:
			includeFileMap["std::tie"] = "tuple"
			ret = append(ret, "std::tie(" + name + ") = {" + value + "};")
		default:
			log.Fatal("NOT SUPPORT MULTI VALUE ASSIGN, token = " + assignStmt.Tok.String())
		}
	} else {
		log.Fatal("NOT SUPPORT MULTI VALUE ASSIGN, token = " + assignStmt.Tok.String())
	}

	return ret
}