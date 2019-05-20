package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	diffCmd = "git diff --cached --name-only"
)

// two options
// 1.
// run-all-files - run linters on all files with .go extension
// 2. (default)
// run linters only on staged to commit files
func main() {
	getUpstagedFiles()
}

func getUpstagedFiles() {
	// git diff --cached --name-only | grep `.go$`
	buildCmd := exec.Command("bash", "-c", diffCmd)
	msg, err := buildCmd.Output()
	sl := strings.Split(string(msg), "\n")
	fmt.Println(sl)
	if err != nil {
		panic(err)
	}
}

func getConfig() {

}
