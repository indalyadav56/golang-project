package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {

	// move
	// cwd, _ := os.Getwd()
	// fmt.Println(cwd)

	// indalPath := filepath.Join(cwd, "indal")

	// err := os.Rename("test.txt", filepath.Join(indalPath, "indal.txt"))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("done...")

	// copy file

	// // src
	// file, err := os.Open("indal.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // destination
	// wObj, err := os.Create("indal/copy.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// byteDataLen, err := io.Copy(wObj, file)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(byteDataLen)

	// file, err := os.Open("large-file.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer file.Close()

	// buff := bufio.NewReader(file)
	// buffer := make([]byte, 10)

	// for {
	// 	n, err := buff.Read(buffer)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }

	fileObj, err := os.Create("test.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer fileObj.Close()

	zipWritterObj := zip.NewWriter(fileObj)
	defer zipWritterObj.Close()

	// which file you want to copy in zip
	fileSrc, _ := os.Open("large-file.txt")
	defer fileSrc.Close()

	writer, err := zipWritterObj.Create("large-file.txt")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(writer, fileSrc)
	if err != nil {
		panic(err)
	}

	// file _. syscall - filedecriptor kernal
	// for {
	// 	_, err := file.Read(buffer)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		fmt.Println("Read error:", err)
	// 		return
	// 	}
	// fmt.Print(string(buffer))
	// }

	// read zip file

	// frObj, _ := os.Open("test.zip")
	// defer frObj.Close()

	// zipReader, _ := zip.OpenReader("test.zip")

	// defer zipReader.Close()

	// for _, f := range zipReader.File {
	// 	fmt.Println(f)
	// }

	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		fmt.Println("Error opening zip:", err)
		return
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fmt.Println("File inside zip:", f.Name)
	}

}

func CreateZipFile() {
	fileObj, err := os.Create("test.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer fileObj.Close()

	zipWriter := zip.NewWriter(fileObj)
	defer zipWriter.Close()

	// Open source file to zip
	fileSrc, err := os.Open("large-file.txt")
	if err != nil {
		fmt.Println("Source file error:", err)
		return
	}
	defer fileSrc.Close()

	// Create file entry inside zip
	writer, err := zipWriter.Create("large-file.txt")
	if err != nil {
		fmt.Println("Zip entry error:", err)
		return
	}

	_, err = io.Copy(writer, fileSrc)
	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}

	fmt.Println("File successfully zipped.")
}

// // Create directory
// err = os.MkdirAll(newDir, 0755)
// if err != nil {
// 	fmt.Println("Error creating directory:", err)
// 	return
// }
// fmt.Println("Directory created at:", newDir)
