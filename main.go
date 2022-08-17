package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os/exec"
)

func printLog(stdout io.ReadCloser) {
	reader := bufio.NewReader(stdout)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Print(readString)
	}
}

func appserver(w http.ResponseWriter, req *http.Request) {

	appserverShellPath := "/data/hswg/deploy.sh"
	cmd := exec.Command(appserverShellPath)
	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	go printLog(stdout)
}

func main() {
	http.HandleFunc("/appserver", appserver)
	http.ListenAndServe(":23998", nil)
}
