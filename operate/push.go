package main

import "flag"

func main() {
	fs := flag.NewFlagSet("pre-push", flag.ExitOnError)

	check := fs.String("check", "", "url of the commit operating file")
	install := fs.String("install", "", "url of the commit operating file")
	run := fs.String("run", "", "url of the commit operating file")

	_ = online
}

