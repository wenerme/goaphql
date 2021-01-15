package gqlg

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wenerme/goaphql/pkg/gqll"
)

// var defaultJavaTypeMap = map[string]string{
// 	"ID":  "String",
// 	"Int": "Int",
// }

type GraphQLJavaGenerator struct {
	CommonGenerator
	Config *JavaGeneratorConfig

	javaTypes map[string]*javaTypeDefinition
	ignores   map[string]struct{}
}

type GenerateJavaFile struct {
	generateFile
	JavaPackage string
}

type JavaGeneratorConfig struct {
	CommonGeneratorConfig
	TypeSystem     *gqll.TypeSystem
	SchemaName     string
	JavaPackage    string
	ImportPackages []string
	JavaTypeMap    map[string]string
	Ignores        []string
}

func NewGraphQLJavaGenerator(config JavaGeneratorConfig) (*GraphQLJavaGenerator, error) {
	conf := &config
	if config.TypeSystem == nil {
		return nil, errors.New("gqlg: generator need TypeSystem")
	}

	g := &GraphQLJavaGenerator{
		Config:    conf,
		javaTypes: make(map[string]*javaTypeDefinition),
		ignores:   make(map[string]struct{}),
	}
	err := g.config(conf.CommonGeneratorConfig)
	if err != nil {
		return nil, err
	}
	for _, v := range conf.Ignores {
		g.ignores[v] = struct{}{}
	}

	g.contextCreator = g.createContext

	g.TypeSystem = config.TypeSystem
	g.Template.Funcs(g.FuncMap(nil))
	return g, g.ScanTemplate()
}

func (self *GraphQLJavaGenerator) createContext(t *TemplateFile, v interface{}) Context {
	var ctx Context
	file := &GenerateJavaFile{
		generateFile: generateFile{buf: new(bytes.Buffer)},
		JavaPackage:  self.Config.JavaPackage,
	}
	if typ, ok := self.javaTypes[gqll.NameOf(v)]; ok {
		if typ.javaPackage != "" {
			file.JavaPackage = typ.javaPackage
		}
	}
	ctx = &generateContext{
		Context:      context.Background(),
		templateFile: t,
		file:         file,
	}

	self.Context = ctx
	self.Files = append(self.Files, ctx.File())
	return ctx
}
func (self *GraphQLJavaGenerator) FuncMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		m = make(map[string]interface{})
	}
	m = self.CommonGenerator.FuncMap(m)
	for k, v := range map[string]interface{}{
		"Config":          func() interface{} { return self.Config },
		"IsSkipped":       self.IsSkipped,
		"GenJavaType":     self.GenJavaType,
		"GenJavaTypeName": self.GenJavaTypeName,
		"GenJavaTypeNames": func(a []string) []string {
			var s []string
			for _, v := range a {
				s = append(s, self.GenJavaTypeName(v))
			}
			return s
		},
		"GenJavaValue": self.GenJavaValue,
	} {
		m[k] = v
	}
	return m
}

func LangFuncMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		m = make(map[string]interface{})
	}
	for k, v := range map[string]interface{}{
		"TypeOf": gqll.TypeOf,
		"TypeNameOf": func(v interface{}) string {
			t := gqll.TypeOf(v)
			if t != nil {
				return t.Name()
			}
			return ""
		},
	} {
		m[k] = v
	}
	return m
}

type javaTypeDefinition struct {
	name        string
	javaType    string
	javaPackage string
}

func (self *GraphQLJavaGenerator) ScanDefinition() (err error) {
	defMap := make(map[string]*javaTypeDefinition)
	for _, v := range self.TypeSystem.GetDefinitions() {
		t := gqll.TypeOf(v)
		if !t.IsTypeDefinition() {
			continue
		}
		name := gqll.NameOf(v)
		javaDef := &javaTypeDefinition{
			name: name,
		}
		defMap[name] = javaDef
		if err = scanJavaType(javaDef, v); err != nil {
			return
		}
	}
	self.javaTypes = defMap
	return nil
}

