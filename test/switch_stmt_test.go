package main

import (
	"fmt"
	"go/ast"
	"log"
	"strconv"
	"testing"
	"time"
)
//
//func TestSwitchStmt(t *testing.T) {
//	id := 12
//	switch id {
//	case 2:
//		log.Print("id is 2")
//	case 12:
//		log.Print("id is 12")
//	default:
//		log.Print("invalid id: " + strconv.Itoa(id))
//	}
//}
//
//func TestSwitchStmt1(t int) int {
//	switch id := 12; id  {
//	case 2:
//		t = id + 2
//	case 12:
//		t = id + 10
//		fmt.Println("Good morning!")
//	default:
//		t = id
//	}
//	t = t + 2
//	return t
//}
//
//func noConditionSwitch() {
//	hour := time.Now().Hour()
//	switch {
//	case hour < 12:
//		fmt.Println("Good morning!")
//		cc := 3
//	case hour < 17:
//		fmt.Println("Good afternoon!")
//		fmt.Println("Good afternoon1!")
//	default:
//		fmt.Println("Good evening!")
//	}
//}

func ParseFallThroughSwitchStmt(switch_stmt int) string {
	id := 2
	id *= 7
bbb:
	switch id {
	case 1:
		//log.Print("id = 1")
		goto bbb
		a := 2
		fallthrough
	case 2:
		//log.Print("id = 2")
		a := 3
		a += 4
		if id > 3 {
			id = 7
			break
		}
	default:
		id = 6
	}

	//return []string{""}
	bb:
	return ""

}
//
//func ParseNoConditionSwitchStmt(switch_stmt *ast.SwitchStmt) [] string {
//	var ret []string
//	switch test := 2; {
//	case test < 2:
//	case test < 5:
//		break
//	}
//	return ret
//}
