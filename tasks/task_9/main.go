package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	dirFlag string
)

func init() {
	flag.StringVar(&dirFlag, "P", "", "points to the directory for the downloaded website")
}

func main() {
	flag.Parse()
	url := os.Args[len(os.Args)-1]

	fmt.Println(dirFlag)

	file, err := createFile()
	if err != nil {
		fmt.Println("failed to create a file:", err)
		return
	}

	content, err := readContent(url)
	if err != nil {
		fmt.Println("failed to read content:", err)
		return
	}

	_, err = io.WriteString(file, string(content))
	if err != nil {
		fmt.Println("failed to write content to the file:", err)
		return
	}
}

func createFile() (*os.File, error) {
	var file *os.File
	var err error

	if dirFlag != "" {
		file, err = os.Create(dirFlag + "/" + dirFlag + ".html")
		if err != nil {
			fmt.Println("failed to create a file:", err)
			return nil, err
		}
	} else {
		file, err = os.Create("output.html")
		if err != nil {
			fmt.Println("failed to create a file:", err)
			return nil, err
		}
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("failed to close the file")
			return
		}
	}(file)

	return file, nil
}

func readContent(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to load the page:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println("failed to close response body")
			return
		}
	}(response.Body)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("failed to read the content:", err)
		return nil, err
	}

	return content, nil
}
