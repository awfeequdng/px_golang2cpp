package main

import (
	"go/ast"
)

func ParseKeyValueExpr(keyValueExpr *ast.KeyValueExpr) string {
	key := keyValueExpr.Key
	val := keyValueExpr.Value
	var ret string
	ret += "{"
	ret += ParseExpr(key) + " , "
	ret += ParseExpr(val) + "}"
	return ret
}