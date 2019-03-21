package filehandler

import (
	"log"
	"os"
)

//Funksjonen har som hensikt å konvertere en tekstfil til bytes,
//og antall bytes basert på lengde på slice.
//@param textfile - Tekstfil av type string.
//@return data - Data ifra tekstfilen i form av en slice av bytes ([]bytes).
//@return count - Antall bytes i data (int).
func MakeFileToByte(textfile string) (data []byte, count int) {
	file, err := os.Open(textfile)
	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	sizeOfSlice := finfo.Size()

	data = make([]byte, sizeOfSlice)
	count, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	return data, count
}
