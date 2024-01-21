package main

import (
	"bufio"
	"fmt"
	"ijafa/lib"
	"log"
	"os"
	"path/filepath"
)

func CheckFileExtension(filename string) bool {
	if filepath.Ext(filename) == ".ijafa" {
		return true
	}
	return false
}

func ConvertFile(suuu string) {

	if !CheckFileExtension(suuu) {
		fmt.Println("File extension not supported")
		return
	}

	file, err := os.Open(suuu)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res = ""

	for scanner.Scan() {

		text := scanner.Text()
		translation := lib.GetTranslation(text)
		res += translation + "\n"
		// fmt.Println(GetTranslation(scanner.Text()))
	}
	os.WriteFile("test/i.ijafa", []byte(res), 0644)
}

func main() {
	fmt.Println("Welcome to ijafa")
	ConvertFile("test/l.ijafa")
}
