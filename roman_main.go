package main

import (
	"fmt"
	"roman"
)

func main() {

	for {
		//var scanner = bufio.NewScanner(os.Stdin)

		var num int
        fmt.Print("Enter a whole number less than 4000: ")
		fmt.Scan(&num)
		r := roman.ArabicToRoman(num)
		fmt.Printf("%d -> %s -> %d\n", num, r, roman.RomanToArabic(r))
	}
}
