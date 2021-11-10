package main

import (
	"go/ast"
	"log"
)

var typeMap = map[string]string {
	"float64": "double",
	"float32": "float",
	"int": "int64_t",
}

func ParseFieldList(fieldList *ast.FieldList) []string {
	var ret []string
	var isTypeNamePair = false
	for _, field := range fieldList.List {
		typ := ParseExpr(field.Type)
		// convert go type to c++ type
		if val, ok := typeMap[typ]; ok {
			typ = val
		}
		if field.Names != nil {
			for _, name := range field.Names {
				ret = append(ret, typ + " " + name.Name)
				isTypeNamePair = true
			}
		} else {
			if isTypeNamePair {
				log.Print("exist type-name pair, so can not exist type only, type: " + typ)
			}
			ret = append(ret, typ)
		}

	}
	return ret
}
