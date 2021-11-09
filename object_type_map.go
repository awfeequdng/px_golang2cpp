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