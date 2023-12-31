package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/yosssi/gohtml"
)

func Unbodied(url string, method string) ResOutput {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalf("Failure to create request to url: %s", url)
	}

	req.Header.Set("Content-Type", "application/xml; charset=utf-8")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failure to connect to url: %s", url)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Could read the response body from url: %s", url)
	}
	defer res.Body.Close()

	out := ResOutput{
		URL:     " \n" + url + "\n",
		Headers: GetHeaders(fmt.Sprint(res.Header), []string{"map", "[", "]"}, []string{"", " ", " \n"}),
		Body:    gohtml.Format(string(resBody)) + "\n",
	}

	return out
}

func Bodied(url string, body string, method string) (ResOutput, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		ReqErr(url)
	}
	req.Header.Set("Content-Type", "application/xml; charset=utf-8")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ReqErr(url)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		ReqErr(url)
	}
	defer res.Body.Close()

	out := ResOutput{
		URL:     " \n" + url + "\n",
		Headers: GetHeaders(fmt.Sprint(res.Header), []string{"map", "[", "]"}, []string{"", " ", " \n"}),
		Body:    gohtml.Format(string(resBody)) + "\n",
	}

	return out, nil
}
