package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/waditya/go-file-handling/filewriter"
)

func main() {
	path := filepath.Join("dir1", "dir2//", "dir3/../dir4", "filename.txt")
	fmt.Println("\n [1] Get the normalized path")
	fmt.Println("\n Using join method in filepath package constructs the filepath as :", path)
	fmt.Println("\n*** DO NOT USE concatenation. ALWAYS USE Failepath.Join ***")
	fmt.Println("\n [2] Get the path directory")
	fmt.Println("\n*** Get the paths directory (All except filename) : ", filepath.Dir(path), " ***")
	fmt.Println("\n [3] Get the last element in the path")
	fmt.Println("\n ", filepath.Base(path))
	fmt.Println("\n [4] Check if a path is an absolute path")
	fmt.Println("\n Check for ", path, " : ", filepath.IsAbs(path))
	fmt.Println("\n\n Check for  e:\\dir1\\dir2\\dir4: ", filepath.IsAbs("e:\\dir1\\dir2\\dir4"))
	fmt.Println("\n [5] Return the extenstion of the filename in a file path")
	fmt.Println("\n*** File extension is : ***", filepath.Ext(path))

	fileMetadata, err := os.Stat("E:\\Kode Kloud\\Golang\\advanced-go-code\\go-file-handling\\main.go")

	if err != nil {
		fmt.Println("Error while getting file metadata", err)
	}

	fmt.Println("\nFile Name is :\n", fileMetadata.Name())
	fmt.Println("\nFile Size is :\n", fileMetadata.Size())
	fmt.Println("\nFile Mode is :\n", fileMetadata.Mode())
	fmt.Println("\nIs it directory :\n", fileMetadata.IsDir())

	file2Metadata, err := os.Stat(path)

	if err != nil {
		fmt.Println("Error while getting file metadata", err)
	}

	fmt.Println("\nFile2 Metadata is :\n", file2Metadata)

	pathTempFile := "E:\\Kode Kloud\\Golang\\advanced-go-code\\go-file-handling\\temp.json"
	data, err := os.ReadFile(pathTempFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n Get data from file")
	fmt.Println(string(data))

	// Using buffers to read large file

	file, err := os.Open(pathTempFile)

	if err != nil {
		fmt.Println("\nError opening the file : ", err)
	}

	bufferToReadFile := make([]byte, 2)

	for {
		numberOfBytesRead, err := file.Read(bufferToReadFile)

		if err != nil {
			fmt.Println("Error reading the file in buffer : \t", err)
			break
		}
		fmt.Println(string(bufferToReadFile[0:numberOfBytesRead]))
	}

	// Writing to large files

	filewriter.WriteStringToFile("E:\\Kode Kloud\\Golang\\advanced-go-code\\go-file-handling\\temp.txt", "This is the title of the file")
}
