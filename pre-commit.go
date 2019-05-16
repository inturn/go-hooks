package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const hookPath = ".git/hooks/"

func main() {
	fs := flag.NewFlagSet("go-hooks", flag.ExitOnError)

	url := fs.String("url",
		"https://raw.githubusercontent.com/inturn/go-hooks/BE-1904_git-hook_solution_for_golang_microservices/operate/commit.go",
		"url of the commit operating file")
	online := fs.Bool("online", true, "is online")

	_ = online

	getLatestPush(*url)
}

func getLatestPush(url string) {

	resp := sendRequest(url)

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Print(os.TempDir())

	file, err := os.Create(hookPath + "commit")

	if err != nil {
		panic(err)
	}
	_, err = file.Write(raw)
	if err != nil {
		panic(err)
	}
	defer file.Close()

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

func compareHashes() {
	h := md5.New()

}
