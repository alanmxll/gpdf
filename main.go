package main

import (
	"fmt"
	"os"

	"github.com/alanmxll/gpdf/htmlParser"
	"github.com/alanmxll/gpdf/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {

	h := htmlParser.New("tmp")
	wkhtml := pdfGenerator.NewWkHtmlToPDF("tmp")

	dataHTML := Data{
		Name: "Alan",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer os.Remove(htmlGenerated)
	fmt.Println("HTML generated", htmlGenerated)

	filePDFName, err := wkhtml.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("PDF generated", filePDFName)
}
