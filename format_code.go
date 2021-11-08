package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func FormatCode(code string) string {
	//cmd := exec.Command("clang-format", "-style={BasedOnStyle: Webkit, ColumnLimit: 120}")
	cmd := exec.Command("clang-format", "-style={Language: Cpp, BasedOnStyle: Google, ColumnLimit: 120, " +
		"AccessModifierOffset: -4, IndentWidth: 4, PointerAlignment: Right}")
	cmd.Stdin = strings.NewReader(code)
	var out bytes.Buffer
	cmd.Stdout = &out

	var err error
	err = cmd.Run()
	if err != nil {
		log.Println("clang-format is not available, the output will look ugly!")
		return code
	} else {
		return out.String()
	}
}
