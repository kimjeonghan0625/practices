package main

import (
    "fmt"
)

func ArrayInsertDemo() {
    var a [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
    var b [11]int
    copy(b[0:], a[0:5])
    b[5] = 100
    copy(b[6:], a[5:])
    fmt.Println(a)
    fmt.Println(b)
}

