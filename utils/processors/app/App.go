package app

import (
	//"fmt"
	"bufio"
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
	outputFileName := os.Args[2]

	// reading inputfile
	buffer, err := os.ReadFile(inputFileName)
	// Handling errors in readingfile
	if err != nil {
		log.Fatal(err.Error())
	}

	// Then process the text and Apply transformation
	text:= processors.ReplaceHexAndBin(string(buffer))
	text = processors.FormatCasing(text)
	text = processors.AdjustPunctuations(text)
	text = processors.AdjustQuotations(text)
	finalProcessedText := processors.ModifyArticle(text)
	fmt.Println(finalProcessedText)

	
	// Attempt to create the Output File and Write to it
	outputFile, err := os.Create(outputFileName)

	if err != nil {
		log.Fatalf("Error creating outputfile: %v \n", err)
	}

	// shceduling the output file to close when the function exits
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	_ , err = writer.WriteString(finalProcessedText)

	if err != nil {
		log.Fatalf("Error Writing to output file: %v\n", err)
	}
	writer.Flush()






	
}
