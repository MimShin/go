package main

import (
	"fmt"
)

func xmain() {
	s1 := "This is"
	s2 := "somthing totally different"
	fmt.Printf("This: " + lcs(s1, s2))
}

func lcs(s1 string, s2 string) string {
	if s1 == "" || s2 == "" {
		return ""
	}
	if s1[0] == s2[0] {
		return s1[0:1] + lcs(s1[1:], s2[1:])
	}
	x1 := lcs(s1[1:], s2)
	x2 := lcs(s1, s2[1:])
	if len(x1) > len(x2) {
		return x1
	}
	return x2
}
