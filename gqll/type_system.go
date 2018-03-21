package gqll

import (
	"github.com/pkg/errors"
	"strings"
)

type TypeSystem struct {
	definitions []Definition
	nameMap     map[string]Definition
	typeMap     map[string][]Definition
}

func (self *TypeSystem) GetNamedDefinition(name string) Definition {
	return self.nameMap[name]
}
func (self *TypeSystem) GetTypedDefinitions(typ string) []Definition {
	return self.typeMap[typ]
}
func (self *TypeSystem) GetDefinitions() []Definition {
	return self.definitions
}

func NewTypeSystem(doc *Document) (*TypeSystem, error) {
	ts := &TypeSystem{
		nameMap: make(map[string]Definition),
		typeMap: make(map[string][]Definition),
	}
	namesMap := make(map[string][]Definition)

	for _, v := range doc.Definitions {
		// name index
		if vn, ok := v.(interface {
			GetName() string
		}); ok {
			name := vn.GetName()
			namesMap[name] = append(namesMap[name], v)
		}

		// type index
		typeName := TypeOf(v).Name()
		ts.typeMap[typeName] = append(ts.typeMap[typeName], v)
		ts.definitions = append(ts.definitions, v)
	}

	var err error
	// aggregate
	for k, v := range namesMap {
		if k == "" || len(v) == 0 {
			continue
		}
		if len(v) == 1 {
			ts.nameMap[k] = v[0]
		} else {
			ts.nameMap[k], err = MergeDefinitions(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return ts, nil
}
func MergeDefinitions(definitions []Definition) (Definition, error) {
	var target Definition
	var targetIdx = 0
	for i, v := range definitions {
		if strings.HasSuffix(TypeOf(v).Name(), "TypeDefinition") {
			target = v
			targetIdx = i
			break
		}
	}
	if target == nil {
		return nil, errors.New("gqll: merge target not found")
	}
	for i, v := range definitions {
		if i == targetIdx {
			continue
		}
		MergeDefinition(target, v)
	}

	return definitions[0], nil
}

func MergeDefinition(target Definition, ext Definition) error {

	if feat, ok := target.(featFieldDefinitions); ok {
		origin := feat.GetFieldDefinitions()
		extra := ext.(featFieldDefinitions).GetFieldDefinitions()

		feat.SetFieldDefinitions(append(origin, extra...))
	}
	if feat, ok := target.(featInterfaces); ok {
		origin := feat.GetInterfaces()
		extra := ext.(featInterfaces).GetInterfaces()

		feat.SetInterfaces(append(origin, extra...))
	}

	if feat, ok := target.(featDirectives); ok {
		origin := feat.GetDirectives()
		extra := ext.(featDirectives).GetDirectives()

		feat.SetDirectives(append(origin, extra...))
	}

	return nil
}
