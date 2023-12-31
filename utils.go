package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func WriteFile(content ResOutput, file string) (string, error) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Could not create file: %s", file)
	}

	_, err = f.WriteString(fmt.Sprintf("%s\n %s", content.Headers, content.Body))
	if err != nil {
		log.Fatal("Could not write output to file")
	}

	err = f.Close()
	if err != nil {
		log.Fatal("Could not close writer")
	}

	return fmt.Sprintf("File: %s has been written to", file), err
}

func GetHeaders(headers string, oldStrings []string, newStrings []string) string {
	newString := []byte(headers)
	for i := range oldStrings {
		newString = bytes.Replace(newString, []byte(oldStrings[i]), []byte(newStrings[i]), -1)
	}

	return string(newString)
}

func ReqErr(url string) error {
	fmt.Print("There was an error, would you like to see the detailed log? y/n -- ")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read input defaulting to basic error:")
		log.Fatalf("Failure to create request to url: %s", url)
	}

	switch line {
	case "y":
		return err
	case "n":
		log.Fatalf("Failure to create post request to url: %s", url)
	}

	return nil
}

func ReadErr(url string) error {
	fmt.Print("There was an error, would you like to see the detailed log? y/n -- ")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Could not read input defaulting to basic error:")
		log.Fatalf("Failure to read request body: %s", url)
	}

	switch line {
	case "y":
		return err
	case "n":
		log.Fatalf("Failure to read request body: %s", url)
	}

	return nil
}
