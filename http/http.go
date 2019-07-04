package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func Post(url string, params map[string]string, body string) (string, error) {
	resp, err := http.Post(url + ConvertToQueryParams(params), "application/json", strings.NewReader(body));
	return ResponseHandle(resp, err)
}

func Get(url string, params map[string]string)(string, error){
	resp, err := http.Get(url + ConvertToQueryParams(params))
	return ResponseHandle(resp, err)
}

func ResponseHandle(resp *http.Response, err error) (string, error){
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ConvertToQueryParams(params map[string]string) string{
	if &params == nil || len(params) == 0{
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString("?")
	for k, v := range params {
		buffer.WriteString(k + "=" + v + "&")
	}
	buffer.Truncate(buffer.Len() - 1)
	return buffer.String()
}