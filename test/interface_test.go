package main

import (
	"fmt"
	"strconv"
)

//func (s *Scanner) Errorf(format string, a ...interface{}) (err error) {
//	str := fmt.Sprintf(format, a...)
//	val := s.r.s[s.lastScanOffset:]
//	var lenStr = ""
//	if len(val) > 2048 {
//		lenStr = "(total length " + strconv.Itoa(len(val)) + ")"
//		val = val[:2048]
//	}
//	err = fmt.Errorf("line %d column %d near \"%s\"%s %s",
//		s.r.p.Line, s.r.p.Col, val, str, lenStr)
//	return
//}

// SetDefaultValue sets the default value.
func (c *ColumnInfo) SetDefaultValue(value interface{}) error {
	c.DefaultValue = value
	return nil
}