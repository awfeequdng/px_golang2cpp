package main

import (
	"go/ast"
	"log"
	"path/filepath"
	"strings"
)

var importMap map[string]string = make(map[string]string)

func genImportMap(key *ast.Ident, pkg string) {
	var k string
	if key == nil {
		k = filepath.Base(pkg)
	} else {
		k = key.Name
	}

	if strings.HasPrefix(k, "\"") {
		k = k[1:]
	}
	if strings.HasSuffix(k, "\"") {
		k = k[:len(k) - 1]
	}
	importMap[k] = pkg
}

func PrintImportMap() {
	log.Print("importMap: \n")
	for k, v := range importMap {
		log.Print(k + " : " + v + "\n")
	}
}

func ParseImport(decl *ast.GenDecl) []string {
	var ret []string
	var names []string
	for _, spec := range decl.Specs {
		if is, ok := spec.(*ast.ImportSpec); ok {
			names = append(names, is.Path.Value)
			// generate import package map
			genImportMap(is.Name, is.Path.Value)
		} else {
			log.Fatal("can not convert to ImportSpec")
		}
	}
	for _, name := range names {
		// do not include this import pack
		ret = append(ret, "// #include " + name);
	}
	return ret
}