package main

import (
	"errors"
	"strings"
)

type ObjectTypeMap struct {
	// object -> object type name
	typeMap map[string]string
	next *ObjectTypeMap
}

func FindObjectMap(objectTypeMap *ObjectTypeMap, objName string) (typeName string,err error) {
	if v, ok := objectTypeMap.typeMap[objName]; ok {
		return v, nil
	}
	if objectTypeMap.next != nil {
		return FindObjectMap(objectTypeMap.next, objName)
	}
	return "", errors.New("can not find object name : " + objName)
}

func getValueType(value string) string {
	if strings.Contains(value, "map") {
		return "map"
	} else if strings.Contains(value, "vector") {
		return "array"
	} else if strings.Contains(value, "\"") {
		return "string"
	} else {
		return "int"
	}
	//else {
	//	log.Fatal("invalid value type")
	//}
	//return ""
}

func (otm* ObjectTypeMap)InsertObjectMap(name string, typeName string, value string) {
	if strings.Contains(typeName, "auto") {
		typeName = getValueType(value)
	}
	otm.typeMap[name] = typeName
}