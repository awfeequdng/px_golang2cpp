package main

import (
	"go/ast"
	"log"
)

func ParseCallExpr(callExpr *ast.CallExpr) string {
	var ret string

	fn := callExpr.Fun
	fnName := ParseExpr(fn)
	args := callExpr.Args
	ret += fnName + "("
	for id, arg := range args {
		argName := ParseExpr(arg)
		if id == 0 {
			ret += argName
		} else {
			ret += "," + argName
		}
	}
	ret += ")"
	if fnName == "make" {
		// cpp do not use make to initial object, so return {}
		log.Print("call fun is: " + ret + ", but we return {}")
		return "{}"
	}
	return ret
}