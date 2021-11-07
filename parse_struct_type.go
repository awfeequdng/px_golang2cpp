package main

import (
	"go/ast"
	"strings"
)

func ParseStructType(structType *ast.StructType) string {
	var ret []string

	ret = append(ret, "struct {")
	if structType.Fields != nil {
		fields := ParseFieldList(structType.Fields)
		for _, field := range fields {
			ret = append(ret, field + ";")
		}
	}
	ret = append(ret, "}")
	return strings.Join(ret, "\n")
}
