package gqlp

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/wenerme/goaphql/pkg/gqll"
	"io/ioutil"
	"testing"
)

func TestSpecParse(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/spec.graphql")
	assert.NoError(t, err)
	doc, err := ParseContent(string(b))
	assert.NoError(t, err)
	_ = doc
}
func TestDefinitionSpec(t *testing.T) {
	for _, test := range []struct {
		raw string
		def interface{}
	}{
		{
			raw: `"Test Directive Def" directive @dA repeatable on VARIABLE_DEFINITION`,
			def: func() interface{} {
				v := new(gqll.DirectiveDefinition)
				v.Description = "Test Directive Def"
				v.Name = "dA"
				v.Repeatable = true
				v.Locations = []string{"VARIABLE_DEFINITION"}
				return v
			}(),
		},
	} {
		doc, err := ParseContent(test.raw)
		if assert.NoError(t, err) {
			doc.Definitions[0].SetSourceLocation(nil)
			if !assert.Equal(t, test.def, doc.Definitions[0]) {
				spew.Dump(doc.Definitions[0])
			}
		}

	}
}
