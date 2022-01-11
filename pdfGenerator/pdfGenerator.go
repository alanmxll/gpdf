package pdfGenerator

import (
	"fmt"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wkhtml struct {
	rootPath string
}

func NewWkHtmlToPDF(rootPath string) PDFGeneratorInterface {
	return &wkhtml{rootPath: rootPath}
}

func (w *wkhtml) Create(htmlFile string) (string, error) {
	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		fmt.Println("CAAA")
		return "", err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
