package lineshift

import (
	filehandler "github.com/Brofo/is105-ica03/files/lineshift/lineshiftpack/filehandler"
)

//Verdi som skal brukes til å sjekke om bytes er CR LF.
var byteValue byte

//boolean verdier som returner true hvis teksten bruker
//CRLF, CR eller LF, og false hvis teksten ikke bruker det.
var CRLF bool = false
var CR bool = false
var LF bool = false

//Alle de tre funksjonene har som hensikt å finne ut om tekstfilen
//inneholder den metoden for linjeskifte som står i funksjonsnavnet.
//@param textfile - Tekstfilen som skal undersøkes.
//@return bool - returnerer true eller false avhengig av resultatet.

func CheckIfCRLF(textfile string) bool {
	data, count := filehandler.MakeFileToByte(textfile)

	for i := 0; i < len(data[:count]); i++ {
		byteValue = data[i]
		if byteValue == 13 {
			i++
			byteValue = data[i]
			if byteValue == 10 {
				CRLF = true
			}
		}
	}
	return CRLF
}

func CheckIfCR(textfile string) bool {
	data, count := filehandler.MakeFileToByte(textfile)

	for i := 0; i < len(data[:count]); i++ {
		byteValue = data[i]
		if byteValue == 13 {
			i++
			byteValue = data[i]
			if byteValue != 10 {
				CR = true
			}
		}
	}
	return CR
}

func CheckIfLF(textfile string) bool {
	data, count := filehandler.MakeFileToByte(textfile)

	for i := 0; i < len(data[:count]); i++ {
		byteValue = data[i]
		if byteValue == 10 {
			i--
			byteValue = data[i]
			if byteValue != 13 {
				LF = true
			}
			//For at ikke loopen skal gå uendelig, setter jeg i til å nå
			//grenseverdien så loopen slutter, dersom LF = true.
			i = len(data[:count]) - 1
		}
	}
	return LF
}
