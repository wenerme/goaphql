package gqlg

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/wenerme/goaphql/gqll"
	"strings"
)

type GraphQLJavaGenerator struct {
	CommonGenerator
	Config     *JavaGeneratorConfig
	TypeSystem *gqll.TypeSystem
}

type JavaGeneratorConfig struct {
	JavaPackage    string
	ImportPackages []string
	JavaTypeMap    map[string]string
}

func (self *GraphQLJavaGenerator) FuncMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		m = make(map[string]interface{})
	}
	for k, v := range map[string]interface{}{
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
		"Config":       func() interface{} { return self.Config },
		"ToUpperCamel": func(s string) string { return strings.Title(s) },
		"Join":         func(a []string, seq string) string { return strings.Join(a, seq) },
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
func (self *GraphQLJavaGenerator) GenJavaTypeName(name string) string {
	switch name {
	case "ID":
		return "String"
	case "String":
		return "String"
	case "Int":
		return "Integer"
	case "Boolean":
		return "Boolean"
	case "DateTime":
		return "Date"
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
