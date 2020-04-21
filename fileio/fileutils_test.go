package fileio

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	filePath := "sample.txt"
	content := ReadFile(filePath)
	fmt.Println(content)
}

func TestReadFileLines(t *testing.T) {
	filePath := "sample.txt"
	lines := ReadFileLines(filePath)
	for _, line := range lines {
		fmt.Println("line -> ", line)
	}
}

func TestWriteToFile(t *testing.T) {
	filePath := "sample.txt"
	err := WriteToFile(filePath, "new line")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func TestAppendToFile(t *testing.T) {
	filePath := "sample.txt"
	err := AppendToFile(filePath, "another new line")
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
