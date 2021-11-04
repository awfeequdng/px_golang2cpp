package main

import "go/ast"

func ParseBasicLit(bl *ast.BasicLit) string {
	return bl.Value
}