// directive @JavaType(package: String, type: String)
func scanJavaType(t *javaTypeDefinition, def gqll.Definition) (err error) {
	if v, ok := def.(gqll.HasDirectives); ok {
		d := v.GetDirective("JavaType")
		if d == nil {
			return
		}
		argument := d.GetArgument("package")
		if argument != nil {
			if val, ok := argument.Value.(*gqll.StringValue); ok {
				t.javaPackage = val.Value
			} else {
				logrus.WithField("Def", d.SourceLocation).WithField("Arg", "package").Fatal("incorrect argument type")
			}
		}

		argument = d.GetArgument("type")
		if argument != nil {
			if val, ok := argument.Value.(*gqll.StringValue); ok {
				t.javaType = val.Value
			} else {
				logrus.WithField("Def", d.SourceLocation).WithField("Arg", "type").Fatal("incorrect argument type")
			}
		}
	}
	return
}

func (self *GraphQLJavaGenerator) GenJavaValue(v interface{}) string {
	switch v := v.(type) {
	case *gqll.IntValue:
		return fmt.Sprint(v.Value)
	case *gqll.FloatValue:
		return fmt.Sprint(v.Value)
	case *gqll.BooleanValue:
		return fmt.Sprint(v.Value)
	case *gqll.NullValue:
		return "null"
	case *gqll.StringValue:
		return fmt.Sprint('"', v.Value, '"')
	default:
		panic("Unexpected")
	}
}
func (self *GraphQLJavaGenerator) GenJavaType(v interface{}) string {
	switch v := v.(type) {
	case *gqll.NonNullType:
		return "@NotNull " + self.GenJavaType(v.Type)
	case *gqll.NamedType:
		return self.GenJavaTypeName(v.Name)
	case *gqll.ListType:
		return "List<" + self.GenJavaType(v.Type) + ">"
	case string:
		return self.GenJavaTypeName(v)
	default:
		panic("Unexpected")
	}
}

// Map a name for Java
func (self *GraphQLJavaGenerator) GenJavaTypeName(name string) string {
	if name, ok := self.Config.JavaTypeMap[name]; ok {
		return name
	}
	if typ, ok := self.javaTypes[name]; ok && typ.javaType != "" {
		return typ.javaType
	}
	def := self.TypeSystem.GetNamedDefinition(name)
	if def == nil {
		logrus.Fatal("javagen: ", name, " type not found")
	}
	switch gqll.TypeOf(def).Name() {
	case "InterfaceTypeDefinition":
		fallthrough
	case "ObjectTypeDefinition":
		return strings.Title(name) + "Resolver"
	}
	return name
}

// Is this definition should be skipped when generate
func (self *GraphQLJavaGenerator) IsSkipped(d gqll.Definition) bool {
	if _, ok := self.ignores[gqll.NameOf(d)]; ok {
		return true
	}
	return false
}
func (self *GraphQLJavaGenerator) generateTags(d gqll.Definition) (tags []string) {
	t := gqll.TypeOf(d)
	if _, ok := self.ignores[gqll.NameOf(d)]; ok {
		return nil
	}
	if t.IsTypeDefinition() {
		typeName := t.Name()
		tags = []string{typeName[:len(typeName)-len("TypeDefinition")]}
	}

	return
}

func (self *GraphQLJavaGenerator) prepareGenerateGraphQLJava() (err error) {
	if err := self.ScanTemplate(); err != nil {
		return errors.Wrap(err, "failed to scan template file")
	}
	if err := self.ScanDefinition(); err != nil {
		return errors.Wrap(err, "failed to scan definiyion")
	}
	return
}
func (self *GraphQLJavaGenerator) GenerateGraphQLJava() (err error) {
	if err = self.prepareGenerateGraphQLJava(); err != nil {
		return
	}
	// Generate type
	for _, v := range self.TypeSystem.GetDefinitions() {

		tags := self.generateTags(v)
		typeName := gqll.TypeOf(v).Name()
		files := self.SelectTemplateFileByTags(tags...)

		if len(files) == 0 {
			continue
		}

		logrus.WithFields(logrus.Fields{
			"Type":  typeName,
			"Name":  gqll.NameOf(v),
			"Count": len(files),
		}).Info("Generate type definition")

		for _, f := range files {
			if err = self.Generate(f, v); err != nil {
				return
			}
		}
	}

	// Generate common
	files := self.SelectTemplateFileByTags("Java")
	for _, f := range files {
		if err = self.Generate(f, self.TypeSystem); err != nil {
			return
		}
	}

	return nil
}
