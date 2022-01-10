package htmlParser

import (
	"os"
	"text/template"

	"github.com/google/uuid"
)

type htmlParser struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
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
