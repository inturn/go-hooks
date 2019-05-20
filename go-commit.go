package main

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	hookPath      = ".git/bootstrap/"
	cachePath     = hookPath + ".cache/"
	preCommitName = "pre-commit"

	prePushHook = "pre-push"
	pushName    = "push"
	commitName  = "commit"

	goFileExt = ".go"
)

// two options:
// go-bootstrap install --> install pre-commit and pre-push bootstrap into .git/bootstrap
// and also creates operational scripts in .cache/push .cache/commit

// When runnin go-bootstrap install multiply times in repo it will update
// .cache/push, .cache/commit, .git/bootstrap/pre-commit and .git/bootstrap/pre-push scripts

// The second option:
// pre-commit (default)
func main() {
	fs := flag.NewFlagSet("go-bootstrap", flag.ExitOnError)
	fs.String("install", "install", "install pre-commit and pre-push into .git/bootstrap")

	// [TEMP]
	preCommitUrl := fs.String("url",
		"https://raw.githubusercontent.com/inturn/go-bootstrap/BE-1904_git-hook_solution_for_golang_microservices/pre-commit.go",
		"url of the commit operating file")

	// [TEMP]
	prePushUrl := fs.String("url2",
		"https://raw.githubusercontent.com/inturn/go-bootstrap/BE-1904_git-hook_solution_for_golang_microservices/pre-push.go",
		"url of the push operating file")


	runAll := fs.Bool("run-all-files", true, "run linters on all files in repo")

	if *runAll {
		executeRunAll()
		return
	}

	installPreCommit(*preCommitUrl)
	installPrePush(*prePushUrl)

}

// cmd to run-all with pre-commit and then with pre-push
func executeRunAll() {

}

// go-commit install command
func installPreCommit(url string) {
	resp, err := sendRequest(url)
	if err != nil {
		panic(err)
	}
	// do not forget to close body
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			//skip
		}
	}()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// if file does not exist we skip the step with comparing the hashes and save file directly
	if _, err := os.Stat(hookPath + preCommitName); os.IsNotExist(err) {
		writeHookFromBody(rawBody)
	}

	// remove if file exist
	err = os.Remove(hookPath + preCommitName)
	if err != nil {
		panic(err)
	}

	//fmt.Print(os.TempDir())

	//cmd := exec.Command("pwd")
	//data, err := cmd.Output()

	//fmt.Print(string(data))
}

func installPrePush(url string) {
	resp, err := sendRequest(url)
	if err != nil {
		panic(err)
	}
	// do not forget to close body
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			//skip
		}
	}()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// if file does not exist we skip the step with comparing the hashes and save file directly
	if _, err := os.Stat(hookPath + preCommitName); os.IsNotExist(err) {
		writeHookFromBody(rawBody)
	}

	// remove if file exist
	err = os.Remove(hookPath + preCommitName)
	if err != nil {
		panic(err)
	}

	//fmt.Print(os.TempDir())

	//cmd := exec.Command("pwd")
	//data, err := cmd.Output()

	//fmt.Print(string(data))
}

func writeHookFromBody(rawBody []byte) {
	// create file with .go extension
	// .git/bootstrap/pre-pre-commit.go
	file, err := os.Create(hookPath + preCommitName + goFileExt)
	if err != nil {
		panic(err)
	}

	// write data to the file
	_, err = file.Write(rawBody)
	if err != nil {
		panic(err)
	}

	// close the fd
	err = file.Close()
	if err != nil {
		panic(err)
	}

	// go build -o
	// build the downloaded go file
	buildCmd := exec.Command("go", "build", hookPath+preCommitName+goFileExt)
	err = buildCmd.Start()
	if err != nil {
		panic(err)
	}

	// move file from current directory to the hook path
	moveCmd := exec.Command("mv", preCommitName, hookPath)
	err = moveCmd.Start()
	if err != nil {
		panic(err)
	}

	// remove old go source file
	// pre-pre-commit.go
	err = os.Remove(hookPath + preCommitName + goFileExt)
	if err != nil {
		panic(err)
	}

	// make the file executable (needed according git instructions)
	cmd := exec.Command("chmod", "a+x", hookPath+preCommitName)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
}

// send request, do not forget to close body
func sendRequest(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	// send request to get the file, not closing body here,
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// close the request body
	defer func() {
		err = req.Body.Close()
		if err != nil {

		}
	}()

	// return response, not closing body here
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
