package wordtrek

import (
    "fmt"
    "dict"
    "util"
    "time"
)

type WordTrek struct {
    visited map[string]bool
    wtc chan WTNode
    dict dict.Dict
    rows, cols int
    wordLengths []int
}

func (wt *WordTrek) Solve(tableStr string, wordLengths []int, dict dict.Dict) {

    t := util.StrToTable(tableStr)
    wt.rows, wt.cols = len(t), len(t[0])
    wt.wtc = make(chan WTNode) 
    wt.wordLengths = wordLengths;
    wt.dict = dict
    wt.visited = make(map[string]bool)

    wtn := WTNode{
        words: []string{}, 
        wordLengths: wordLengths,
        table: t }

    wtn.Print(true)
    go func(){ wt.wtc <- wtn }()

    for {
        select {
        case wtn := <- wt.wtc: 
            key := wtn.Key()
            if _, found := wt.visited[key]; !found {
                wt.visited[key] = true
                go wt.findWord(wtn)
            }
        // terminate if there is no node available for 1 second!
        case <-time.After(1 * time.Second):
            fmt.Println("No activity in 1 second. That's all folks!")
            return
        }
    }
}

func (wt *WordTrek) findWord(wtn WTNode) {

    //fmt.Printf("findWord %c, %s\n", wtn.table, wtn.words);
    if len(wt.wordLengths) == len(wtn.words) {
        wtn.Print(false)
        return
    }

    t := wtn.table
    for r:=0; r<len(t); r++ {
        for c:=0; c<len(t[0]); c++ {
            go wt.findWordAtRC(wtn.Clone(), r, c, "") 
        }
    }
}

func (wt *WordTrek) findWordAtRC(wtn WTNode, row int, col int, prefix string) {
    //fmt.Printf("findWordsAtRC: %c, %d, %d, %d, %q\n", wtn.table, row, col, prefix)

    t := wtn.table

    ch := t[row][col]

    if ch == '.' {
        return
    } 
    t[row][col] = '.'

    if len(prefix) == wt.wordLengths[len(wtn.words)] - 1 {
        if wt.dict.Look(prefix + string(ch)) {
            wt.wtc <- WTNode{
                words: append(wtn.words, prefix + string(ch)), 
                table: wtn.table.Clone().DropDown() }
        }
        t[row][col] = ch
        return
    }

    for ri := -1; ri<2; ri++ {
        r := row + ri
        if r >= 0 && r < wt.rows {
            for ci := -1; ci<2; ci++ {
                c := col + ci
                if c >= 0 && c < wt.cols {
                    wt.findWordAtRC(wtn, r, c, prefix + string(ch))
                }
            }
        }
    }

    t[row][col] = ch
}
