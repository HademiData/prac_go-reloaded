package app

import (
	"fmt"
	"log"
	"os"
)

func App() {
	fmt.Println("Hello, World!")

	// First read file to string and handle errors

	// Checking for Correct Number of Arguments
	if len(os.Args) != 3 {
		log.Fatal("Invalid Number of Arguments, Expected 2 arguments corresponding to input/output file names")
	}

	inputFileName := os.Args[1]
	// outputFileName := os.Args[2]

	// reading inputfile
	buffer, err := os.ReadFile(inputFileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(buffer))

	// Then process the text and Apply transformation
	// Attempt to create the Output File and Write to it

}
