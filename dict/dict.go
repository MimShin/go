package dict

import (
    "fmt"
    "bufio"
    "os"
    "sort"
    "strings"
    "regexp"
)

type Dict struct{
    filename string
    words []string
}

func (d* Dict) ReadFile(dictFile string) {

    d.filename = dictFile
    f, err := os.Open(dictFile)
    if err != nil {
        panic(err)
    }

    reader := bufio.NewReader(f);

    re, err := regexp.Compile(`^[a-z]+\n$`)
    if err != nil {
        panic(err)
    }

    for {
        str, err :=  reader.ReadString('\n');
        if (err != nil) {
            break
        }
        strl := strings.ToLower(str)
        if re.MatchString(strl) {
            d.words = append(d.words, strings.TrimSuffix(strl, "\n"))
        }
    }

    if len(d.words) >  0 {
        sort.Strings(d.words)
        fmt.Printf("%d words imported: %s...%s\n", len(d.words), d.words[0], d.words[len(d.words)-1])
    }
}

func (d* Dict) Look(w string) bool {
    w = strings.ToLower(w)
    idx := sort.SearchStrings(d.words, w)
    //fmt.Println(idx, w, d.words[idx])
    if idx < len(d.words) && d.words[idx] == w {
        return true
    }

    return false
}
