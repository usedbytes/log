// SPDX-License-Identifier: MIT
// Copyright (c) 2020 Brian Starkey <stark3y@gmail.com>
package log

import (
	"log"
)

var verbose bool

func SetVerbose(v bool) {
	verbose = v
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Println(args ...interface{}) {
	log.Println(args...)
}

func Printf(args ...interface{}) {
	log.Printf(args[0].(string), args[1:]...)
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
