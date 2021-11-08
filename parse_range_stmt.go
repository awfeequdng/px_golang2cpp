package main

import (
	"go/ast"
	"log"
	"strings"
)

func ParseRangeStmt(stmt *ast.RangeStmt, objectTypeMap *ObjectTypeMap) []string {
	var ret[]string

	var key string
	var val string
	tok := stmt.Tok.String()

	x := ParseExpr(stmt.X)

	objType, err := FindObjectMap(objectTypeMap, x)
	if err != nil {
		log.Fatal("can not find object type: " + x)
	}
	if strings.HasPrefix(objType, "unknown") {
		log.Fatal("unknown object type: " + x)
	}
	// whether range object is a map
	var isMap = strings.Contains(objType, "map")

	if stmt.Key != nil {
		key = ParseExpr(stmt.Key)
	}
	if stmt.Value != nil {
		val = ParseExpr(stmt.Value)
	}
	var keyValType string
	if tok == ":=" {
		keyValType = "auto "
	} else if tok == "=" {
		keyValType = ""
	} else {
		log.Fatal("invalid token type in range for: " + tok)
	}

	if isMap {
		// range for map
		ret = append(ret, "for (" + keyValType + " [" + key + ", " + val + "] : " + x + ") {")
	} else {
		if key == "_" {
			ret = append(ret, "for (" + keyValType + " " + val + " : " + x + ") {")
		} else {
			ret = append(ret, "for (" + keyValType + " " + key + " = 0; " + key + "< " + x + ".size(); " + key + "++ ) {")
			ret = append(ret, keyValType + val + " = " + x + "[" + key + "];")
		}
	}
	if stmt.Body == nil {
		ret = append(ret, "{}")
	} else {
		ret = append(ret, ParseBlockStmt(stmt.Body, objectTypeMap)...)
	}
	ret = append(ret, "}")

	return ret
}
