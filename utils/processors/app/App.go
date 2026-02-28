package app

import (
	//"fmt"
	"fmt"
	"log"
	"os"

	"prac_go-reloaded/utils/processors"
)

func App() {

	// First read file to string and handle errors

	// Checking for Correct Number of Arguments
	if len(os.Args) != 3 {
		log.Fatal("Invalid Number of files, Expected 2 arguments corresponding to input/output file names")
	}

	inputFileName := os.Args[1]
	// outputFileName := os.Args[2]

	// reading inputfile
	buffer, err := os.ReadFile(inputFileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println(buffer)

	// for i := range buffer {
	// 	fmt.Println(string(buffer[i]))
	// }

	text := processors.ReplaceHexAndBin(string(buffer))
	text = processors.FormatCasing(text)
	fmt.Println(text)

	// Then process the text and Apply transformation
	// Attempt to create the Output File and Write to it

}
