package main

import (
	"go/ast"
	"log"
	"strings"
)

var structDeclMap map[string]string = make(map[string]string)

func GetStructDeclMap() map[string]string {
	return structDeclMap
}

func GetStructDeclAndDefinition() []string {
	var ret []string

	// parse struct's member function
	stFuncDeclMap := GetStructFuncDeclMap()
	for name, decl := range GetStructDeclMap() {
		if sig, ok := stFuncDeclMap[name]; ok {
			lastBrace := strings.LastIndex(decl, "}")
			decl = decl[:lastBrace]
			ret = append(ret, "typedef ")
			ret = append(ret, decl)
			ret = append(ret, sig...)
			ret = append(ret, "}")
			ret = append(ret, name + ";")
		}
	}

	definitionMap := GetStructFuncDefinitionMap()
	for _, stDefinitions := range definitionMap {
		ret = append(ret, stDefinitions...)
	}

	return ret
}

func ParseGenDeclType(decl *ast.GenDecl) []string {
	var ret []string
	for _, spec := range decl.Specs {
		var name string
		if ts, ok := spec.(*ast.TypeSpec); ok {
			name = ts.Name.Name
			typ := ParseExpr(ts.Type)
			//ret = append(ret, "typedef " + typ + " " + name + ";")
			structDeclMap[name] = typ
		} else {
			log.Fatal("invalid spec type in ParseType function ")
		}
	}

	return ret
}