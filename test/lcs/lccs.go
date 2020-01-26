package main

import (
	"fmt"
)

func xmain() {
	s1 := "This is illegal"
	s2 := "Totally different"
	fmt.Printf("This: " + lccs("", s1, s2))
}

func lccs(prefix string, s1 string, s2 string) string {
	if s1 == "" || s2 == "" {
		return prefix
	}
	x1 := ""
	if s1[0] == s2[0] {
		x1 = lccs(prefix+s1[0:1], s1[1:], s2[1:])
	} else {
		x1 = prefix
	}

	x2 := lccs("", s1[1:], s2)
	x3 := lccs("", s1, s2[1:])
	if len(x1) > len(x2) && len(x1) > len(x3) {
		return x1
	}
	if len(x2) > len(x1) && len(x2) > len(x3) {
		return x2
	}
	return x3
}
