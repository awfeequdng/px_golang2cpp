package main

import "go/ast"

var typeMap = map[string]string {
	"float64": "double",
	"float32": "float",
	"int": "int64_t",
}

func ParseFieldList(fieldList *ast.FieldList) []string {
	var ret []string
	// var id = 0;
	for _, field := range fieldList.List {
		typ := ParseExpr(field.Type)
		// convert go type to c++ type
		if val, ok := typeMap[typ]; ok {
			typ = val
		}
		if field.Names != nil {
			for _, name := range field.Names {
				ret = append(ret, typ + " " + name.Name)
				// if id == 0 {
				// 	ret = append(ret, typ + " " + name.Name)
				// } else {
				// 	ret = append(ret, ", " + typ + " " + name.Name)
				// }
				// id++
			}
		} else {
			ret = append(ret, typ)
		}

	}
	return ret
}
