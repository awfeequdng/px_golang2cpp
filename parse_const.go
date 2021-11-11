package main

import (
	"go/ast"
	"go/token"
	"log"
	"strconv"
	"strings"
)

func ParseGenDeclConst(decl *ast.GenDecl, objectTypeMap *ObjectTypeMap) []string {
	var ret []string

	iotaStart := 0
	incFlag := false
	shiftFlag := false

	for sIdx, spec := range decl.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			var names []string
			for _, name := range vs.Names {
				names = append(names, name.Name)
			}
			var values []string
			for vIdx, value := range vs.Values {
				if v, ok := value.(*ast.BasicLit); ok {
					switch v.Kind {
					case token.STRING:
						values = append(values, v.Value)
						break
					case token.INT:
						values = append(values, v.Value)
						break
					case token.FLOAT:
						values = append(values, v.Value)
						break
					case token.CHAR:
						values = append(values, v.Value)
						break
					default:
						log.Fatal("invalid basicLit type")
					}
				} else {
					val := ParseExpr(value)
					if strings.Contains(val, "iota") {
						if vIdx != 0 {
							log.Fatal("only one value when encounter iota")
						}
						if val == "iota" {
							iotaStart = 0
							incFlag = true
							shiftFlag = false
							values = append(values, "0")
						} else if strings.Contains(val, "1<<iota") {
							iotaStart = 0
							incFlag = false
							shiftFlag = true
							values = append(values, "1ul")
						} else if strings.Contains(val, "iota+1") {
							iotaStart = 1
							incFlag = false
							shiftFlag = true
							values = append(values, "1ul")
						} else {
							log.Fatal("unknown value : " + val)
						}
					} else {
						values = append(values, val)
					}
				}
			}
			if sIdx > 0 && len(vs.Values) == 0 {
				if incFlag {
					iotaStart++
					values = append(values, strconv.Itoa(iotaStart))
				} else if shiftFlag {
					iotaStart++
					values = append(values, "(1ul << " + strconv.Itoa(iotaStart) + ")")
				}
			}

			if len(names) != len(values) {
				if len(values) == 1 {
					for _, name := range names {
						ret = append(ret, "#define " + name + " " + values[0])
					}
				} else {
					log.Print("invalid names size and values size")
					log.Print("names: " + strings.Join(names, " "))
					log.Print("values: " + strings.Join(values, " "))
					log.Fatal("invalid sizes")
				}
			} else {
				for i, name := range names {
					ret = append(ret, "#define " + name + " " + values[i])
				}
			}
		}
	}
	return ret
}