package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type ResOutput struct {
	URL     string
	Headers string
	Body    string
}

func main() {
	var out ResOutput
	output := flag.Bool("o", false, "Outputs response to a file of choice")
	post := flag.Bool("post", false, "Sends a POST method request")
	put := flag.Bool("put", false, "Sends a PUT method request")
	delete := flag.Bool("delete", false, "Sends a DELETE method request")

	flag.Parse()

	if *post {
		out, _ = Bodied(os.Args[2], os.Args[3], http.MethodPost)
		if *output {
			WriteFile(out, os.Args[4])
			fmt.Printf("Output written to file %s", os.Args[4])
		} else {
			fmt.Print(out)
		}
	} else if *put {
		out, _ = Bodied(os.Args[2], os.Args[3], http.MethodPut)
		if *output {
			WriteFile(out, os.Args[4])
			fmt.Printf("Output written to file %s", os.Args[4])
		} else {
			fmt.Print(out)
		}
	} else if *delete {
		if os.Args[1] == "-o" {
			out = Unbodied(os.Args[2], http.MethodDelete)
		} else {
			out = Unbodied(os.Args[1], http.MethodDelete)
		}
		if *output {
			WriteFile(out, os.Args[3])
			fmt.Printf("Output written to file %s", os.Args[3])
		} else {
			fmt.Print(out)
		}
	} else {
		if os.Args[1] == "-o" {
			out = Unbodied(os.Args[2], http.MethodGet)
		} else {
			out = Unbodied(os.Args[1], http.MethodGet)
		}
		if *output {
			WriteFile(out, os.Args[3])
			fmt.Printf("Output written to file %s", os.Args[3])
		} else {
			fmt.Print(out)
		}
	}
}
