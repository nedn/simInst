package main

import (
    "io/ioutil"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        println("Usage: ./simInst <source file> <outputfile>")
        return
    }

    // open input file
    sourceByte, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        panic(err)
    }
    print(compile(string(sourceByte)))
}

func compile(source string) string {
    return source
}
