package wordtrek

import (
    "fmt"
    "util"
)

type WTNode struct{ 
	words []string
	wordLengths []int
	level int
	table util.Table
}

func (wt *WTNode) Clone() WTNode {

    wl := make([]int, len(wt.wordLengths))
    copy(wl, wt.wordLengths)

    w := make([]string, len(wt.words))
    copy(w, wt.words)

    return WTNode {
        words: w,
        wordLengths: wl,
        level: wt.level,
        table: wt.table.Clone()}
}

func (wt *WTNode) Print() {
    fmt.Printf("%d: %d, %s\n", wt.level, wt.wordLengths, wt.words)
    wt.table.Print()
}
