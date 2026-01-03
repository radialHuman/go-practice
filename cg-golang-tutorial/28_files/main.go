package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func errorHandelling(err error, lineNumber int) { //#doutb why not use this instead of writting the whole thing again and agains
	if err != nil {
		fmt.Println("!!!! Line number : ", lineNumber)
		panic(err)
	}
} // dont use this, as lines numbers are not static, if something is edited in the first line, all the calls will ave to be edited

func main() {
	fmt.Println(("======================Files============================="))
	fileObject, err := os.Open("file_1.txt") // to read file
	if err != nil {                          //always sue this when ever a function returns an error
		// fmt.Println("Error reading file_1.txt path\n", err)
		// log it
		// panic it
		panic(err)
	}
	defer fileObject.Close() // ones read, this has to be used so that it closes at the end of the function
	// so that even if there is an error it will be closed for sure

	fileInfo, err := fileObject.Stat() // returens file info and error
	if err != nil {                    //always sue this when ever a function returns an error
		// fmt.Println("Error reading file_1.txt's info")
		panic(err)
	}
	fmt.Println("File name :", fileInfo.Name())
	fmt.Println("Is folder? :", fileInfo.IsDir())
	fmt.Println("File size :", fileInfo.Size())
	fmt.Println("Permissions :", fileInfo.Mode())
	fmt.Println("Modified at :", fileInfo.ModTime())

	// reading the content
	fmt.Println(("======================File read============================="))
	// one way to do it, is using buffer memoery and store tha values in it
	// arrray of bytes
	buf := make([]byte, fileInfo.Size()) // making the size dynamic, as per tht file size
	content, err := fileObject.Read(buf)
	errorHandelling(err, 44)
	fmt.Println("Content in the file is : ", content, string(buf))

	// easy but not always useful as it stores the content of the file, no matter the size, in one go which can be difficult to handel
	// for more control and batch read sue the previous method, may be use read at
	fmt.Println(("======================File read easy way============================="))
	fileContent, err := os.ReadFile("file_1.txt")
	errorHandelling(err, 49)
	// defer fileContent.Close() // #doubt not required?
	fmt.Println("File content (the easy way) : ", string(fileContent))

	fmt.Println(("======================Folder read============================="))
	folder, err := os.Open("../")
	errorHandelling(err, 56)
	defer folder.Close()
	fmt.Println("Folder read : ", folder.Name())

	folderInfo, err := folder.Stat()
	errorHandelling(err, 60)
	fmt.Println("Folder Size : ", folderInfo.Size())
	fmt.Println("Folder? : ", folderInfo.IsDir())

	folderRead, err := folder.ReadDir(-1) // all, else give a numebr as per the content in it, which one might not know
	errorHandelling(err, 64)
	for _, i := range folderRead {
		if strings.Contains(i.Name(), ".md") {
			fmt.Println(i.Name())
		}
	}

	fmt.Println(("======================Create file ============================="))
	fileObjectTwo, err := os.Create("file_2.txt")
	errorHandelling(err, 76)    // order matter, this first
	defer fileObjectTwo.Close() // then close
	fmt.Println("Creating file : ", fileObjectTwo)

	fmt.Println(("======================Write into file ============================="))
	newFileContent, err := os.ReadFile("file_2.txt")
	fmt.Println("New file content : ", string(newFileContent))

	fileObjectTwo.WriteString("Something")
	fileObjectTwo.WriteString("Else")               // keeps appending wihtout space or new line
	newFileContent, err = os.ReadFile("file_2.txt") // no := since its updating an existing variable
	fmt.Println("Now file is : ", string(newFileContent))

	fmt.Println(("======================Write into file binary============================="))
	bytes := []byte("\nAnything")
	fileObjectTwo.Write(bytes)
	newFileContent, err = os.ReadFile("file_2.txt")
	fmt.Println("Now file is : ", string(newFileContent))

	// not allin one go but in a streaming manner
	fmt.Println(("======================Copying content using stream============================="))
	// readin them into memory
	sourceFile, err := os.Open("file_2.txt")
	errorHandelling(err, 105)
	defer sourceFile.Close()

	destFile, err := os.Create("file_3.txt")
	errorHandelling(err, 109)
	defer destFile.Close()

	// copying in a streaming manner
	reader := bufio.NewReader(sourceFile) // 4096 is defualt buffer size
	writer := bufio.NewWriter(destFile)
	// one way fo doing this
	for {
		// first read
		bytes, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" { // to prevent end of file error
				panic(err)
			}
			break
		}

		// now write
		err = writer.WriteByte(bytes)
		errorHandelling(err, 122)
	}
	writer.Flush() // to flush anythign else
	fileContent, err = os.ReadFile("file_3.txt")
	errorHandelling(err, 130)
	fmt.Println("After copying content from file 2, file 1 is : ", string(fileContent))
	// #doutb why does this not work on exsting files?, liek doing this on file 1 did not work
	fmt.Println(("======================Delete a file_1.txt============================="))
	// targetFile, err := os.Open("file_1.txt")
	// errorHandelling(err, 133)
	// defer targetFile.Close()

	err = os.Remove("file_1.txt")

	// checking in fodler if its deleted
	errorHandelling(err, 136)
	folderNow, err := os.Open(".")
	errorHandelling(err, 140)
	defer folderNow.Close()
	folderRead, err = folderNow.ReadDir(-1) // all, else give a numebr as per the content in it, which one might not know
	errorHandelling(err, 141)
	for _, i := range folderRead {
		fmt.Println(i.Name())

	}

}

/*
OUTPUT

======================Files=============================
File name : file_1.txt
Is folder? : false
File size : 6
Permissions : -rw-rw-r--
Modified at : 2026-01-03 09:48:32.064167241 +0530 IST
======================File read=============================
Content in the file is :  6 File 1
======================File read easy way=============================
File content (the easy way) :  File 1
======================Folder read=============================
Folder read :  ../
Folder Size :  4096
Folder? :  true
template.md
2.md
======================Create file =============================
Creating file :  &{0xc000066300}
======================Write into file =============================
New file content :
Now file is :  SomethingElse
======================Write into file binary=============================
Now file is :  SomethingElse
Anything
======================Copying content using stream=============================
After copying content from file 2, file 1 is :  SomethingElse
Anything
======================Delete a file=============================
main.go
file_3.txt
file_2.txt

*/
