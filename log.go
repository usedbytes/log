// SPDX-License-Identifier: MIT
// Copyright (c) 2020 Brian Starkey <stark3y@gmail.com>
package log

import (
	"fmt"
	"os"
	"log"
)

var verbose bool
var uselog bool = true

func SetVerbose(v bool) {
	verbose = v
}

func SetUseLog(v bool) {
	uselog = v
}

func Fatalln(args ...interface{}) {
	Println(args)
	os.Exit(1)
}

func Println(args ...interface{}) {
	if uselog {
		log.Println(args...)
	} else {
		fmt.Println(args...)
	}
}

func Print(args ...interface{}) {
	if uselog {
		log.Print(args...)
	} else {
		fmt.Print(args...)
	}
}

func Printf(args ...interface{}) {
	if uselog {
		log.Printf(args[0].(string), args[1:]...)
	} else {
		fmt.Printf(args[0].(string), args[1:]...)
	}
}

func Verbose(args ...interface{}) {
	if !verbose {
		return
	}
	Print(append([]interface{}{"VERBOSE:"}, args...)...)
}

func Verboseln(args ...interface{}) {
	if !verbose {
		return
	}
	Println(append([]interface{}{"VERBOSE:"}, args...)...)
}

func Verbosef(args ...interface{}) {
	if !verbose {
		return
	}
	Printf(append([]interface{}{"VERBOSE: " + args[0].(string)}, args[1:]...)...)
}
