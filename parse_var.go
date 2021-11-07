package main

import (
	"go/ast"
	"log"
	"strconv"
	"strings"
)

func ParseGenDeclVar(decl *ast.GenDecl, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	for _, spec := range decl.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			var names []string
			var values []string
			var tname string

			if vs.Type == nil {
				tname = "auto"
			} else {
				tname = ParseExpr(vs.Type)
			}

			for _, name := range vs.Names {
				names = append(names, name.Name)
			}
			for _, value := range vs.Values {
				values = append(values, ParseExpr(value))
			}

			if len(names) == len(values) {
				for idx, name := range names {
					ret = append(ret, tname + " " + name + " = " + values[idx] + ";")
				}
			} else if len(names) > 1 && len(values) == 1 {
				for _, name := range names {
					ret = append(ret, tname + " " + name + " = " + values[0])
				}
			} else if len(values) == 0 && len(names) > 0 {
				// have name and type
				if vs.Type == nil {
					log.Fatal("can not be nil when values is nil")
				}
				for _, name := range names {
					ret = append(ret, tname + " " + name + ";")
				}
 			} else {
				 log.Print("names: " + strings.Join(names, ","))
				 log.Print("values: ")
				 log.Print(strings.Join(values, ","))
				 log.Fatal("invalid names and values size , name size: " + strconv.Itoa(len(names)) + ", value size: " + strconv.Itoa(len(values)))
			}
		} else {
			log.Fatal("NOT IMPLEMENTED YET")
		}
	}

	return ret
}