package main

import "flag"

func init() {
	showVersion := flag.Bool("v", false, "show current version")
	if *showVersion {
		println("1.0.1")
	}
}
