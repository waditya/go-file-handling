package filewriter

import (
	"fmt"
	"os"
)

func WriteStringToFile(filepath string, data string) {
	filePointer, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println("\n Error encountered while opening the file!\n", err)
		return
	}

	defer filePointer.Close()

	numberOfBytesWritten, err := filePointer.WriteString(data)

	if err != nil {
		fmt.Println("\n Encountered error while writing to file!\n", err)
		return
	}
	fmt.Println(numberOfBytesWritten, " bytes written to the file")

}
