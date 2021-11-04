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
	VAR_ARRAY_TYPE = iota
	VAR_MAP_TYPE
	VAR_NORMAL_TYPE
	VAR_NONE_TYPE
)

func ParseMapType(Type *ast.MapType, Elts []ast.Expr) (t string, ret []string) {
	var key, val string
	if k, ok := Type.Key.(*ast.Ident); ok {
		key = k.Name
	} else {
		log.Fatal("key can not convert to Ident")
	}

	if v, ok := Type.Key.(*ast.Ident); ok {
		val = v.Name
	} else {
		log.Fatal("value can not convert to Ident")
	}

	ret = append(ret, "{");
	t = "unordered_map<" + key + "," + val + "> "

	for _, elt := range Elts {
		ret = append(ret, ParseExpr(elt))
		// if keyVal, ok := elt.(*ast.KeyValueExpr); ok {
		// 	// todo: resolve independent
		// 	// ret = append(ret, "{ " + keyVal.Key.(*ast.BasicLit).Value + ", " + keyVal.Value.(*ast.BasicLit).Value + "},");
		// 	ret = append(ret, ParseKeyValueExpr(keyVal));
		// }
	}
	ret = append(ret, "};")
	return t, ret
}


func ParseArrayType(Type *ast.ArrayType, Elts []ast.Expr) (tname string, values []string) {

	values = append(values, "{");
	if ident, ok := Type.Elt.(*ast.Ident); ok {
		tname = ident.Name
	} else {
		log.Fatal("INVALID ELEMENT")
	}

	for _, elt := range Elts {
		values = append(values, ParseExpr(elt) + ",")
		// if bl, ok := elt.(*ast.BasicLit); ok {
		// 	// todo: resolve independent
		// 	values = append(values, bl.Value + ", ");
		// } else {
		// 	log.Fatal("invalid array values")
		// }
	}
	values = append(values, "};")
	return tname, values
}

type type_value struct {
	var_type int
	tname string
	values[]string
}

func genTypeValues(name string, tvalue type_value) string {
	var ret string
	switch tvalue.var_type {
	case VAR_ARRAY_TYPE:
		ret = tvalue.tname + " " + name + "[]" + strings.Join(tvalue.values, "\n")
		break
	case VAR_MAP_TYPE:
		ret = tvalue.tname + " " + name + strings.Join(tvalue.values, "\n");
		break
	case VAR_NORMAL_TYPE:
		ret = tvalue.tname + " " + name + "=" + strings.Join(tvalue.values, "\n")
		break
	case VAR_NONE_TYPE:
		_, file, line, _ := runtime.Caller(0)
		log.Fatal("INVALID VAR TYPE" + file + ", lien: " + strconv.Itoa(line))
	}

	return ret
}

func ParseVar(decl *ast.GenDecl) []string {
	var ret []string
	var names []string

	var tvalues []type_value
	for _, spec := range decl.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			for _, name := range vs.Names {
				names = append(names, name.Name)
			}
			for _, value := range vs.Values {
				if cl, ok := value.(*ast.CompositeLit); ok {
					typ, tname, vals := ParseCompositeLit(cl)
					tv := type_value{typ, tname, vals}
					tvalues = append(tvalues, tv)
				} else {
					typ := VAR_NORMAL_TYPE
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
			} else if (len(names) == 1 && len(tvalues) > 0) {
				log.Fatal("NOT IMPLEMENTED YET")
			} else {
				log.Fatal("invalid names and values size")
			}
		} else {
			log.Fatal("NOT IMPLEMENTED YET")
		}
	}

	return ret
}