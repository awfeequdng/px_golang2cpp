package main

import (
	"go/ast"
)

var sliceTemplate2 = "template<typename T>\n" +
	"std::vector<T> slice(std::vector<T> const &v, int m, int n)\n" +
	"{\n" +
	"    auto first = v.cbegin() + m;\n" +
	"    auto last = v.cbegin() + n;\n" +
	"\n" +
	"    std::vector<T> vec(first, last);\n" +
	"    return vec;\n" +
	"}\n"

var sliceTemplate1 = "template<typename T>\n" +
	"std::vector<T> slice(std::vector<T> &v, int m, int n)\n" +
	"{\n" +
	"    std::vector<T> vec(n - m);\n" +
	"    std::copy(v.begin() + m, v.begin() + n, vec.begin());\n" +
	"    return vec;\n" +
	"}\n"

var sliceTemplate = "template<typename T>\n" +
	"std::vector<T> slice(std::vector<T> &v, std::size_t low, std::size_t high = -1)\n" +
	"{\n" +
	"    std::vector<T> vec;\n" +
	"    if (high == -1) {\n" +
	"    	std::copy(v.begin() + low, v.end(), std::back_inserter(vec));\n" +
	" 	 } else {\n" +
	"	 	std::copy(v.begin() + low, v.begin() + high, std::back_inserter(vec));\n" +
	"	 }\n" +
	"    return vec;\n" +
	"}\n"

func GetSliceTemplate() string {
	return sliceTemplate
}

func ParseSliceExpr(expr *ast.SliceExpr) string {
	x := ParseExpr(expr.X)
	var low string = "0"
	if expr.Low != nil {
		low = ParseExpr(expr.Low)
	}
	var high string = "-1"
	if expr.High != nil {
		high = ParseExpr(expr.High)
	}
	//var warn = "// TODO: Warning: this code is convert from go slice, it may get error in c++"
	return "slice(" + x + ", " + low + ", " + high + ")"
}
