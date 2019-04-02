package main

import "fmt"

func check_slice(x []int) {
    z := make([]int, 5)
    if z != nil && len(x) != 0 {
        fmt.Println(len(x))
    }
}

func main() {
    var s []int

    a := [5]int{1, 2, 3, 4, 5}

    s = a[2:4]
    check_slice(s)
}
