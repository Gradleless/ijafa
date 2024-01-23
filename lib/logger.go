package lib

import (
	"fmt"
)

const (
	errorColor = "\033[31m"
	resetColor = "\033[0m"
	infoColor  = "\033[33m"
	doneColor  = "\033[32m"
)

func PrintError(s string) {
	fmt.Println(errorColor + "[Error] " + s + resetColor)
}

func PrintInfo(s string) {
	fmt.Println(infoColor + "[Info] " + s + resetColor)
}

func PrintDone(s string) {
	fmt.Println(doneColor + s + resetColor)
}
