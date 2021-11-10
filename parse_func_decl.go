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
		// log.Fatal("rev's type is not StartExpr")
		//log.Print("rev's type is not StartExpr")
		rec = ParseExpr(list.List[0].Type)
	}
	return rec, obj
}

func ParseObjectTypeMapFromParams(objectTypeMap *ObjectTypeMap, params []string) {
	for _, param := range params {
		res := strings.Split(strings.TrimSpace(param), " ")
		if len(res) != 2 {
			log.Fatal("not key-value pair in param: " + param)
		}
		objectTypeMap.typeMap[res[1]] = res[0]
	}
}

func ParseFuncSignature(decl *ast.FuncDecl, objectTypeMap *ObjectTypeMap) (funcRet, funcName, funcParams, funcVars, retValues string) {
	name := decl.Name.Name
	funcType := decl.Type

	params := ParseFieldList(funcType.Params)
	ParseObjectTypeMapFromParams(objectTypeMap, params)

	var results []string
	if funcType.Results != nil {
		results = ParseFieldList(funcType.Results)
	}

	funcParams = "(" + strings.Join(params, ",") + ")"
	funcName = name
	if len(results) == 0 {
		funcRet = "void"
	} else if len(results) == 1 {
		res := strings.TrimSpace(results[0])
		if strings.Contains(res, " ") {
			reses := strings.Split(res, " ")
			if len(reses) != 2 {
				log.Fatal("not key-value pair: " + strings.Join(reses, " "))
			}
			funcRet = reses[0]
			retValues = reses[1]
			funcVars = results[0] + ";"
		} else {
			funcRet = results[0]
		}
	} else if len(results) == 2 {
		includeFileMap["std::pair"] = "utility"
		//funcRet = "std::pair<" + results[0] + "," + results[1] + ">"
		res1 := strings.TrimSpace(results[0])
		res2 := strings.TrimSpace(results[1])
		if strings.Contains(res1, " ") && strings.Contains(res2, " ") {
			res1s := strings.Split(res1, " ")
			res2s := strings.Split(res2, " ")
			if len(res1s) != 2 || len(res2s) != 2 {
				log.Fatal("not key value pair: " + strings.Join(res1s, " ") + ", " + strings.Join(res2s, " "))
			}
			funcRet = "std::pair<" + res1s[0] + "," + res2s[0] + ">"
			retValues = "{" + res1s[01] + ", " + res2s[1] + "}"
			funcVars = results[0] + ";"
			funcVars += results[1] + ";"
		} else if !strings.Contains(res1, " ") && !strings.Contains(res2, " ") {
			funcRet = "std::pair<" + results[0] + "," + results[1] + ">"
		} else {
			log.Fatal("not all key-value pair: " + results[0] + ", " + results[1])
		}
	} else {
		includeFileMap["std::tuple"] = "tuple"
		if strings.Contains(strings.TrimSpace(results[0]), " ") {
			// key-value pair
			tuple := "std::tuple<"
			retValues = "{ "
			for id, r := range results {
				res := strings.TrimSpace(r)
				if !strings.Contains(res, " ") {
					log.Fatal("not key-value pair: " + r)
				}
				reses := strings.Split(res, " ")
				if len(reses) != 2 {
					log.Fatal("not key-value pair: " + r)
				}

				if id == 0 {
					tuple += reses[0]
					retValues += reses[1]
				} else {
					tuple += ", " + reses[0]
					retValues += ", " + reses[1]
				}
				funcVars += r + ";"
			}
			tuple += ">"
			retValues += "}"
			funcRet = tuple
		} else {
			// only type
			tuple := "std::tuple<"
			for id, r := range results {
				if strings.Contains(strings.TrimSpace(r), " ") {
					log.Fatal("is key-value pair: " + r)
				}
				if id == 0 {
					tuple += r
				} else {
					tuple += ", " + r
				}
			}
			tuple += ">"
			funcRet = tuple
		}
	}
	return funcRet, funcName, funcParams, funcVars, retValues
}

func ParseCommonFuncDecl(decl *ast.FuncDecl, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	funcRet, funcName, funcParams, funcVars, retValues := ParseFuncSignature(decl, objectTypeMap)
	ret = append(ret, funcRet + " " + funcName + funcParams)
	body := ParseBlockStmt(decl.Body, objectTypeMap)
	strBody := strings.Join(body, "\n")
	strBody = "{" + funcVars + strBody + "}"

	if strings.Contains(strBody, "return;") {
		strBody = strings.ReplaceAll(strBody, "return;", "return " + retValues + ";")
	}
	ret = append(ret, strBody)
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
	funcRet, funcName, funcParams, funcVars, retValues := ParseFuncSignature(decl, objectTypeMap)
	// add struct type to function map
	structFuncDeclMap[rev] = append(structFuncDeclMap[rev], funcRet + " " + funcName + funcParams + ";")
	body := ParseBlockStmt(decl.Body, objectTypeMap)

	strBody := strings.Join(body, "\n")
	strBody = "{ " + funcVars + " " + strBody  + " }"
	strBody = strings.ReplaceAll(strBody, revObj + ".", "this->")
	if strings.Contains(strBody, "return;") {
		strBody = strings.ReplaceAll(strBody, "return;", "return " + retValues + ";")
	}

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