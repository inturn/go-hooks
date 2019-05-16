package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)



func main() {
	fs := flag.NewFlagSet("go-hooks", flag.ExitOnError)

	url := fs.String("url",
		"https://raw.githubusercontent.com/inturn/go-hooks/BE-1904_git-hook_solution_for_golang_microservices/operate/commit.go",
		"url of the commit operating file")
	online := fs.Bool("online", true, "is online")

	_ = online

	getLatestPreCommit(*url)
	execute()
}

func execute() {

}

func getLatestPreCommit(url string) {

	resp, err := sendRequest(url)

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// do not forget to close body
	defer resp.Body.Close()

	// if file does not exist we skip the step with comparing the hashes and save file directly
	if _, err := os.Stat(hookPath + hookName); os.IsNotExist(err) {
		file, err := os.Create(hookPath + hookName)
		if err != nil {
			panic(err)
		}
		_, err = file.Write(rawBody)
		if err != nil {
			panic(err)
		}

		file.Close()

		cmd := exec.Command("chmod", "a+x", hookPath+hookName)
		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Print(string(out))
	}

	//fmt.Print(os.TempDir())

	//cmd := exec.Command("pwd")
	//data, err := cmd.Output()

	//fmt.Print(string(data))
}

// send request
func sendRequest(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func compareHashes(rawBody []byte) bool {
	fileHash := md5.Sum([]byte("fsfsfdfdfd"))
	bodyHash := md5.Sum(rawBody)

	// in go we can compare two arrays of the same type
	// [16]byte is the md5 hash
	if (fileHash != bodyHash) {
		return false
	}

	return true
}
