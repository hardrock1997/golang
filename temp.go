package main

import (
    "fmt"
    "strconv"
    "bytes"
)

// if we want to declare a variable at the packagae level then we cant use the := syntax.

var (
    fname string="yash"
    lname string="mundra"
)
var (
    age int=27
)

func main() {
    var buffer bytes.Buffer
    buffer.WriteString(fname)
    buffer.WriteString(" ")
    buffer.WriteString(lname)
    
    name:=buffer.String()
    fmt.Println(name)

    strAge:=strconv.Itoa(age)
    fmt.Println("my name is ",name," and my age is ",strAge)
    

}