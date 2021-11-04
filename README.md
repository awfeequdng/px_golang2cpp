# px_golang2cpp
convert go to cpp
- convert map to std::unordered_map
- convert array to vector
- convert 'var test bool' to 'bool test{}'
- convert 'var test = 2' to 'auto test = 2;'
- convert 'var test []string' to 'vector<string> test{};'
- convert 'var test map[int]string' to std::unordered_map<int,string> test{};

# usage
### clone source code
- git clone git@github.com:awfeequdng/px_golang2cpp.git
### build the project
- 1. cd px_golang2cpp source directory
- 2. exec the cmd: `go build *.go`,then execuable file `main` is created

### generate a cpp file from go source code
- 1. we have test file in directory `test`, we now chose the `test/func.go`
- 2. execute the command: `./main test/func.go -o test.cc`, then we get a cpp file: `test.cc`

### build the created cpp source file
- we now have the test.cc file, then we can execute the command like this: `g++ test.cc`
- if there have no error in building the cpp file `test.cc`, then we get the correct cpp source form go

### notice
- the created cpp source file is just for testing, and some converted code build failed dure to lack of dependence file.

# example
- convert var in go to cpp: `./main test/var_test.go -o var_test.cc`

// var_test.go
```
package var_test

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

```

// var_test.cc (this is part of code)
```
std::vector<string> endings{
"{",
",",
"}",
":",
};
auto  test = (-1);
auto  test1 = 2;
auto  g1 = 23;
auto  g2 = "fjdd";
auto  switchExpressionCounter = (-1);
bool firstCase{};
string switchLabel{};
int labelCounter{};
int iotaNumber{};
int deferCounter{};
bool unfinishedDeferFunction{};
```
notice: cpp code not format yet(we will do it later).

- convert go map to cpp unordered_map: `./main test/map_test.go -o map_test.cc`

// map_test.go
```
package map_test

var includeFileMap map[string]string = make(map[string]string)

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

var testMap map[string]string
```

// map_test.cc (this is part of code)
```
std::unordered_map<string,string> includeFileMap = {};
std::unordered_map<string,string> includeMap{
{"std::tuple" , "tuple"},
{"std::endl" , "iostream"},
{"std::cout" , "iostream"},
{"std::string" , "string"},
{"std::size" , "iterator"},
{"std::unordered_map" , "unordered_map"},
{"std::hash" , "functional"},
{"std::size_t" , "cstddef"},
{"std::int8_t" , "cinttypes"},
{"std::int16_t" , "cinttypes"},
{"std::int32_t" , "cinttypes"},
{"std::int64_t" , "cinttypes"},
{"std::uint8_t" , "cinttypes"},
{"std::uint16_t" , "cinttypes"},
{"std::uint32_t" , "cinttypes"},
{"std::uint64_t" , "cinttypes"},
{"printf" , "cstdio"},
{"fprintf" , "cstdio"},
{"sprintf" , "cstdio"},
{"snprintf" , "cstdio"},
{"std::stringstream" , "sstream"},
{"std::is_pointer" , "type_traits"},
{"std::experimental::is_detected_v" , "experimental/type_traits"},
{"std::shared_ptr" , "memory"},
{"std::nullopt" , "optional"},
{"EXIT_SUCCESS" , "cstdlib"},
{"EXIT_FAILURE" , "cstdlib"},
{"std::vector" , "vector"},
{"std::unique_ptr" , "memory"},
{"std::runtime_error" , "stdexcept"},
{"std::regex_replace" , "regex"},
{"std::regex_constants" , "regex"},
{"std::to_string" , "string"},
};
std::unordered_map<string,string> testMap{};
```

notice: not format yet