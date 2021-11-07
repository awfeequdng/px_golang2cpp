package main

import (
	"go/ast"
	"go/token"
	"log"
	"strconv"
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
		default:
			ret = append(ret, name + assignStmt.Tok.String() + value + ";")
		}
	} else if valueSize == 1 {
		// names size > 1 and value size = 1
		for id, n := range names {
			if id == 0 {
				name += n
			} else {
				name += ", " + n
			}
		}
		value = values[0]
		switch assignStmt.Tok {
		case token.DEFINE:
			ret = append(ret, "auto [" + name + "] = " + value + ";")
		case token.ASSIGN:
			includeFileMap["std::tie"] = "tuple"
			ret = append(ret, "std::tie(" + name + ") = {" + value + "};")
		default:
			log.Fatal("NOT SUPPORT MULTI VALUE ASSIGN, token = " + assignStmt.Tok.String())
		}
	} else if nameSize == valueSize {
		var nameType string
		switch assignStmt.Tok {
		case token.DEFINE:
			nameType = "auto "
		case token.ASSIGN:
			nameType = ""
		default:
			log.Fatal("NOT SUPPORT MULTI VALUE ASSIGN, token = " + assignStmt.Tok.String())
		}
		for idx, n := range names {
			name = n
			value = values[idx]
			ret = append(ret, nameType + name + " = " + value + ";")
		}
	} else {
		log.Fatal("nameSize does not equal valueSize, nameSize: " + strconv.Itoa(nameSize) + ", valueSize: " + strconv.Itoa(valueSize))
	}

	return ret
}