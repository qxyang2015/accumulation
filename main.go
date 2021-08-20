package main

import (
	"bytes"
	"fmt"
	"github.com/qxyang2015/accumulation/tools/http_tools"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start")
	path := "C:\\Users\\qxyan\\Desktop\\1.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	w := &bytes.Buffer{}
	writen, err := io.Copy(w, file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(writen)
	url := "localhost:9099/web/formdata"

	res, err := http_tools.HttpWhithOpt(url, http.MethodPost, "",
		http_tools.TimeOut(10),
		http_tools.OptWrite(w))
	fmt.Println(string(res), err)
	fmt.Println("end")
}
