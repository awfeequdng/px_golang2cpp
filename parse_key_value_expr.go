package main

import (
	"go/ast"
)

// return: type and values
func ParseKeyValueExpr(key_value_expr *ast.KeyValueExpr) string {
	key := key_value_expr.Key
	val := key_value_expr.Value
	var ret string
	ret += "{"
	ret += ParseExpr(key) + " , "
	ret += ParseExpr(val) + "},"
	return ret
}