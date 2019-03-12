package main

import (
    "fmt"
    "time"
    "sudoku"
)

func main() {

    start := time.Now()
    //s := sudoku.Init("123456749 234567891 345678912 456789123 567891234 678912345 789123456 891234567 9123456789")
    s := sudoku.Init("53..7.... 6..195... .98....6. 8...6...3 4..8.3..1 7...2...6 .6....28. ...419..5 ....8..79")
    s.Print()
    /*
    s.CheckCell(0, 0)
    */
    fmt.Printf("Elapsed time: %s\n", time.Since(start));
}
