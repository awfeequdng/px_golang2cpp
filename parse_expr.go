package main

import (
	"go/ast"
	"log"
)

func ParseExpr(expr ast.Expr) string {

	switch expr.(type) {
	// Expressions
	case *ast.BadExpr:
		log.Fatal("bad expr")
	case *ast.Ident:
		return ParseIdent(expr.(*ast.Ident))
	case *ast.BasicLit:
	// nothing to do
		return ParseBasicLit(expr.(*ast.BasicLit))
	case *ast.Ellipsis:
	// if n.Elt != nil {
	// 	Walk(v, n.Elt)
	// }

	case *ast.FuncLit:
	// Walk(v, n.Type)
	// Walk(v, n.Body)

	case *ast.CompositeLit:
		return ParseCompositeLite(expr.(*ast.CompositeLit))
		//log.Fatal("can not call composite lit here")
	case *ast.ParenExpr:
		parenExpr := expr.(*ast.ParenExpr)
		return "(" + ParseExpr(parenExpr.X) + ")"
	case *ast.SelectorExpr:
		return ParseSelectorExpr(expr.(*ast.SelectorExpr))

	case *ast.IndexExpr:
		return ParseIndexExpr(expr.(*ast.IndexExpr))
	// Walk(v, n.X)
	// Walk(v, n.Index)

	case *ast.SliceExpr:
		return ParseSliceExpr(expr.(*ast.SliceExpr))
	// Walk(v, n.X)
	// if n.Low != nil {
	// 	Walk(v, n.Low)
	// }
	// if n.High != nil {
	// 	Walk(v, n.High)
	// }
	// if n.Max != nil {
	// 	Walk(v, n.Max)
	// }

	case *ast.TypeAssertExpr:
	// Walk(v, n.X)
	// if n.Type != nil {
		// Walk(v, n.Type)
	// }

	case *ast.CallExpr:
		return ParseCallExpr(expr.(*ast.CallExpr))

	case *ast.StarExpr:
		return ParseStarExpr(expr.(*ast.StarExpr))
	// Walk(v, n.X)

	case *ast.UnaryExpr:
		return ParseUnaryExpr(expr.(*ast.UnaryExpr))
	// Walk(v, n.X)

	case *ast.BinaryExpr:
		return ParseBinaryExpr(expr.(*ast.BinaryExpr))
	// Walk(v, n.X)
	// Walk(v, n.Y)

	case *ast.KeyValueExpr:
		return ParseKeyValueExpr(expr.(*ast.KeyValueExpr))
	// Walk(v, n.Key)
	// Walk(v, n.Value)

	case *ast.StructType:
		return ParseStructType(expr.(*ast.StructType))
	case *ast.ArrayType:
		return ParseArrayType(expr.(*ast.ArrayType))
	case *ast.MapType:
		return ParseMapType(expr.(*ast.MapType))
	}


	return ""
}
