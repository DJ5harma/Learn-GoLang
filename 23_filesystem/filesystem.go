package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// properTraditional()
	// simpleForSmallFiles()
	// folderReading()
	// createFile()
	streamCopy()
}

func properTraditional() {
	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()

	if err != nil {
		panic(err)
	}

	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.Mode())
	// fmt.Println(fileInfo.ModTime())

	// ------------------------
	// READ

	buf := make([]byte, fileInfo.Size())

	d, err2 := f.Read(buf)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("data", d, string(buf))
}

func simpleForSmallFiles() {
	data, err := os.ReadFile("./example.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func folderReading() {
	dir, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	fileInfo, err := dir.ReadDir(0)

	for _, fi := range fileInfo {
		fmt.Println(fi.Name(), "isDir:", fi.IsDir())
	}
}

func createFile() {
	file, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// file.WriteString("My name is anthony ghonsalvis\n")
	// file.WriteString("Mai duniya mei akela hun")
	// file.Write([]byte("some here bytes"))
	file.WriteAt([]byte("OGGY"), fileInfo.Size())
}

func streamCopy() {
	srcFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if err := writer.WriteByte(b); err != nil {
			panic(err)
		}
	}

	writer.Flush()
	fmt.Println("Written to new file")

	destFile.Close()
	err = os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("deleted file")
}
