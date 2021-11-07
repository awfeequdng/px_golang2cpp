package main

import (
	"go/ast"
	"log"
	"runtime"
	"strconv"
	"strings"
)

// parsed var type
const (
	VarArrayType = iota
	VarMapType
	VarNormalType
	VarNoneType
)

// func ParseArrayType(Type *ast.ArrayType, Elts []ast.Expr) (tname string, values []string) {

// 	values = append(values, "{");
// 	if ident, ok := Type.Elt.(*ast.Ident); ok {
// 		tname = ident.Name
// 	} else {
// 		log.Fatal("INVALID ELEMENT")
// 	}

// 	for _, elt := range Elts {
// 		values = append(values, ParseExpr(elt) + ",")
// 		// if bl, ok := elt.(*ast.BasicLit); ok {
// 		// 	// todo: resolve independent
// 		// 	values = append(values, bl.Value + ", ");
// 		// } else {
// 		// 	log.Fatal("invalid array values")
// 		// }
// 	}
// 	values = append(values, "};")
// 	return tname, values
// }

type type_value struct {
	var_type int
	tname string
	values[]string
}

func genTypeValues(name string, tvalue type_value) string {
	var ret string
	switch tvalue.var_type {
	case VarArrayType:
		ret = tvalue.tname + " " + name + strings.Join(tvalue.values, "\n") + ";"

	case VarMapType:
		ret = tvalue.tname + " " + name + strings.Join(tvalue.values, "\n") + ";"

	case VarNormalType:
		// normal is not append with '{}', so use '=' instead
		ret = tvalue.tname + " " + name + " = " + strings.Join(tvalue.values, "\n") + ";"

	case VarNoneType:
		_, file, line, _ := runtime.Caller(0)
		log.Fatal("INVALID VAR TYPE" + file + ", lien: " + strconv.Itoa(line))
	}

	return ret
}

func ParseGenDeclVar(decl *ast.GenDecl, objectTypeMap *ObjectTypeMap) []string {
	var ret []string
	for _, spec := range decl.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			var names []string
			var tvalues []type_value
			for _, name := range vs.Names {
				names = append(names, name.Name)
			}
			for _, value := range vs.Values {
				if cl, ok := value.(*ast.CompositeLit); ok {
					typ, tname, vals := ParseCompositeLit(cl)
					tv := type_value{typ, tname, vals}
					tvalues = append(tvalues, tv)
				} else if _, ok := value.(*ast.CallExpr); ok {
					typ := VarNormalType
					var tname string
					if vs.Type == nil {
						tname = "auto"
					} else {
						tname = ParseExpr(vs.Type)
					}
					var vals []string
					vals = append(vals, ParseExpr(value))
					tv := type_value{typ, tname, vals}
					tvalues = append(tvalues, tv)
				} else {
					typ := VarNormalType
					tname := "auto "
					var vals []string
					vals = append(vals, ParseExpr(value))
					tv := type_value{typ, tname, vals}
					tvalues = append(tvalues, tv)
				}
			}

			if len(names) == len(tvalues) {
				for i, name := range names {
					ret = append(ret, genTypeValues(name, tvalues[i]))
				}
			} else if len(names) == 1 && len(tvalues) > 0 {
				log.Fatal("NOT IMPLEMENTED YET")
			} else if len(tvalues) == 0 {
				// have name and type
				if vs.Type == nil {
					log.Fatal("can not be nil when values is nil")
				}
				tname := ParseExpr(vs.Type)
				typ := VarMapType
				vals := []string{"{}"}
				tv := type_value{typ, tname, vals}
				tvalues = append(tvalues, tv)
				for _, name := range names {
					ret = append(ret, genTypeValues(name, tv))
				}

				// if ident, ok := vs.Type.(*ast.Ident); ok {
				// 	typ := VAR_NORMAL_TYPE
				// 	tname := ident.Name
				// 	vals := []string{"{}"}
				// 	tv := type_value{typ, tname, vals}
				// 	// tvalues = append(tvalues, tv)
				// 	for _, name := range names {
				// 		ret = append(ret, genTypeValues(name, tv))
				// 	}
				// } else {
				// 	log.Fatal("type must be ast.Ident")
				// }
 			} else {
				 log.Print("names: " + strings.Join(names, ","))
				 log.Print("values: ")
				 for _, v := range tvalues {
					 log.Print(strings.Join(v.values, " "))
				 }
				log.Fatal("invalid names and values size , name size: " + strconv.Itoa(len(names)) + ", value size: " + strconv.Itoa(len(tvalues)))
			}
		} else {
			log.Fatal("NOT IMPLEMENTED YET")
		}
	}

	return ret
}