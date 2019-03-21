package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	argument := flag.Arg(0)
	filename := flag.String("f", argument, "info about the file")
	flag.Parse()

	fi, err := os.Lstat(*filename)
	if err != nil {
		log.Fatal(err)
	}

	sizeBytes := float64(fi.Size())
	kibiBytes := float64(sizeBytes / 1024)
	mibiBytes := float64(kibiBytes / 1024)
	gibiBytes := float64(mibiBytes / 1024)

	fmt.Println("\nFilename: ", fi.Name())
	fmt.Println("Path: ", *filename)
	fmt.Printf("Size in... \nBytes: %f. \nKibibytes: %f. \nMibibytes: %f. \nGibibytes: %f\n",
		sizeBytes, kibiBytes, mibiBytes, gibiBytes)
	fmt.Printf("Unix permission bits: %#o\n", fi.Mode().Perm())
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("Is a regular file")
	case mode.IsDir():
		fmt.Println("Is a directory")
	case mode&os.ModeAppend != 0:
		fmt.Println("Is append only")
	case mode&os.ModeDevice != 0:
		fmt.Println("Is a device file")
	case mode&os.ModeCharDevice != 0:
		fmt.Println("Is a unix character device")
	case mode&os.ModeSocket != 0:
		fmt.Println("Is a unix block device")
	case mode&os.ModeSymlink != 0:
		fmt.Println("Is a symbolic link")
	}
}
