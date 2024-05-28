package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Define input and output directories
	inputDir := "input"
	outputDir := "output"

	// Read Markdown files from input directory
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		os.Exit(1)
	}

	// Loop through Markdown files
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".md") {
			// Read Markdown content
			mdContent, err := ioutil.ReadFile(inputDir + "/" + file.Name())
			if err != nil {
				fmt.Println("Error reading Markdown file:", err)
				continue
			}

			// Convert Markdown to HTML (basic conversion)
			htmlContent := []byte("<html><body>" + string(mdContent) + "</body></html>")

			// Generate HTML page
			htmlFileName := strings.TrimSuffix(file.Name(), ".md") + ".html"
			htmlFilePath := outputDir + "/" + htmlFileName

			// Write HTML to file
			err = ioutil.WriteFile(htmlFilePath, htmlContent, 0644)
			if err != nil {
				fmt.Println("Error writing HTML file:", err)
				continue
			}

			fmt.Println("Generated HTML file:", htmlFileName)
		}
	}

	fmt.Println("Website generation completed.")
}
