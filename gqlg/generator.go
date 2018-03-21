package gqlg

import (
	"bytes"
	"text/template"
)

type TemplateFile struct {
	ParseName string
	Tags      []string
	Filename  *template.Template
	Content   *template.Template
	Generator Generator
}

func (self *TemplateFile) Apply(v interface{}) (err error) {
	file := self.Generator.GetGenerateContext().File

	if self.Filename != nil {
		buf := new(bytes.Buffer)
		if err = self.Filename.Execute(buf, v); err != nil {
			return
		}
		file.Name = buf.String()
	}

	if err = self.Content.Execute(file.Buf, v); err != nil {
		return
	}

	return nil
}

type Generator interface {
	//GetGenerateFile(name string) *GenerateFile
	GetGenerateContext() *GenerateContext
}
type GenerateContext struct {
	File         *GenerateFile
	TemplateFile *TemplateFile
}

var _ Generator = (*CommonGenerator)(nil)

type CommonGenerator struct {
	Template        *template.Template
	Files           []*GenerateFile
	TemplateFiles   []*TemplateFile
	GenerateContext *GenerateContext
}

func (self *CommonGenerator) GetGenerateContext() *GenerateContext {
	return self.GenerateContext
}

type GenerateFile struct {
	Name string
	Buf  *bytes.Buffer
}

func (self *CommonGenerator) SelectTemplateFileByTags(tags ...string) []*TemplateFile {
	var files []*TemplateFile

	for _, f := range self.TemplateFiles {
		for _, t := range f.Tags {
			for _, tag := range tags {
				if t == tag {
					files = append(files, f)
					goto NEXT_FILE
				}
			}
		}
	NEXT_FILE:
	}

	return files
}
