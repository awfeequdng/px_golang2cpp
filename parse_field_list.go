package main

import "go/ast"

func ParseFieldList(field_list *ast.FieldList) []string {
	var ret []string
	var id = 0;
	for _, field := range field_list.List {
		typ := ParseExpr(field.Type)
		// var names []string
		for _, name := range field.Names {
			// names = append(names, name.Name)
			if id == 0 {
				ret = append(ret, typ + " " + name.Name)
			} else {
				ret = append(ret, ", " + typ + " " + name.Name)
			}
			id++
		}
	}
	return ret
}
