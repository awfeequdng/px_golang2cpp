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
