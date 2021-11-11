package main

import (
	"errors"
	"fmt"
	"strings"
)

type ObjectTypeMap struct {
	// object -> object type name
	typeMap map[string]string
	next *ObjectTypeMap
}

func StrBetween(str, left, right string) string {
	lPos := strings.Index(str, left)
	rPos := strings.Index(str, right)
	return str[lPos + len(left): rPos]
}

func InsertInitObjectMap(objectTypeMap *ObjectTypeMap) {
	objectTypeMap.InsertObjectMap("n.Fields", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Items", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.TableHints", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Columns", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Lists", "std::vector", "")
	objectTypeMap.InsertObjectMap("list", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Setlist", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Keys", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Options", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Cols", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Constraints", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Tables", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.IndexPartSpecifications", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.NewColumns", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.NewConstraints", "std::vector", "")
	objectTypeMap.InsertObjectMap("n.Specs", "std::vector", "")

	// model.go
	objectTypeMap.InsertObjectMap("t.Columns", "std::vector", "")
	objectTypeMap.InsertObjectMap("t.Indices", "std::vector", "")
	objectTypeMap.InsertObjectMap("t.ForeignKeys", "std::vector", "")
	objectTypeMap.InsertObjectMap("index.Columns", "std::vector", "")
	objectTypeMap.InsertObjectMap("db.Tables", "std::vector", "")
	objectTypeMap.InsertObjectMap("idx.Columns", "std::vector", "")
}

func FindObjectMap(objectTypeMap *ObjectTypeMap, objName string) (typeName string,err error) {
	// objName may be a function name or a temporary value
	if v, ok := objectTypeMap.typeMap[objName]; ok {
		return v, nil
	}
	if strings.HasPrefix(objName, "std::vector") {
		return "std::vector", nil
	} else if strings.HasPrefix(objName, "std::map") {
		return "std::map", nil
	}

	if objectTypeMap.next != nil {
		return FindObjectMap(objectTypeMap.next, objName)
	}
	return "", errors.New("can not find object name : " + objName)
}

func getValueType(objectTypeMap *ObjectTypeMap, value string) string {
	if strings.Contains(value, "std::map") {
		return "std::map"
	} else if strings.Contains(value, "std::vector") {
		return "std::vector"
	} else if strings.HasPrefix(value, "\"") || strings.HasPrefix(value, "std::string") {
		return "std::string"
	} else {
		// parse slice object
		if strings.Contains(value, "slice(") && strings.Contains(value, ",") {
			// value is a slice
			var sliceObj = StrBetween(value, "slice(", ",")
			if typ, err := FindObjectMap(objectTypeMap, strings.TrimSpace(sliceObj)); err != nil {
				return "known"
			} else {
				return typ
			}
		}
		return "unknown"
	}
	//else {
	//	log.Fatal("invalid value type")
	//}
	//return ""
}

func (otm* ObjectTypeMap)InsertObjectMap(name string, typeName string, value string) {
	if strings.HasPrefix(typeName, "auto") {
		typeName = getValueType(otm, value)
	}
	otm.typeMap[name] = typeName
}

func (otm* ObjectTypeMap)PrintObjectMap() {
	for k, v := range otm.typeMap {
		fmt.Println("key = " + k + ", value = " + v)
	}
}