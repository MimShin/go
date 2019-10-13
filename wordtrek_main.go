package main

import (
    "fmt"
    "dict"
    "time"
    "wordtrek"
)

func main() {

	// Load the dictionary
    d := dict.Dict{}
    d.ReadFile("/usr/share/dict/words")

    s := ""
    wl := []int{}
    
    s = "ce ak"
    wl = []int{4}

    s = "enee idgk bwua orfc"
    wl = []int{4, 5, 7}

    wl = []int{4, 6, 6}; 
    s = "wses ogau rlnr bcef";

    wl = []int{3, 6, 7, 9};
    s = "mgeee oarpf resaf ntpwo ninec"; 

    wl = []int{5, 4, 9, 7};
    s = "adlgb nedtn oiawi giuqa rbsat"; 
    
    wl = []int{7, 4, 8, 3, 3};
    s = "brigp eibfe ogaba lcbec cotea"; 
    
    wl = []int{4, 3, 4, 3, 4, 3, 4};
    s = "xflga oadtk vblbo oooob glbol"; 

    wl = []int{3, 6, 6, 5, 5}
    s = "ydmcr erate ajjte nepub hosui"

    wl = []int{4, 6, 5, 5, 5}
    s = "ynfgr eocla ljati gesri udeht"

    wl = []int{5, 4, 6, 10}
    s = "rtptr oelep ceaco ktnie ikehl"

    wl = []int{9, 8, 4, 4}
    s = "rdrep eagpa nanws izext eamet"

    wt := wordtrek.WordTrek{}
    start := time.Now()
    wt.Solve(s, wl, d);
    fmt.Printf("Elapsed time: %s\n", time.Since(start));
}
