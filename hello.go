package main

import "fmt"

// if we want to declare a variable at the packagae level then we cant use the := syntax.

var a int=24

func main() {
    fmt.Println("Hello, World!");
    var i int // just declaring the variable
    i=42
    fmt.Println(i)
    //or 
    var j int =27
    fmt.Println(j)
    //or
    k:=99
    fmt.Println(k)
    // printing the value and type of the variable
    fmt.Printf("%v, %T",j,j)

    //a will be accessible which is defined outside the main().
    fmt.Println(a)

}