package http_tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type reqOption func(r *http.Request)

func OptCookie(cookie *http.Cookie) reqOption {
	return func(r *http.Request) {
		r.AddCookie(cookie)
	}
}

func OptHeader(headerMap map[string]string) reqOption {
	return func(r *http.Request) {
		for k, v := range headerMap {
			r.Header.Set(k, v)
		}
	}
}

func OptUrlQuery(paramMap map[string]string) reqOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		for k, v := range paramMap {
			q.Add(k, v)
		}
		r.URL.RawQuery = q.Encode()
	}
}

func OptWrite(w io.Writer) reqOption {
	return func(r *http.Request) {
		r.Write(w)
	}
}

type clientFuntion func(c *http.Client)

func TimeOut(t time.Duration) clientFuntion {
	return func(c *http.Client) {
		c.Timeout = t * time.Second
	}
}

func HttpWhithOpt(url, method, data string, clientFunc clientFuntion, reqOpt ...reqOption) (result []byte, err error) {
	if method != http.MethodPost && method != http.MethodGet {
		return nil, fmt.Errorf("不支持该请求方式[%v]", method)
	}
	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return
	}
	for _, rf := range reqOpt {
		rf(req)
	}
	//设置超时时长
	client := &http.Client{}
	clientFunc(client)
	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status[%v] is not StatusOK", response.Status)
		return
	}
	return
}
