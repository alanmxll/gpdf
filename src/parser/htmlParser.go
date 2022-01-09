package parser

import (
	"html/template"
	"os"

	parser "github.com/alanmxll/gpdf/src/parser/interfaces"
	"github.com/google/uuid"
)

type htmlParser struct {
	rootPath string
}

func New(rootPath string) parser.HTMLParserInterface {
	return &htmlParser{rootPath: rootPath}
}

func (p *htmlParser) Create(templateName string, data interface{}) (string, error) {
	templateGenerator, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}

	fileName := p.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err = templateGenerator.Execute(fileWriter, data); err != nil {
		return "", err
	}

	return fileName, nil
}
