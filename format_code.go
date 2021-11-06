package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func FormatCode(code string) string {
	//cmd := exec.Command("clang-format", "-style={BasedOnStyle: Webkit, ColumnLimit: 120}")
	cmd := exec.Command("clang-format", "-style={BasedOnStyle: Google, ColumnLimit: 120}")
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
