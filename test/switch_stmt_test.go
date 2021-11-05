package main

import (
	"go/ast"
	"log"
	"strconv"
	"testing"
)

func TestSwitchStmt(t *testing.T) {
	id := 12
	switch id {
	case 2:
		log.Print("id is 2")
	case 12:
		log.Print("id is 12")
	default:
		log.Print("invalid id: " + strconv.Itoa(id))
	}
}


func TestSwitchStmt1(t int) int {
	switch id := 12; id  {
	case 2:
		t = id + 2
	case 12:
		t = id + 10
	default:
		t = id
	}
	t = t + 2
	return t
}

