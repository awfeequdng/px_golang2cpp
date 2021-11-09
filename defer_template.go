package main

var deferTemplate = "// C++ implementation of golang defer keyword, the concept of delayed execution, often used for resource release\n" +
	"template <typename FnType>\n" +
	"class ScopedLambda {\n" +
	"public:\n" +
	"    explicit ScopedLambda(FnType fn) : fn_(std::move(fn)), active_(true)\n" +
	"    {}\n" +
	"    // Default movable." +
	"\n" +
	"    ScopedLambda(ScopedLambda&&) = default;\n" +
	"    ScopedLambda& operator=(ScopedLambda&&) = default;\n" +
	"    // Non-copyable. In particular, there is no good reasoning about which copy\n" +
	"    // remains active.\n" +
	"    ScopedLambda(const ScopedLambda&) = delete;\n" +
	"    ScopedLambda& operator=(const ScopedLambda&) = delete;\n" +
	"    ~ScopedLambda() {\n" +
	"        if (active_)\n" +
	"            fn_();\n" +
	"    }\n" +
	"    void run_and_expire() {\n" +
	"        if (active_)\n" +
	"            fn_();\n" +
	"        active_ = false;\n" +
	"    }\n" +
	"    void activate() {\n" +
	"        active_ = true;\n" +
	"    }\n" +
	"    void deactivate() {\n" +
	"        active_ = false;\n" +
	"    }\n" +
	"\n" +
	"private:\n" +
	"    FnType fn_;\n" +
	"    bool active_ = true;\n" +
	"};\n" +
	"\n" +
	"template <typename FnType>\n" +
	"ScopedLambda<FnType> MakeScopedLambda(FnType fn) {\n" +
	"    return ScopedLambda<FnType>(std::move(fn));\n" +
	"}\n" +
	"\n" +
	"#define TOKEN_PASTE(x, y) x##y\n" +
	"#define TOKEN_PASTE2(x, y) TOKEN_PASTE(x, y)\n" +
	"#define SCOPE_UNIQUE_NAME(name) TOKEN_PASTE2(name, __LINE__)\n" +
	"#define NAMED_DEFER(name, ...) auto name = MakeScopedLambda([&] { __VA_ARGS__; })\n" +
	"#define DEFER(...) NAMED_DEFER(SCOPE_UNIQUE_NAME(defer_varname), __VA_ARGS__)\n"

func GetDeferTemplte() string {
	return deferTemplate
}