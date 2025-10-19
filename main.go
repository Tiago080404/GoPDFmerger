package main

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func main() {
	inputFiles := []string{"/home/tiago/Documents/Berichtsheft/2025_06_30_TD.pdf", "/home/tiago/Documents/Berichtsheft/2025_07_14_TD.pdf"}

	outputFile := "/home/tiago/Documents/Berichtsheft/mergedtry.pdf"
	config := model.NewDefaultConfiguration()

	err := api.MergeCreateFile(inputFiles, outputFile, false, config)
	if err != nil {
		fmt.Printf("Error merging PDFs: %v\n", err)
		return
	}

	fmt.Println("PDFs merged successfully into this file:", outputFile)
}
