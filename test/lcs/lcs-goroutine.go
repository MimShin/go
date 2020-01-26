package main

import (
	"fmt"
)

func main() {
	s1 := "Tyhyiys is sexy hi"
	s2 := "Txhxixs is something"
	//	s1 = "This is"
	//	s2 = "somthing totally different"
	fmt.Printf("This: " + lcs(s1, s2))
}

func lcs(s1 string, s2 string) string {
	ch := make(chan string)
	go lcsgo(s1, s2, ch)
	return <-ch
}

func lcsgo(s1 string, s2 string, ch chan string) {

	ch1 := make(chan string)
	defer close(ch1)

	if s1 == "" || s2 == "" {
		ch <- ""
		return
	}

	if s1[0] == s2[0] {
		go lcsgo(s1[1:], s2[1:], ch1)
		ch <- s1[0:1] + <-ch1
		return
	}

	go lcsgo(s1[1:], s2, ch1)
	go lcsgo(s1, s2[1:], ch1)

	x1 := <-ch1
	x2 := <-ch1
	if len(x1) > len(x2) {
		ch <- x1
	} else {
		ch <- x2
	}
}
