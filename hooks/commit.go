package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	getUnstagetFiles()
}

func getUnstagetFiles() {
	// git diff --cached --name-only | grep `.go$`
	rawCmd := "git diff --cached --name-only | grep '.go$'"
	buildCmd := exec.Command("bash", "-c", rawCmd)
	msg, err := buildCmd.Output()
	sl := strings.Split(string(msg), "\n")
	fmt.Println(sl)
	if err != nil {
		panic(err)
	}
}
