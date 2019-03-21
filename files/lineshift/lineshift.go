package main

import (
	"fmt"

	lineshift "github.com/Brofo/is105-ica03/files/lineshift/lineshiftpack"
)

const textfile = "C:/Users/Sindre/go/src/github.com/Brofo/is105-ica03/files/text1.txt"

//Funksjonen henter inn boolean verdier, og printer ut om de
//er true eller false.
func main() {
	CRLF := lineshift.CheckIfCRLF(textfile)
	CR := lineshift.CheckIfCR(textfile)
	LF := lineshift.CheckIfLF(textfile)

	fmt.Println("CRLF: ", CRLF)
	fmt.Println("CR:   ", CR)
	fmt.Println("LF:   ", LF)
}
