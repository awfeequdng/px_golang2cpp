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
- convert var in go to cpp: 
// var_test.go
```
  var (
	switchExpressionCounter = -1
	firstCase               bool
	switchLabel             string
	labelCounter            int
	iotaNumber              int // used for simple increases of iota constants
	deferCounter            int
	unfinishedDeferFunction bool
)
```
// var_test.cc
```
auto  switchExpressionCounter = (-1);
bool firstCase{};
string switchLabel{};
int labelCounter{};
int iotaNumber{};
int deferCounter{};
bool unfinishedDeferFunction{};
``  
