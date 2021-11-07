package main

import "go/ast"

func ParseSelectorExpr(selectorExpr *ast.SelectorExpr) string {
	var ret string
	x := selectorExpr.X
	xName := ParseExpr(x)
	sel := selectorExpr.Sel
	selName := ParseExpr(sel)
	if _, ok := importMap[xName]; ok {
		// x_name is a package name
		ret = xName + "::" + selName
	} else {
		// x_name is a object name
		ret = xName + "." + selName
	}
	return ret
}