package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    ret := make([][]uint8, dy)

    for x := 0; x < dy; x++ {
        ret[x] = make([]uint8, dx)

        for y := 0; y < dx; y++ {
            ret[x][y] = uint8(x ^ y)
        }
    }

    return ret
}

func main() {
    pic.Show(Pic)
}
