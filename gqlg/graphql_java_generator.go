package gqlg

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/wenerme/goaphql/gqll"
	"strings"
)

type GraphQLJavaGenerator struct {
	CommonGenerator
	Config *JavaGeneratorConfig
}

type JavaGeneratorConfig struct {
	CommonGeneratorConfig
	TypeSystem     *gqll.TypeSystem
	SchemaName     string
	JavaPackage    string
	ImportPackages []string
	JavaTypeMap    map[string]string
}

func NewGraphQLJavaGenerator(config JavaGeneratorConfig) (*GraphQLJavaGenerator, error) {
	conf := &config
	if config.TypeSystem == nil {
		return nil, errors.New("gqlg: generator need TypeSystem")
	}

	g := &GraphQLJavaGenerator{
		Config: conf,
	}
	err := g.config(conf.CommonGeneratorConfig)
	if err != nil {
		return nil, err
	}

	g.TypeSystem = config.TypeSystem
	g.Template.Funcs(g.FuncMap(nil))
	g.ScanTemplate()
	return g, nil
}

func (self *GraphQLJavaGenerator) FuncMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		m = make(map[string]interface{})
	}
	m = self.CommonGenerator.FuncMap(m)
	for k, v := range map[string]interface{}{
		"Config":          func() interface{} { return self.Config },
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
	return ""
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
	return ""
}

// Map a name for Java
func (self *GraphQLJavaGenerator) GenJavaTypeName(name string) string {
	if name, ok := self.Config.JavaTypeMap[name]; ok {
		return name
	}
	switch name {
	case "ID":
		return "String"
	case "String":
		return "String"
	case "Int":
		return "Integer"
	case "Boolean":
		return "Boolean"
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

func (self *GraphQLJavaGenerator) GenerateGraphQLJava() (err error) {
	// Generate type
	for _, v := range self.TypeSystem.GetDefinitions() {
		typeName := gqll.TypeOf(v).Name()
		if !strings.HasSuffix(typeName, "TypeDefinition") {
			continue
		}
		tags := []string{typeName[:len(typeName)-len("TypeDefinition")]}
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
