package main

import (
	"go/ast"
	"go/token"
	"log"
)

func ParseAssignStmt(assign_stmt *ast.AssignStmt) []string {
	var ret []string
	var names []string
	var values []string
	switch assign_stmt.Tok {
		case token.DEFINE:
		case token.ASSIGN:
		default:
			log.Fatal("only support ':=' and '=' assignment, not support: " + assign_stmt.Tok.String())
	}

	for _, lhs := range assign_stmt.Lhs {
		names = append(names, ParseExpr(lhs))
	}
	for _, rhs := range assign_stmt.Rhs {
		values = append(values, ParseExpr(rhs))
	}
	name_size := len(names)
	value_size := len(values)
	// if name_size != value_size {
	// 	log.Fatalf("name size: %d not equal value size: %d", name_size, value_size)
	// }

	var name string
	var value string

	if name_size == 0 {
		log.Fatal("name size is 0")
	}
	if value_size == 0 {
		log.Fatal("value size if 0")
	}

	if name_size == 1 {
		name = names[0]
		value = values[0]
		switch assign_stmt.Tok {
		case token.DEFINE:
			ret = append(ret, "auto " + name + " = " + value + ";")
		case token.ASSIGN:
			ret = append(ret, name + " = " + value + ";")
		}
	} else {
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
				name += ", " + v
			}
		}
		switch assign_stmt.Tok {
		case token.DEFINE:
			ret = append(ret, "auto [" + name + "] = " + value + ";")
		case token.ASSIGN:
			includeFileMap["std::tie"] = "tuple"
			ret = append(ret, "std::tie(" + name + ") = " + value + ";")
		}
	}

	return ret
}