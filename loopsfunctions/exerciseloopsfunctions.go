package main

import (
    "fmt"
    "math")

const delta = 1e-10

func Newton(z, x float64) float64 {
    return z - (math.Pow(z,2) - x) / (2 * z)
}

func SqrtTenIterations(x float64) float64 {
    if x == 0 {
        return 0
    }

    z := float64(1)

    for i := 0; i < 10; i++ {
        z = Newton(z, x)
    }

    return z
}

func Sqrt(x float64) float64 {
    z := float64(1)
    ret := float64(0)

    for {
        z = Newton(z, x)

        if math.Abs(z - ret) < delta {
            break
        }

        ret = z
    }

    return ret
}

func main() {
    fmt.Println(SqrtTenIterations(2))
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
    fmt.Println("-")
    fmt.Println(SqrtTenIterations(3))
    fmt.Println(Sqrt(3))
    fmt.Println(math.Sqrt(3))
    fmt.Println("-")
    fmt.Println(SqrtTenIterations(5))
    fmt.Println(Sqrt(5))
    fmt.Println(math.Sqrt(5))
    fmt.Println("-")
    fmt.Println(SqrtTenIterations(7))
    fmt.Println(Sqrt(7))
    fmt.Println(math.Sqrt(7))
    fmt.Println("-")
    fmt.Println(SqrtTenIterations(10))
    fmt.Println(Sqrt(10))
    fmt.Println(math.Sqrt(10))
    fmt.Println("-")
    fmt.Println(SqrtTenIterations(11))
    fmt.Println(Sqrt(11))
    fmt.Println(math.Sqrt(11))
    fmt.Println("-")
    fmt.Println(SqrtTenIterations(99))
    fmt.Println(Sqrt(99))
    fmt.Println(math.Sqrt(99))
}
