package main

import (
	"fmt"
	"os"
)

func removePath(path string) {
	_, error := os.Stat(path)
	if error == nil {
		os.RemoveAll(path)
	}
}

func main() {

	// read file to bytestream
	bs, err := os.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	// if output.txt file exists remove it
	removePath("output.txt")

	// write to file
	file, err := os.Create("output.txt")
	if err != nil {
		return
	}
	defer file.Close() // close file at end
	file.WriteString(str + "\nand a newly added string")

	// remove the tmp1 directory if it exists
	removePath("tmp1")

	// make a new directory, add a folder and a file in the folder
	directory_error := os.MkdirAll("tmp1/folder1", 0750) // the 0750 is the read/write permission bits in linux.
	if directory_error != nil {
		fmt.Println("directory error: ", directory_error)
		fmt.Println("cannot create new directory!")
	}
	os.WriteFile("tmp1/folder1/temp_file.txt", []byte("a temp file has been created"), 0660)

	// list the directory contents
	folder_list, err := os.ReadDir("tmp1")
	fmt.Println("contents of the tmp1 directory: ", folder_list)
	folder_list, err = os.ReadDir("tmp1/folder1")
	fmt.Println("contents of tmp1/folder1: ", folder_list)

	// move the temp_file.txt out of folder1 AND rename the file
	os.Rename("tmp1/folder1/temp_file.txt", "tmp1/temp_file_renamed_dos.txt")
	fmt.Println(os.ReadDir("tmp1"))

}
