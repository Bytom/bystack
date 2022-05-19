package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please setup the public directory path")
		return
	}

	inPath := os.Args[1]
	outPath := "dashboard.go"
	var files []string
	if err := filepath.Walk(inPath, visit(&files)); err != nil {
		panic(err)
	}

	out, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	writeString(out, "package dashboard")
	writeString(out, "\n\n")
	writeString(out, "var Files = map[string]string{")
	writeString(out, "\n")

	for _, file := range files {
		bs, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(file, err)
		}

		_, name := filepath.Split(file)
		writeString(out, "\"")
		writeString(out, name)
		writeString(out, "\": ")
		writeString(out, strconv.Quote(string(bs)))
		writeString(out, ",")
		writeString(out, "\n")
		fmt.Println("out:", file)
	}

	writeString(out, "}")
	writeString(out, "\n")
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if info.IsDir() {
			fmt.Printf("Skip directory: %s\n", path)
			return nil
		}

		*files = append(*files, path)
		return nil
	}
}

func writeString(file *os.File, str string) {
	_, err := file.WriteString(str)
	if err != nil {
		panic(err)
	}
}
