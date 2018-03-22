package gqlg

import "strings"

var _TemplateHelper = &TemplateHelper{}

// Provide some helper function for template use
type TemplateHelper struct {
}

func (self *TemplateHelper) FuncMap(m map[string]interface{}) map[string]interface{} {
	if m == nil {
		m = make(map[string]interface{})
	}
	for k, v := range map[string]interface{}{
		"ToUpperCamel": func(s string) string { return strings.Title(s) },
		"Join":         func(a []string, seq string) string { return strings.Join(a, seq) },
	} {
		m[k] = v
	}
	return m
}
