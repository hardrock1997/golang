package main

import "fmt"

// if we want to declare a variable at the packagae level then we cant use the := syntax.

var a int=24

func main() {
    // we cannot define a same variable in the same scope.
    //eg->
    var a int=30
    // a:=34 // cannot re declare the same variable a inside the same scope (here scope is the main function)
    // But we can re assign the same variable inside the same scope
    a=45
    fmt.Println(a)

}