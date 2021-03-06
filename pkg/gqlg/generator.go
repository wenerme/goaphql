package gqlg

import (
	"bytes"
	"context"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/wenerme/goaphql/pkg/gqll"
)

var _ GenerateFile = (*generateFile)(nil)
var _ Context = (*generateContext)(nil)

type generateFile struct {
	name string
	buf  *bytes.Buffer
}

func (self *generateFile) Name() string {
	return self.name
}
func (self *generateFile) SetName(v string) {
	self.name = v
}
func (self *generateFile) Buf() *bytes.Buffer {
	return self.buf
}

type TemplateFile struct {
	ParseName string
	Tags      []string
	Filename  *template.Template
	Content   *template.Template
}

func (self *TemplateFile) Execute(ctx Context, v interface{}) (err error) {
	file := ctx.File()

	if self.Filename != nil {
		buf := new(bytes.Buffer)
		if err = self.Filename.Execute(buf, v); err != nil {
			return
		}
		file.SetName(buf.String())
	}

	if err = self.Content.Execute(file.Buf(), v); err != nil {
		return
	}

	return nil
}

type Generator interface {
	Context() *generateContext
}
type GenerateFile interface {
	Name() string
	SetName(string)
	Buf() *bytes.Buffer
}

type Context interface {
	context.Context
	File() GenerateFile
	TemplateFile() *TemplateFile
}
type generateContext struct {
	context.Context
	file         GenerateFile
	templateFile *TemplateFile
	// schemaName   string
}

func (self *generateContext) File() GenerateFile {
	return self.file
}

func (self *generateContext) TemplateFile() *TemplateFile {
	return self.templateFile
}

type CommonGenerator struct {
	Template       *template.Template
	Files          []GenerateFile
	TemplateFiles  []*TemplateFile
	Context        Context
	TypeSystem     *gqll.TypeSystem
	contextCreator func(*TemplateFile, interface{}) Context
	conf           *CommonGeneratorConfig
}
type CommonGeneratorConfig struct {
	// Base template
	Template *template.Template
	// Extra functions
	FuncMap map[string]interface{}
}

func (self *CommonGenerator) config(config CommonGeneratorConfig) error {

	if config.Template == nil {
		config.Template = template.New("GraphQLGenerator")
	}

	self.conf = &config
	self.Template = config.Template
	return nil
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

func (self *CommonGenerator) FuncMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		m = make(map[string]interface{})
	}
	m = _TemplateHelper.FuncMap(m)
	m = LangFuncMap(m)
	for k, v := range map[string]interface{}{
		"Context": func() interface{} { return self.Context },
	} {
		m[k] = v
	}
	return m
}

func (self *CommonGenerator) Generate(t *TemplateFile, v interface{}) error {
	ctx := self.createContext(t, v)
	return t.Execute(ctx, v)
}

func (self *CommonGenerator) createContext(t *TemplateFile, v interface{}) Context {
	var ctx Context
	if self.contextCreator != nil {
		ctx = self.contextCreator(t, v)
	} else {
		ctx = &generateContext{
			Context:      context.Background(),
			templateFile: t,
			file:         &generateFile{buf: new(bytes.Buffer)},
		}
	}

	self.Context = ctx
	self.Files = append(self.Files, ctx.File())
	return ctx
}
func (self *CommonGenerator) ScanTemplate() (err error) {
	tpl := self.Template
	var files []*TemplateFile
	for _, t := range tpl.Templates() {
		if !strings.Contains(t.ParseName, "#") {
			continue
		}

		f := &TemplateFile{}
		files = append(files, f)
		f.ParseName = t.ParseName
		f.Content = t
		split := strings.Split(t.ParseName[:len(t.ParseName)-len(".tmpl")], "#")

		switch len(split) {
		case 2:
			// contain tags
			f.Tags = strings.Split(split[0], ",")
			f.Filename, err = tpl.New(f.ParseName + ".fn").Parse(split[1])
			if err != nil {
				return
			}
		default:
			return errors.New("gqlg: unexpected template filename")
		}
	}
	self.TemplateFiles = files
	return nil
}
