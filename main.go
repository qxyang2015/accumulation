package main

import (
	"strings"
	"text/scanner"
)

func main() {
	src := "a+b"

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	tok := s.Scan()
	for tok != scanner.EOF {
		// do something with tok
		tok = s.Scan()
	}
}
