package main

import "errors"

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