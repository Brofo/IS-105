package main

import (
	"fmt"

	"github.com/Brofo/ica02/fileutils"
)

const text1 = "C:/Users/Sindre/go/src/github.com/Brofo/is105-ica03/files/text1.txt"
const text2 = "C:/Users/Sindre/go/src/github.com/Brofo/is105-ica03/files/text2.txt"

func main() {
	a := fileutils.FileToByteslice(text1)
	b := fileutils.FileToByteslice(text2)
	fmt.Printf("%q", a)
	fmt.Println(" ")
	fmt.Printf("%q", b)
}
