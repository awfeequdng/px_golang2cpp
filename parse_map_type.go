package main

import "go/ast"

func ParseMapType(Type *ast.MapType) string {
	key := ParseExpr(Type.Key)
	val := ParseExpr(Type.Value)
	if typ, ok := typeMap[key]; ok {
		key = typ
	}

	if typ, ok := typeMap[val]; ok {
		val = typ
	}
	ret := "std::unordered_map<" + key + "," + val + ">"

	return ret
}
