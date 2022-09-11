package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var age int64

// 方式3。
//var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// 方式2。
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	// 方式3。
	//cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.Int64Var(&age, "age", 1, "年龄")
}

func main() {
	// 方式1。
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	// 方式3。
	//cmdLine.Parse(os.Args[1:])
	flag.Parse()
	//fmt.Println(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("age:, %d!\n", age)
}
