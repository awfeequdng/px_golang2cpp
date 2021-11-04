package main

import (
	"go/ast"
	"log"
)

func ParseCallExpr(call_expr *ast.CallExpr) string {
	var ret string

	fn := call_expr.Fun
	fn_name := ParseExpr(fn)
	args := call_expr.Args
	ret += fn_name + "("
	for id, arg := range args {
		arg_name := ParseExpr(arg)
		if id == 0 {
			ret += arg_name
		} else {
			ret += "," + arg_name
		}
	}
	ret += ")"
	if fn_name == "make" {
		// cpp do not use make to initial object, so return {}
		log.Print("call fun is: " + ret + ", but we return {}")
		return "{}"
	}
	return ret
}