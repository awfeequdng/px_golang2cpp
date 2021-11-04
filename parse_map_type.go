package main

import "go/ast"

func ParseMapType(Type *ast.MapType) string {
	key := ParseExpr(Type.Key)
	val := ParseExpr(Type.Value)

	ret := "std::unordered_map<" + key + "," + val + ">"

	return ret
}
