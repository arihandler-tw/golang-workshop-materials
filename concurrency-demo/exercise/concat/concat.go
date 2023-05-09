package concat

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Merge(dir string, output string, workers int8) error {
	concat, err := Concat(dir, workers)
	if err != nil {
		return err
	}

	return writeToFile(concat, output)
}

func Concat(dir string, workers int8) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error reading directory %v: %v", dir, err)
		return nil, err
	}

	numFiles := len(files)
	ch := make(chan []string, workers)

	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if path == dir || info.IsDir() {
			return nil
		}

		go readFileToUppercase(path, ch)

		return nil
	})

	if err != nil {
		log.Fatalf("Error traversing directory %v: %v", dir, err)
		return nil, err
	}

	var result []string
	for i := 0; i < numFiles; i++ {
		result = append(result, <-ch...)
	}

	return result, nil
}

func readFileToUppercase(file string, ch chan []string) {
	readFile, err := os.Open(file)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fileLines = append(fileLines, strings.ToUpper(line))
	}

	ch <- fileLines
}

func writeToFile(concat []string, output string) error {
	outputFile, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	for _, l := range concat {
		if _, err := fmt.Fprintln(outputFile, l); err != nil {
			return err
		}
	}
	return nil
}
