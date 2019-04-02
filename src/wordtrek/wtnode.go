package wordtrek

import (
    "fmt"
    "util"
    "sort"
    "strings"
)

type WTNode struct{ 
    words []string
    wordLengths []int
    table util.Table
}

func (wtn *WTNode) Clone() WTNode {

    wl := make([]int, len(wtn.wordLengths))
    copy(wl, wtn.wordLengths)

    w := make([]string, len(wtn.words))
    copy(w, wtn.words)

    return WTNode {
        words: w,
        wordLengths: wl,
        table: wtn.table.Clone()}
}

func (wtn *WTNode) Key() string {
    s := make([]string, len(wtn.words))
    copy(s, wtn.words)
    sort.Strings(s)
    return wtn.table.ToStr() + strings.Join(s, "")
}

func (wtn *WTNode) Print(printTable bool) {
    fmt.Printf("%s, %d\n", wtn.words, wtn.wordLengths)
    if printTable {
        wtn.table.Print()
    }
}
