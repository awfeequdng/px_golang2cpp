package main

import "go/ast"

func ParseSelectorExpr(selector_expr *ast.SelectorExpr) string {
	var ret string
	x := selector_expr.X
	x_name := ParseExpr(x)
	sel := selector_expr.Sel
	sel_name := ParseExpr(sel)
	if _, ok := importMap[x_name]; ok {
		// x_name is a package name
		ret = x_name + "::" + sel_name
	} else {
		// x_name is a object name
		ret = x_name + "." + sel_name
	}
	return ret
}