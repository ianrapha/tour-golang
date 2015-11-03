package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    countMap := map[string]int{}
    arrString := strings.Fields(s)

    for i := 0; i < len(arrString); i++ {
        if ok := countMap[arrString[i]]; ok > 0 {
            countMap[arrString[i]]++
        } else {
            countMap[arrString[i]] = 1
        }
    }

    return countMap
}

func main() {
    wc.Test(WordCount)
}