package main

import "fmt"

var a = 10

func printPointer() {
    fmt.Println(a)
    fmt.Println(&a)
}

