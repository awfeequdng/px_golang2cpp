package main

import (
	"go/ast"
	"log"
	"strings"
)

var globalObjectMap map[string]string

func ParseFuncRev(list *ast.FieldList) (rec, obj string) {
	if len(list.List) > 1 {
		log.Fatal("rev field count greater than 1")
	}
	if len(list.List[0].Names) > 1 {
		log.Fatal("rev field's name count greater than 1")
	}
	obj = ParseExpr(list.List[0].Names[0])
	if starExpr, ok := list.List[0].Type.(*ast.StarExpr); ok {
		rec = ParseExpr(starExpr.X)
	} else {
		log.Fatal("rev's type is not StartExpr")
	}
	return rec, obj
}

func ParseFuncSignature(decl *ast.FuncDecl, objectTypeMap *ObjectTypeMap) (FuncRet, FuncName, FuncParams string) {
	name := decl.Name.Name
	funcType := decl.Type
	params := ParseFieldList(funcType.Params)
	var results []string
	if funcType.Results != nil {
		results = ParseFieldList(funcType.Results)
	}
	FuncParams = "(" + strings.Join(params, ",") + ")"
	FuncName = name
	if len(results) == 0 {
		FuncRet = "void"
	} else if len(results) == 1 {
		FuncRet = results[0]
	} else if len(results) == 2 {
		includeFileMap["std::pair"] = "utility"
		FuncRet = "std::pair<" + results[0] + "," + results[1] + ">"
	} else {
		includeFileMap["std::tuple"] = "tuple"
		tuple := "std::tuple<"
		for id, r := range results {
			if id == 0 {
				tuple += r
			} else {
				tuple += ", " + r
			}
		}
		tuple += ">"
		FuncRet = tuple
	}
	return FuncRet, FuncName, FuncParams
}

func ParseCommonFuncDecl(decl *ast.FuncDecl, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	funcRet, funcName, funcParams := ParseFuncSignature(decl, objectTypeMap)
	ret = append(ret, funcRet + funcName + funcParams)
	body := ParseBlockStmt(decl.Body, objectTypeMap)
	ret = append(ret, body...)

	return ret
}

var structFuncDeclMap map[string][]string = make(map[string][]string)

func GetStructFuncDeclMap() map[string][]string {
	return structFuncDeclMap
}

var structFuncDefinitionMap map[string][]string = make(map[string][]string)

func GetStructFuncDefinitionMap() map[string][]string {
	return structFuncDefinitionMap
}

func ParseMemberFuncDecl(decl *ast.FuncDecl, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	var rev string
	var revObj string
	rev, revObj = ParseFuncRev(decl.Recv)
	funcRet, funcName, funcParams := ParseFuncSignature(decl, objectTypeMap)
	// add struct type to function map
	structFuncDeclMap[rev] = append(structFuncDeclMap[rev], funcRet + " " + funcName + funcParams + ";")
	body := ParseBlockStmt(decl.Body, objectTypeMap)
	strBody := strings.Join(body, "\n")
	strBody = strings.ReplaceAll(strBody, revObj + ".", "this->")

	structFuncDefinitionMap[rev] = append(structFuncDefinitionMap[rev],
		funcRet + " " + rev + "::" + funcName + funcParams + strBody)

	return ret
}

func ParseFuncDecl(decl *ast.FuncDecl, objectTypeMap *ObjectTypeMap) []string {
	if decl.Recv != nil {
		// member function
		return ParseMemberFuncDecl(decl, objectTypeMap)
	}
	// common function
	return ParseCommonFuncDecl(decl, objectTypeMap)
}