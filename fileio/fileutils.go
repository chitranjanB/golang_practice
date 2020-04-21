package fileio

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) string {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	return string(dat)
}

func ReadFileLines(filePath string) []string {
	readFile, err := os.Open(filePath)
	check(err)
	defer func() {
		if readFile != nil {
			readFile.Close()
		}
	}()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	return fileTextLines
}

func WriteToFile(filePath string, content string) error {
	data := []byte(content)
	return ioutil.WriteFile(filePath, data, 0644)
}

func AppendToFile(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + content); err != nil {
		return err
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
