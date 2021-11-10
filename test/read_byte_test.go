package main

import "unicode"

func (r *reader) readByte() (ch rune) {
	ch = r.peek()
	if ch == unicode.ReplacementChar && r.eof() {
		return
	}
	r.inc()
	return
}
