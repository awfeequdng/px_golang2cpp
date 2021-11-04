package main

import (
	"go/ast"
	"strings"
)

var includeMap = map[string]string{
	"std::tuple":                       "tuple",
	"std::endl":                        "iostream",
	"std::cout":                        "iostream",
	"std::string":                      "string",
	"std::size":                        "iterator",
	"std::unordered_map":               "unordered_map",
	"std::hash":                        "functional",
	"std::size_t":                      "cstddef",
	"std::int8_t":                      "cinttypes",
	"std::int16_t":                     "cinttypes",
	"std::int32_t":                     "cinttypes",
	"std::int64_t":                     "cinttypes",
	"std::uint8_t":                     "cinttypes",
	"std::uint16_t":                    "cinttypes",
	"std::uint32_t":                    "cinttypes",
	"std::uint64_t":                    "cinttypes",
	"printf":                           "cstdio",
	"fprintf":                          "cstdio",
	"sprintf":                          "cstdio",
	"snprintf":                         "cstdio",
	"std::stringstream":                "sstream",
	"std::is_pointer":                  "type_traits",
	"std::experimental::is_detected_v": "experimental/type_traits",
	"std::shared_ptr":                  "memory",
	"std::nullopt":                     "optional",
	"EXIT_SUCCESS":                     "cstdlib",
	"EXIT_FAILURE":                     "cstdlib",
	"std::vector":                      "vector",
	"std::unique_ptr":                  "memory",
	"std::runtime_error":               "stdexcept",
	"std::regex_replace":               "regex",
	"std::regex_constants":             "regex",
	"std::to_string":                   "string",
	// TODO: complex64, complex128
}

var endings = []string{"{", ",", "}", ":"}

var test = -1
var test1 = 2

var g1, g2 = 23, "fjdd"

var (
	switchExpressionCounter = -1
	firstCase               bool
	switchLabel             string
	labelCounter            int
	iotaNumber              int // used for simple increases of iota constants
	deferCounter            int
	unfinishedDeferFunction bool
)

var globalObjectMap map[string]string

func ParseFuncDecl(decl *ast.FuncDecl) []string {
	var ret []string
	name := decl.Name.Name
	func_type := decl.Type
	params := ParseFieldList(func_type.Params)
	results := ParseFieldList(func_type.Results)
	if len(results) == 0 {
		ret = append(ret, "void " + name + "(" + strings.Join(params, " ") + ");")
	}
	return ret
}
