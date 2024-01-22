package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/**
 *	GetTranslation
 *	Get the translation of a word
 *	@param {string} word
 *	@return {string} translation
 */
func CheckFileExtension(filename string) bool {
	return filepath.Ext(filename) == ".ijafa"
}

/**
 *	ConvertFile
 *	Convert a file from ijafa to java
 *	@param {string} suuu (filename)
 */
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

	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	lengthFile := float64(fileinfo.Size())
	scanner := bufio.NewScanner(file)

	var (
		res   string
		bytes float64
	)

	for scanner.Scan() {

		text := scanner.Text()
		bytes += float64(len(text + "\n"))
		translation := GetTranslation(text)
		res += translation + "\n"
		translationTime := bytes / lengthFile * 100
		fmt.Printf("\r%.2f%%\n", translationTime)
		// time.Sleep(1000 * time.Millisecond) // debug pour test la barre de progression
		// fmt.Println(GetTranslation(scanner.Text()))
	}
	os.WriteFile(strings.Split(suuu, ".")[0]+".java", []byte(res), 0644)
}
